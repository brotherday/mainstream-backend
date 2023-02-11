package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brotherday/mainstream/backend/orbit"
	godotenv "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Start OrbitDB connection...")
	orbitdbName := os.Getenv("ORBITDB_NAME")
	DBName := fmt.Sprint("test.", strings.ToLower(orbitdbName))
	fmt.Println("Create Orbit DB instance...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orbitdb := orbit.Must(orbit.NewLocalController(ctx, DBName, orbit.DocumentStore, orbit.DefaultCreateDBOptions(), nil))
	comment := orbit.NewComment(0, []byte(strconv.Itoa(257)), []byte(strconv.Itoa(32)), nil, time.Now(), "Hey! This is my first comment.")
	id := comment.Id.String()
	document := map[string]any{id: comment, "_id": id}
	db, err := orbitdb.DocumentStore()
	if err != nil {
		log.Fatalln("orbitdb.DocumentStore()", err)
	}
	op, err := db.Put(ctx, document)
	if err != nil {
		log.Fatalln("Put error:", err)
	}
	fmt.Println("Key added.")
	// same as id or fmt.Sprint(document["_id"])
	key := *op.GetKey()
	val, err := db.Get(ctx, key, orbit.CommentStoreGetOptions())
	if err != nil {
		log.Fatalln("Get error:", err)
	}
	fmt.Println("Get value:", val[0])
	op, err = db.Delete(ctx, key)
	if err != nil {
		log.Fatalln("Delete error:", err)
	}
	fmt.Print("Key deleted: ")
	// same as op.GetValue()
	fmt.Println(op.Marshal())
	db.Close()
	fmt.Println("Store closed.")
}
