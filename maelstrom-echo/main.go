package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message req as an loosely-typed map.
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		
		res := make(map[string]any, 2)
		res["type"] = "echo_ok"
		res["message"] = req["message"]

		// Echo the original message back with the updated message type.
		return n.Reply(msg, res)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
