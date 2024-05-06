package main

import (
	"encoding/json"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {
	broadcastOk := make(map[string]any, 1)
	broadcastOk["type"] = "broadcast_ok"
	var broadcastVal []int

	n := maelstrom.NewNode()
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		f := int(req["message"].(float64))
		broadcastVal = append(broadcastVal, f)
		return n.Reply(msg, broadcastOk)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		res := make(map[string]any, 2)
		res["messages"] = broadcastVal
		res["type"] = "read_ok"

		return n.Reply(msg, res)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}

		res := make(map[string]any, 1)
		res["type"] = "topology_ok"

		return n.Reply(msg, res)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
