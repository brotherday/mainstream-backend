package comment

import (
	"context"
	"encoding/base32"
	"log"
	"math/big"
	"time"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	"github.com/ethereum/go-ethereum/common"
)

type Comment struct {
	Id           *big.Int
	Presentation common.Address
	Owner        common.Address
	PrevComment  *big.Int
	Datetime     time.Time
	Content      string
}

func NewComment(id uint64, presentation []byte, owner []byte, prevComment *uint64, datetime time.Time, content string) *Comment {
	var prevId *big.Int
	if prevComment != nil {
		prevId = big.NewInt(int64(*prevComment))
	}
	return &Comment{
		Id:           big.NewInt(int64(id)),
		Presentation: common.BytesToAddress(presentation),
		PrevComment:  prevId,
		Datetime:     datetime,
		Content:      content,
	}
}

type CommentStore interface {
	iface.DocumentStore
}

func NewCommentStore(ctx context.Context, name string, dbOptions *orbitdb.CreateDBOptions, callback func(ctx context.Context, name string, dbOptions *orbitdb.CreateDBOptions) (CommentStore, error)) (CommentStore, error) {
	if callback == nil {
		log.Fatalln("Please provide orbit.")
	}
	address := base32.HexEncoding.EncodeToString([]byte(name))
	if dbOptions.Replicate == nil {
		dbOptions.Replicate = boolPtr(false)
	}
	if dbOptions.Replicate == nil {
		dbOptions.LocalOnly = boolPtr(false)
	}
	return callback(ctx, address, dbOptions)
}

func Must(store CommentStore, err error) CommentStore {
	if err != nil {
		log.Fatalln(err)
	}
	return store
}

func boolPtr(b bool) *bool {
	return &b
}
