package main

// /Volumes/mac-extention/workDir/github/go-grpc-node/go-client/services/go-grpc-node/post_service.pb.go
import (
	"context"
	"fmt"
	"log"
	"time"

	go_grpc_node "github.com/vinayakPhutane/go-grpc-node/services/go-grpc-node"
	"google.golang.org/grpc"
)

func getGRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial("localhost:8081", opts...)

	if err != nil {
		log.Fatal("Fail to dial %v", err)
	}
	return conn
}

func main() {
	conn := getGRPCClient()

	defer conn.Close()

	messageData := go_grpc_node.Post{
		Id:    10,
		Title: "Mayur",
		Text:  "phutane",
	}

	client := go_grpc_node.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 190*time.Second)

	posts, err := client.GetPosts(ctx, &go_grpc_node.Empty{})

	if err != nil {
		log.Fatal("err %v", err)
	}

	for _, post := range posts.GetPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Text)
	}

	post, er := client.PostPosts(ctx, &messageData)

	if er != nil {
		log.Fatal("error while post call %v", er)
	}

	fmt.Println("data %v", post)
}
