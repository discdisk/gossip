package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var (
	lock      sync.RWMutex
	store     = make(map[float64]float64)
	topology  map[string][]string
	neighbers []string
)

func main() {
	n := maelstrom.NewNode()
	pid := os.Getpid()
	n.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "generate_ok"
		body["id"] = fmt.Sprintf("%s-%d-%d-%d", n.ID(), pid, time.Now().UnixNano(), rand.Int31n(11111))

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "broadcast_ok"
		lock.Lock()
		_msg := (body["message"]).(float64)
		store[_msg] = _msg
		lock.Unlock()
		// for _, node := range neighbers {
		// 	n.Send(node)
		// }
		delete(body, "message")

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
		lock.RLock()
		defer lock.RUnlock()
		keys := make([]float64, 0, len(store))
		for v := range store {
			keys = append(keys, v)
		}
		sort.Float64s(keys)
		body["messages"] = keys

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})
	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		b, _ := json.Marshal(body["topology"])
		_ = json.Unmarshal(b, &topology)
		log.Printf("topology is %v ", topology)
		for k := range topology {
			if k == n.ID() {
				neighbers = topology[k]
			}
		}
		log.Println(neighbers)

		// Update the message type to return back.
		body["type"] = "topology_ok"
		delete(body, "topology")

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
