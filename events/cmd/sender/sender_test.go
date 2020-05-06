package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/metno/S-ENDA-Prototype/events/internal/config"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats-server/v2/test"
)

const TestPort = 8369

func RunServerOnPort(port int) *server.Server {
	opts := test.DefaultTestOptions
	opts.Port = port
	return RunServerWithOptions(&opts)
}

func RunServerWithOptions(opts *server.Options) *server.Server {
	return test.RunServer(opts)
}

func TestSomething(t *testing.T) {

	s := RunServerOnPort(TestPort)
	defer s.Shutdown()

	sURL := fmt.Sprintf("nats://127.0.0.1:%d", TestPort)
	var env config.EnvConfig
	env.NATSServer = sURL
	env.Subject = "sample"

	ch := make(chan bool)
	go startDataEventsServer(ch, env)
	time.Sleep(3 * time.Second)
	stopDataEventsServer(ch)
}
