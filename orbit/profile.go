package orbit

import (
	"berty.tech/go-orbit-db/iface"
	"github.com/ethereum/go-ethereum/common"
)

type Profile struct {
	Owner   common.Address `json:"owner"`
	Name    *string        `json:"name"`
	Email   *[]byte        `json:"email"`
	Picture *[]byte        `json:"picture"`
	Role    common.Hash    `json:"role"`
}

func NewProfile(owner common.Address, name *string, email, Picture *[]byte, Role *common.Hash) (*Profile, error) {
	if Role == nil {
		*Role = common.HexToHash("Public")
	}
	return &Profile{
		Owner:   owner,
		Name:    name,
		Email:   email,
		Picture: Picture,
		Role:    *Role,
	}, nil
}

type ProfileStore interface {
	iface.DocumentStore
}

func ProfileStoreGetOptions() *iface.DocumentStoreGetOptions {
	return &iface.DocumentStoreGetOptions{
		CaseInsensitive: true,
		PartialMatches:  false,
	}
}
