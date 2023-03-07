package kafka

import (
	"context"

	kfk "github.com/segmentio/kafka-go"
)

func Producer(topic string, messages [][]byte) (error, int32) {
	writer := &kfk.Writer {
		Addr: kfk.TCP("localhost:9092"),
		Topic: topic,
	}

	var count int32

	for _, m := range(messages) {
		err := writer.WriteMessages(context.Background(), kfk.Message{
			Value: m,
		})	
		if err != nil {
			return err, count
		}

		count++
	}

	return nil, count
}

func Consumer(topic string, groupID string) (error, []string){
	reader := kfk.NewReader(kfk.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: groupID,
		Topic: topic,
		MinBytes: 0,
		MaxBytes: 10e6, //10MB
	})
	defer reader.Close()

	var menssages []string
	for i := 0; i < 10; i++ {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return err, nil
		} 

		menssages = append(menssages, string(msg.Value))
	}

	return nil, menssages
}