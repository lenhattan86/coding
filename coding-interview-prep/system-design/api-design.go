package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// --- Conceptual REST API Client ---
func callRestAPI() {
	fmt.Println("\n--- Calling Conceptual REST API ---")
	// Simulate GET request
	resp, err := http.Get("http://example.com/api/users/1")
	if err != nil {
		fmt.Printf("REST GET Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("REST GET Response (Status: %s): %s\n", resp.Status, string(body))

	// Simulate POST request
	user := map[string]string{"name": "Alice", "email": "alice@example.com"}
	jsonUser, _ := json.Marshal(user)
	resp, err = http.Post("http://example.com/api/users", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		fmt.Printf("REST POST Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("REST POST Response (Status: %s): %s\n", resp.Status, string(body))
}

// --- Conceptual GraphQL API Client ---
func callGraphQLAPI() {
	fmt.Println("\n--- Calling Conceptual GraphQL API ---")
	query := `
		query {
			user(id: "1") {
				name
				email
			}
		}
	`
	requestBody, _ := json.Marshal(map[string]string{"query": query})
	resp, err := http.Post("http://example.com/graphql", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("GraphQL Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("GraphQL Response (Status: %s): %s\n", resp.Status, string(body))
}

// --- Conceptual gRPC Client (simplified, without actual gRPC setup) ---
// In a real scenario, this would involve generated client code from .proto files.
func callGRPCAPI() {
	fmt.Println("\n--- Calling Conceptual gRPC API ---")
	// Simulate a gRPC call
	// In a real Go gRPC client, you'd have:
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// client := pb.NewUserServiceClient(conn)
	// resp, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: "1"})
	// fmt.Printf("gRPC Response: %v\n", resp)

	fmt.Println("gRPC call simulated: Requesting user '1' via a binary protocol.")
	fmt.Println("Actual gRPC client setup would involve protocol buffers and generated code.")
}

func main() {
	fmt.Println("--- Conceptual API Design Example ---")

	// Simulate network delays for demonstration
	http.DefaultClient.Timeout = 2 * time.Second

	callRestAPI()
	callGraphQLAPI()
	callGRPCAPI()

	fmt.Println("\n--- API Design Example Finished ---")
}
