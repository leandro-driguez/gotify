package core

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type RedisPeer struct {
	kademlia.FullNode
}

func NewRedisPeer(isBootstrapNode bool) *RedisPeer {
	db := persistence.NewRedisDb()
	newPeer := kademlia.NewFullNode("0.0.0.0", 8080, 32140, db, isBootstrapNode)

	return &RedisPeer{*newPeer}
}

func (p *RedisPeer) Store(data *[]byte) (string, error) {

	hash := sha1.Sum(*data)
	key := base64.RawStdEncoding.EncodeToString(hash[:])

	fmt.Println("Before StoreValue()")
	_, err := p.StoreValue(key, data)
	if err != nil {
		return "", nil
	}
	fmt.Println("After StoreValue()")

	return key, nil
}