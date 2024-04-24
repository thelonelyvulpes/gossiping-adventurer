package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
	"strings"
)

const NodeValueSpace = 1000000000000

func main() {
	n := maelstrom.NewNode()
	var idx *uint64

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		if *idx == 0 {
			for i := 0; i < len(n.NodeIDs()); i++ {
				if strings.Compare(n.ID(), n.NodeIDs()[i]) == 0 {
					*idx = uint64(i+1) * NodeValueSpace
					break
				}
			}
		}

		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "broadcast_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "read_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "topology_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
