package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	cloudeventsnats "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/nats"
	"github.com/kelseyhightower/envconfig"
	"github.com/metno/S-ENDA-Prototype/events/internal/config"
	"github.com/metno/S-ENDA-Prototype/events/pkg/datastructs"
	"golang.org/x/net/websocket"
)

var eventsChannel chan datastructs.Example

func updatesHandler(ws *websocket.Conn) {
	for {
		event := <-eventsChannel
		ws.Write([]byte(event.Message))
	}
}

func main() {
	var env config.EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
	os.Exit(_main(os.Args[1:], env))
}

func receive(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	fmt.Printf("Got Event Context: %+v\n", event.Context)

	data := &datastructs.Example{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}
	fmt.Printf("Got Data: %+v\n", data)

	eventsChannel <- *data

	fmt.Printf("----------------------------\n")
	return nil
}

func _main(args []string, env config.EnvConfig) int {
	fmt.Printf("Receiver running...\n")
	ctx := context.Background()

	eventsChannel = make(chan datastructs.Example, 10)

	log.Printf("Starting websocket server")
	http.Handle("/updates", websocket.Handler(updatesHandler))
	go http.ListenAndServe(":8084", nil)

	t, err := cloudeventsnats.New(env.NATSServer, env.Subject)
	if err != nil {
		log.Fatalf("failed to create nats transport, %s", err.Error())
	}
	c, err := client.New(t)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}

	log.Printf("Starting receiver...")
	if err := c.StartReceiver(ctx, receive); err != nil {
		log.Fatalf("failed to start nats receiver, %s", err.Error())
	}

	<-ctx.Done()
	return 0
}
