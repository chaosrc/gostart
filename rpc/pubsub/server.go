package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/moby/moby/pkg/pubsub"

	"example.com/hello/rpc/pubsub/sub"
	"google.golang.org/grpc"
)

type PublishServiceImp struct {
	sub *pubsub.Publisher
}

func (p *PublishServiceImp) Publish(ctx context.Context, arg *sub.String) (*sub.String, error) {
	p.sub.Publish(arg.GetValue())
	return &sub.String{}, nil
}
func (p *PublishServiceImp) Subsribe(arg *sub.String, stream sub.PubsubService_SubsribeServer) error {
	ch := p.sub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&sub.String{Value: v.(string)}); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	serv := grpc.NewServer()
	sub.RegisterPubsubServiceServer(serv, &PublishServiceImp{pubsub.NewPublisher(time.Millisecond*100, 10)})

	lis, err := net.Listen("tcp", ":8004")
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
