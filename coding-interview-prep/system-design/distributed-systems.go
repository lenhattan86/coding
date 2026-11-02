package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

// --- Server Side ---

// Args defines the arguments for the RPC call.
type Args struct {
	A, B int
}

// Calculator is the service that provides RPC methods.
type Calculator int

// Add is an RPC method that adds two numbers.
func (t *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// Multiply is an RPC method that multiplies two numbers.
func (t *Calculator) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func startServer() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	fmt.Println("RPC Server listening on :1234")
	rpc.Accept(listener) // Blocks until listener is closed
}

// --- Client Side ---

func startClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	args := Args{7, 8}
	var reply int

	// Call Add method
	err = client.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Fatal("Calculator.Add error:", err)
	}
	fmt.Printf("Client: Calculator.Add(%d, %d) = %d\n", args.A, args.B, reply)

	// Call Multiply method
	args = Args{5, 6}
	err = client.Call("Calculator.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Calculator.Multiply error:", err)
	}
	fmt.Printf("Client: Calculator.Multiply(%d, %d) = %d\n", args.A, args.B, reply)
}

func main() {
	fmt.Println("--- Conceptual RPC Client/Server Example ---")

	// Start the server in a goroutine
	go startServer()

	// Give the server a moment to start up
	time.Sleep(1 * time.Second)

	// Start the client
	startClient()

	fmt.Println("\n--- RPC Client/Server Example Finished ---")
}
