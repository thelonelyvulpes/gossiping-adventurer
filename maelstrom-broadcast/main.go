package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
	"strconv"
	"strings"
)

func main() {
	n := maelstrom.NewNode()

	var idx *uint64
	var broadcastVal *int
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		// Unmarshal the message req as an loosely-typed map.
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		r := body["message"].(string)
		val, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		*broadcastVal = val

		// Update the message type to return back.
		req["type"] = "broadcast_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, req)
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshal the message reqbody as an loosely-typed map.
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}

		// Update the message type to return back.
		req["type"] = "read_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, req)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshal the message req as an loosely-typed map.
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		
		
		// Update the message type to return back.
		req["type"] = "topology_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, req)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
