package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("echo", func(msg maelstrom.Message) error {
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}

		res := make(map[string]any, 2)
		res["type"] = "echo_ok"
		res["echo"] = req["echo"]

		return n.Reply(msg, res)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
