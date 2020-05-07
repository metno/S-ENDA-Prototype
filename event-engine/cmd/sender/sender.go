package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	cloudeventsnats "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/nats"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/metno/S-ENDA-Prototype/events/internal/config"
	"github.com/metno/S-ENDA-Prototype/events/pkg/datastructs"
)

func main() {
	ch := make(chan bool)

	var env config.EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
	os.Exit(startDataEventsServer(ch, env))
}

// Simple holder for the sending sample.
type Demo struct {
	Message string
	Source  url.URL
	Target  url.URL

	Client client.Client
}

func (d *Demo) Send(eventContext cloudevents.EventContext, i int) (context.Context, *cloudevents.Event, error) {
	event := cloudevents.Event{
		Context: eventContext,
		Data: &datastructs.Example{
			Sequence: i,
			Message:  d.Message,
		},
	}
	return d.Client.Send(context.Background(), event)
}

func stopDataEventsServer(ch chan bool) {
	ch <- true
}

func startDataEventsServer(ch chan bool, env config.EnvConfig) int {
	source, err := url.Parse("https://github.com/metno/S-ENDA-Prototype/events/cmd/sender")
	if err != nil {
		log.Fatalf("failed to parse source url, %v", err)
	}

	seq := 0

	contentType := "application/json"
	t, err := cloudeventsnats.New(env.NATSServer, env.Subject)
	if err != nil {
		log.Fatalf("Failed to create nats transport. NATSServer %s, Subject %s, %s", env.NATSServer, env.Subject, err.Error())
	}
	c, err := client.New(t)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}

	d := &Demo{
		Message: fmt.Sprintf("Heartbeat!"),
		Source:  *source,
		Client:  c,
	}

	now := time.Now()

	ctx := cloudevents.EventContextV1{
		SpecVersion:     "1.0",
		Type:            "com.cloudevents.sample.sent",
		Source:          types.URIRef{URL: d.Source},
		ID:              uuid.New().String(),
		Time:            &types.Timestamp{Time: now},
		DataContentType: &contentType,
	}

	fmt.Printf("Sending an event every %d seconds\n", 2)
	for {
		if _, _, err := d.Send(&ctx, seq); err != nil {
			log.Fatalf("failed to send: %v", err)
			return 1
		}
		seq++
		time.Sleep(2000 * time.Millisecond)

		select {
		case <-ch:
			fmt.Println("Received termination message. Switching off.")
			return 0
		default:
			continue
		}
	}
}
