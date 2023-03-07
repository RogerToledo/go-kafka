package main

import (
	"fmt"
	"log"
	"me/go-kafka/common"
	"me/go-kafka/kafka"
)

func main() {
	//producer()
	consumer()
}

func producer() {
	file := "data/data.txt"

	lines := common.ReadFile(file)
		
	err, count := kafka.Producer("go", lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %d messages", count)
}

func consumer() {
	err, messages := kafka.Consumer("go", "consumer1")
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range(messages) {
		fmt.Println(msg)
	}
}