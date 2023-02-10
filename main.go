package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	"github.com/brotherday/mainstream/backend/comment"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	godotenv "github.com/joho/godotenv"
)

func createStoreInstance(ctx context.Context, name string, dbOptions *orbitdb.CreateDBOptions) (comment.CommentStore, error) {
	fmt.Println("Start IPFS connection...")
	ipfs, err := httpapi.NewLocalApi()
	if err != nil {
		log.Fatalln("Connection error:", err)
	}
	var newOrbitDBOptions *orbitdb.NewOrbitDBOptions
	fmt.Println("Start OrbitDB connection...")
	orbit, err := orbitdb.NewOrbitDB(ctx, ipfs, newOrbitDBOptions)
	if err != nil {
		log.Fatalln("Database error:", err)
	}
	return orbit.Docs(ctx, name, dbOptions)
}

func main() {
	godotenv.Load()
	fmt.Println("Start OrbitDB connection...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var orbitdbName = os.Getenv("ORBITDB_NAME")
	var DBName = fmt.Sprint("test.", strings.ToLower(orbitdbName))
	var dbOptions = &orbitdb.CreateDBOptions{
		LocalOnly: boolPtr(true),
		Replicate: boolPtr(false),
	}
	fmt.Println("Create Orbit DB instance...")
	db := comment.Must(comment.NewCommentStore(ctx, DBName, dbOptions, createStoreInstance))
	comment := comment.NewComment(0, []byte(strconv.Itoa(257)), []byte(strconv.Itoa(32)), nil, time.Now(), "Hey! This is my first comment.")
	id := comment.Id.String()
	document := map[string]any{"example": comment, "_id": id}
	op, err := db.Put(ctx, document)
	if err != nil {
		log.Fatalln("Put error:", err)
	}
	fmt.Println("Key added.")
	// same as id or fmt.Sprint(document["_id"])
	key := *op.GetKey()
	var opts = &iface.DocumentStoreGetOptions{
		CaseInsensitive: false,
		PartialMatches:  true,
	}
	val, err := db.Get(ctx, key, opts)
	if err != nil {
		log.Fatalln("Get error:", err)
	}
	fmt.Println("Get value:", val[0])
	op, err = db.Delete(ctx, key)
	if err != nil {
		log.Fatalln("Delete error:", err)
	}
	fmt.Print("Key deleted: ")
	fmt.Println(op.Marshal()) // same as op.GetValue()
	db.Close()
	fmt.Println("Store closed.")
}

func boolPtr(b bool) *bool {
	return &b
}
