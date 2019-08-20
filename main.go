package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	notification "github.com/300481/3141-notification"

	"cloud.google.com/go/pubsub"
	"github.com/300481/mq"
)

type Operator struct {
	SystemID string
	Ref      string
	Command  string
	Args     []string
}

const (
	errMissingCommand = "Required environment variable 'COMMAND' not specified."
)

// new creates new Operator struct
func new() (op *Operator, err error) {
	o := &Operator{
		SystemID: os.Getenv("SYSTEM_ID"),
		Ref:      os.Getenv("REF"),
		Command:  os.Getenv("COMMAND"),
		Args:     strings.Split(os.Getenv("ARGS"), " "),
	}
	if len(o.Command) == 0 {
		return nil, errors.New(errMissingCommand)
	}
	return o, nil
}

// handleMessage handles incoming messages
func (o *Operator) handleMessage(ctx context.Context, m *pubsub.Message) {
	message := notification.NewFromJson(m.Data)

	if message.IsSelected(o.SystemID, o.Ref) {
		log.Printf("Received updates for this system. SystemID: %s\n", o.SystemID)

		var metadata []string
		metadata = append(metadata, message.Repository, message.CommitID, strconv.FormatInt(message.PushedAt.Unix(), 10))

		var args []string
		args = append(metadata, o.Args...)

		cmd := exec.Command(o.Command, args...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		cmd.Run()
	} else {
		log.Printf("No updates received for this system. SystemID: %s\n", o.SystemID)
	}

	m.Ack()
}

func main() {
	log.Println("Start 3141-operator")

	// create new Operator configuration from environment
	op, err := new()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Loaded operator config.")

	// create new message queue configuration from environment
	mq := mq.NewGCP()

	// subscribe to the message queue
	err = mq.Subscribe(op.handleMessage)
	if err != nil {
		log.Fatalln(err)
	}
}
