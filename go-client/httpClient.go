package main

import (
	"encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"time"
	// "bytes"

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

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}
	fmt.Printf("server: request body: %s\n", reqBody)

	target := os.Getenv("TARGET")

	if target == "" {
		target = "World"
	}

	conn := getGRPCClient()
	

	defer conn.Close()

	// messageData := go_grpc_node.Post{
	// 	Id:    reqBody.id || "",
	// 	Title: reqBody.title || "",
	// 	Text:  reqBody.text || "",
	// }

	client := go_grpc_node.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var messageData  go_grpc_node.Post

	err1 := json.Unmarshal(reqBody, &messageData)

	if err1 != nil {
		fmt.Println("Error while parsing data %v", err1)
        msg := "Request body must only contain a single JSON object 1"
        http.Error(w, msg, http.StatusBadRequest)
        return
    }

	post, er := client.PostPosts(ctx, &messageData)

	if er != nil {
		log.Fatal("error while post call %v", er)
	}

	fmt.Fprintf(w, "Hello %s!\n", post)
}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}