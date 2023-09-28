package main

import (
	"context"
	"fmt"
	"gRPC-Streaming-Data-Transfer-App/main_task/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	protos.UnimplementedNumberStreamServer
	clientConn    *grpc.ClientConn
	clientStream  protos.NumberStream_StartStreamServer
	stopSignal    chan struct{}
	streamClients map[string]chan struct{}
}

func (s *Server) Connect(ctx context.Context, req *protos.LoginMessage) (*protos.Empty, error) {
	fmt.Printf("Client login: %s, Client password: %s\n", req.Username, req.Password)
	return &protos.Empty{}, nil
}

func (s *Server) StartStream(req *protos.StartStreamMessage, stream protos.NumberStream_StartStreamServer) error {
	s.clientStream = stream
	interval := time.Millisecond * time.Duration(req.IntervalMs)
	s.stopSignal = make(chan struct{})
	ctx, cancel := context.WithCancel(stream.Context())

	stopped := false

	// Горутина для отправки чисел клиенту с указанным интервалом
	go func() {
		defer cancel()

		value := int32(1)
		for {
			select {
			case <-ctx.Done(): // Завершаем горутину при закрытии контекста
				return
			default:
				err := s.clientStream.Send(&protos.Number{Value: value, Timestamp: time.Now().Unix()})
				if err != nil {
					stopped = true
					return
				}
				value++
				time.Sleep(interval)
			}
		}
	}()

	<-ctx.Done()

	if !stopped {
		// Если горутина не была остановлена, закрыть поток
		s.clientStream.Send(&protos.Number{Value: -1}) // Отправляем специальное значение для индикации завершения потока
	}

	return nil
}

func (s *Server) StopStream(ctx context.Context, req *protos.StopStreamMessage) (*protos.Empty, error) {
	if s.stopSignal != nil {
		close(s.stopSignal)
	}
	return &protos.Empty{}, nil
}

func main() {
	s := grpc.NewServer()
	srv := &Server{}
	protos.RegisterNumberStreamServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
