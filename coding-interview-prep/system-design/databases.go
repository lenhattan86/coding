package main

import (
	"fmt"
	"time"
)

// User represents a simple data structure for a user.
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// MockDatabase simulates a database interface.
type MockDatabase interface {
	Connect() error
	Disconnect() error
	InsertUser(user User) error
	GetUser(id string) (*User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
}

// SQLMockDB simulates a SQL database.
type SQLMockDB struct {
	users map[string]User
}

func NewSQLMockDB() *SQLMockDB {
	return &SQLMockDB{
		users: make(map[string]User),
	}
}

func (db *SQLMockDB) Connect() error {
	fmt.Println("SQL DB: Connecting...")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("SQL DB: Connected.")
	return nil
}

func (db *SQLMockDB) Disconnect() error {
	fmt.Println("SQL DB: Disconnecting...")
	time.Sleep(20 * time.Millisecond)
	fmt.Println("SQL DB: Disconnected.")
	return nil
}

func (db *SQLMockDB) InsertUser(user User) error {
	fmt.Printf("SQL DB: Inserting user %s...\n", user.ID)
	if _, exists := db.users[user.ID]; exists {
		return fmt.Errorf("user with ID %s already exists", user.ID)
	}
	db.users[user.ID] = user
	time.Sleep(30 * time.Millisecond)
	fmt.Printf("SQL DB: User %s inserted.\n", user.ID)
	return nil
}

func (db *SQLMockDB) GetUser(id string) (*User, error) {
	fmt.Printf("SQL DB: Getting user %s...\n", id)
	user, exists := db.users[id]
	time.Sleep(20 * time.Millisecond)
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	fmt.Printf("SQL DB: User %s retrieved.\n", id)
	return &user, nil
}

func (db *SQLMockDB) UpdateUser(user User) error {
	fmt.Printf("SQL DB: Updating user %s...\n", user.ID)
	if _, exists := db.users[user.ID]; !exists {
		return fmt.Errorf("user with ID %s not found for update", user.ID)
	}
	db.users[user.ID] = user
	time.Sleep(25 * time.Millisecond)
	fmt.Printf("SQL DB: User %s updated.\n", user.ID)
	return nil
}

func (db *SQLMockDB) DeleteUser(id string) error {
	fmt.Printf("SQL DB: Deleting user %s...\n", id)
	if _, exists := db.users[id]; !exists {
		return fmt.Errorf("user with ID %s not found for deletion", id)
	}
	delete(db.users, id)
	time.Sleep(15 * time.Millisecond)
	fmt.Printf("SQL DB: User %s deleted.\n", id)
	return nil
}

// NoSQLMockDB simulates a NoSQL document database.
type NoSQLMockDB struct {
	documents map[string]map[string]interface{} // key -> document (map of string to interface)
}

func NewNoSQLMockDB() *NoSQLMockDB {
	return &NoSQLMockDB{
		documents: make(map[string]map[string]interface{}),
	}
}

func (db *NoSQLMockDB) Connect() error {
	fmt.Println("NoSQL DB: Connecting...")
	time.Sleep(40 * time.Millisecond)
	fmt.Println("NoSQL DB: Connected.")
	return nil
}

func (db *NoSQLMockDB) Disconnect() error {
	fmt.Println("NoSQL DB: Disconnecting...")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("NoSQL DB: Disconnected.")
	return nil
}

func (db *NoSQLMockDB) InsertDocument(id string, doc map[string]interface{}) error {
	fmt.Printf("NoSQL DB: Inserting document %s...\n", id)
	if _, exists := db.documents[id]; exists {
		return fmt.Errorf("document with ID %s already exists", id)
	}
	db.documents[id] = doc
	time.Sleep(25 * time.Millisecond)
	fmt.Printf("NoSQL DB: Document %s inserted.\n", id)
	return nil
}

func (db *NoSQLMockDB) GetDocument(id string) (map[string]interface{}, error) {
	fmt.Printf("NoSQL DB: Getting document %s...\n", id)
	doc, exists := db.documents[id]
	time.Sleep(18 * time.Millisecond)
	if !exists {
		return nil, fmt.Errorf("document with ID %s not found", id)
	}
	fmt.Printf("NoSQL DB: Document %s retrieved.\n", id)
	return doc, nil
}

func main() {
	fmt.Println("--- Conceptual Database Interaction Example ---")

	// --- SQL Database Simulation ---
	fmt.Println("\n--- SQL Database Operations ---")
	sqlDB := NewSQLMockDB()
	sqlDB.Connect()
	defer sqlDB.Disconnect()

	user1 := User{ID: "1", Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()}
	sqlDB.InsertUser(user1)

	retrievedUser, err := sqlDB.GetUser("1")
	if err == nil {
		fmt.Printf("Retrieved SQL User: %+v\n", retrievedUser)
	}

	user1.Email = "alice.smith@example.com"
	sqlDB.UpdateUser(user1)

	sqlDB.DeleteUser("1")
	sqlDB.GetUser("1") // Should fail

	// --- NoSQL Database Simulation ---
	fmt.Println("\n--- NoSQL Database Operations ---")
	nosqlDB := NewNoSQLMockDB()
	nosqlDB.Connect()
	defer nosqlDB.Disconnect()

	doc1 := map[string]interface{}{
		"id":        "doc1",
		"title":     "My First Document",
		"author":    "Bob",
		"tags":      []string{"go", "nosql"},
		"createdAt": time.Now(),
	}
	nosqlDB.InsertDocument("doc1", doc1)

	retrievedDoc, err := nosqlDB.GetDocument("doc1")
	if err == nil {
		fmt.Printf("Retrieved NoSQL Document: %+v\n", retrievedDoc)
	}

	doc1["title"] = "My Updated Document"
	nosqlDB.InsertDocument("doc1", doc1) // In a real NoSQL, this might be an upsert or update

	nosqlDB.GetDocument("doc1")

	fmt.Println("\n--- Database Interaction Example Finished ---")
}
