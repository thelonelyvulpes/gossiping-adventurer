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
	idx := make([]uint64, 1)
	
	n.Handle("generate", func(msg maelstrom.Message) error {
		if idx[0] == 0 {
			for i := 0; i < len(n.NodeIDs()); i++ {
				if strings.Compare(n.ID(), n.NodeIDs()[i]) == 0 {
					idx[0] = uint64(i + 1) * NodeValueSpace
					break
				}
			}
		}

		// Unmarshal the message req as an loosely-typed map.
		var req map[string]any
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}
		
		res := make(map[string]any, 2)
		// Update the message type to return back.
		res["type"] = "generate_ok"
		res["id"] = idx[0]
		idx[0]++

		// Echo the original message back with the updated message type.
		return n.Reply(msg, res)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
