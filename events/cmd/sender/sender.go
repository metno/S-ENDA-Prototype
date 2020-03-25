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
	var env config.EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
	os.Exit(_main(os.Args[1:], env))
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

func _main(args []string, env config.EnvConfig) int {
	source, err := url.Parse("https://github.com/metno/S-ENDA-Prototype/events/cmd/sender")
	if err != nil {
		log.Printf("failed to parse source url, %v", err)
		return 1
	}

	seq := 0

	contentType := "application/json"
	t, err := cloudeventsnats.New(env.NATSServer, env.Subject)
	if err != nil {
		log.Printf("failed to create nats transport, %s", err.Error())
		return 1
	}
	c, err := client.New(t)
	if err != nil {
		log.Printf("failed to create client, %s", err.Error())
		return 1
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
			log.Printf("failed to send: %v", err)
			return 1
		}
		seq++
		time.Sleep(2000 * time.Millisecond)

	}

	return 0
}
