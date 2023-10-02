package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC-Streaming-Data-Transfer-App/main_task/protos"
	"google.golang.org/grpc"
	"log"
	"time"
)

func printBuffer(buffer []*protos.Number) {
	for _, num := range buffer {
		timestamp := time.Unix(num.Timestamp, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Value: %d\tTimestamp: %s\n", num.Value, timestamp)
	}
}

func main() {
	// Определение флагов для строк и чисел
	loginPtr := flag.String("Login", "", "Client login")
	passwordPtr := flag.String("Password", "", "Client password")
	intervalMsPtr := flag.Int("IntervalMs", 500, "Interval in milliseconds")
	bufferSizePtr := flag.Int("BufferSize", 10, "Buffer size")
	stopAfterPtr := flag.Duration("StopAfter", 5*time.Second, "Duration to stop after")

	clientHelpPtr := flag.Bool("help", false, "Show help message for the client")

	// Разбор флагов
	flag.Parse()

	if *clientHelpPtr {
		flag.Usage()
		return
	}

	fmt.Println("Interval in ms: ", int32(*intervalMsPtr))
	fmt.Println("Buffer size: ", int32(*bufferSizePtr))
	fmt.Println("Stream will stop after: ", stopAfterPtr.String())

	// Установка соединения с сервером
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Print(err)
	}

	// Создание клиентского объекта
	c := protos.NewNumberStreamClient(conn)

	// Установка соединения и отправка логин-информации
	_, err = c.Connect(context.Background(), &protos.LoginMessage{Username: *loginPtr, Password: *passwordPtr})
	if err != nil {
		log.Fatal(err)
	}

	// Начало стрима с сервера
	stream, err := c.StartStream(context.Background(), &protos.StartStreamMessage{IntervalMs: int32(*intervalMsPtr)})
	if err != nil {
		log.Fatal(err)
	}

	// Буфер для хранения данных
	buffer := make([]*protos.Number, 0, *bufferSizePtr)

	// Таймер для остановки клиента после указанного времени
	stopTimer := time.NewTimer(*stopAfterPtr)

	for {
		select {
		case <-stopTimer.C:
			_, err := c.StopStream(context.Background(), &protos.StopStreamMessage{})
			if err != nil {
				log.Fatalf("Failed to stop stream: %v", err)
			}
			return
		default:
			num, err := stream.Recv()
			if err != nil {
				log.Printf("Failed to receive data from stream: %v", err)
				return
			}
			buffer = append(buffer, num)

			if len(buffer) == *bufferSizePtr {
				printBuffer(buffer)
				buffer = nil
			}
		}
	}
}
