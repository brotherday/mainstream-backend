package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	godotenv "github.com/joho/godotenv"
)

func Main() {
	main()
}

func main() {
	godotenv.Load()
	fmt.Println("Start OrbitDB connection...")
	orbitdbName := os.Getenv("ORBITDB_NAME")
	DBName := "test." + orbitdbName
	fmt.Println("Create Orbit DB instance...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orbitdb := Must(NewLocalController(ctx, DBName, DocumentStore, DefaultCreateDBOptions(), nil))
	comment := NewComment(0, []byte("257"), []byte("32"), nil, time.Now(), "Hey! This is my first comment.")
	id := comment.Id.String()
	document := map[string]any{id: comment, "_id": id}
	database, err := orbitdb.DocumentStore()
	if err != nil {
		log.Fatalln("orbitdb.DocumentStore()", err)
	}
	op, err := database.Put(ctx, document)
	if err != nil {
		log.Fatalln("Put error:", err)
	}
	fmt.Println("Key added.")
	// same as id or fmt.Sprint(document["_id"])
	key := *op.GetKey()
	val, err := database.Get(ctx, key, CommentStoreGetOptions())
	if err != nil {
		log.Fatalln("Get error:", err)
	}
	fmt.Println("Get value:", val[0])
	op, err = database.Delete(ctx, key)
	if err != nil {
		log.Fatalln("Delete error:", err)
	}
	fmt.Print("Key deleted: ")
	// same as op.GetValue()
	fmt.Println(op.Marshal())
	database.Close()
	fmt.Println("Store closed.")
}
