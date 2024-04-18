package p2p

import "github.com/mosheblox/ssv-spec/types"

// Broadcaster is the interface used to abstract message broadcasting
type Broadcaster interface {
	Broadcast(message *types.SignedSSVMessage) error
}

// Subscriber is used to abstract topic management
type Subscriber interface {
	Subscribe(vpk types.ValidatorPK) error
}
