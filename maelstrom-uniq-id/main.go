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
	var idx uint64

	n.Handle("generate", func(msg maelstrom.Message) error {
		if idx == 0 {
			for i := 0; i < len(n.NodeIDs()); i++ {
				if strings.Compare(n.ID(), n.NodeIDs()[i]) == 0 {
					idx = uint64(i+1) * NodeValueSpace
					break
				}
			}
		}

		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}

		res := make(map[string]any, 2)
		res["type"] = "generate_ok"
		res["id"] = idx
		idx++

		return n.Reply(msg, res)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
