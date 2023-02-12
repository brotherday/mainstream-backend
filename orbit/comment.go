package orbit

import (
	"math/big"
	"time"

	"berty.tech/go-orbit-db/iface"
	"github.com/ethereum/go-ethereum/common"
)

type Comment struct {
	Id           *big.Int       `json:"id"`
	Presentation common.Address `json:"presentation"`
	Owner        common.Address `json:"owner"`
	PrevComment  *big.Int       `json:"prev_comment"`
	Datetime     time.Time      `json:"datetime"`
	Content      string         `json:"content"`
}

func NewComment(id uint64, presentation, owner []byte, prevComment *uint64, datetime time.Time, content string) *Comment {
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
