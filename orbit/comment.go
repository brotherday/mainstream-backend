package orbit

import (
	"math/big"
	"time"

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

func CommentStoreGetOptions() *iface.DocumentStoreGetOptions {
	return &iface.DocumentStoreGetOptions{
		CaseInsensitive: false,
		PartialMatches:  true,
	}
}
