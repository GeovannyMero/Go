package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time string
}

func main() {
	m := Message{"Geocanny", "Hello", "1231"}
	fmt.Println(m)

	b, err := json.Marshal(m)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	fmt.Println("Convert JSON a estructura")

	input := `{
		"Name": "Golden Retriever",
		"Body": "8",
		"time": "Paws"
	}`

	var msg Message

	err = json.Unmarshal([]byte(input), &msg)

	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	fmt.Println(msg)
}
