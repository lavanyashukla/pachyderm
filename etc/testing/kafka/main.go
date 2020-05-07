package main

import (
	"archive/tar"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const defaultGroupID = "test"
const defaultTimeout = 20
const defaultNamedPipe = "/pfs/out"

func main() {
	// Get the connection info from the ENV vars
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	topic := os.Getenv("TOPIC")

	// Set the default values of the configurable variables
	var (
		groupID = defaultGroupID
		timeout = defaultTimeout
		pipe    = defaultNamedPipe
	)

	for {
		if err := func() error {
			// And create a new kafka reader
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:  []string{host + ":" + port},
				Topic:    topic,
				GroupID:  groupID,
				MinBytes: 10e1,
				MaxBytes: 10e6,
			})
			defer reader.Close()

			// read a message
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
			defer func() {
				cancel()
			}()
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				return fmt.Errorf("read message: %v", err)
			}

			// Open the /pfs/out pipe with write only permissons (the pachyderm spout will be reading at the other end of this)
			// Note: it won't work if you try to open this with read, or read/write permissions
			out, err := os.OpenFile(pipe, os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer out.Close()

			tw := tar.NewWriter(out)
			defer tw.Close()

			// give it a unique name
			name := fmt.Sprintf("%v-%v", m.Topic, m.Offset)
			// write the header
			for err = tw.WriteHeader(&tar.Header{
				Name: name,
				Mode: 0600,
				Size: int64(len(m.Value)),
			}); err != nil; {
				if !strings.Contains(err.Error(), "broken pipe") {
					return fmt.Errorf("%s: %v", m.Value, err)
				}
				fmt.Println("broken pipe error")
				// if there's a broken pipe, just give it some time to get ready for the next message
				time.Sleep(time.Duration(timeout) * time.Millisecond)
			}
			// and the message
			fmt.Println("writing", name)
			for _, err = tw.Write(m.Value); err != nil; {
				if !strings.Contains(err.Error(), "broken pipe") {
					return err
				}
				// if there's a broken pipe, just give it some time to get ready for the next message
				time.Sleep(time.Duration(timeout) * time.Millisecond)
			}
			return nil
		}(); err != nil {
			panic(err)
		}
	}
}
