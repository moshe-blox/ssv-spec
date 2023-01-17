package types

import (
	"bytes"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	ssz "github.com/ferranbt/fastssz"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// DomainType is a unique identifier for signatures, 2 identical pieces of data signed with different domains will result in different sigs
type DomainType []byte

var (
	PrimusTestnet = DomainType("primus_testnet")
	ShifuTestnet  = DomainType("shifu_testnet")
)

type SignatureType [4]byte

func (sigType SignatureType) Equal(other SignatureType) bool {
	return bytes.Equal(sigType[:], other[:])
}

var (
	QBFTSignatureType    SignatureType = [4]byte{1, 0, 0, 0}
	PartialSignatureType SignatureType = [4]byte{2, 0, 0, 0}
	DKGSignatureType     SignatureType = [4]byte{3, 0, 0, 0}
)

type BeaconSigner interface {
	// SignBeaconObject returns signature and root.
	SignBeaconObject(obj ssz.HashRoot, domain spec.Domain, pk []byte, domainType spec.DomainType) (Signature, []byte, error)
	// IsAttestationSlashable returns error if attestation is slashable
	IsAttestationSlashable(pk []byte, data *spec.AttestationData) error
	// IsBeaconBlockSlashable returns error if the given block is slashable
	IsBeaconBlockSlashable(pk []byte, block *bellatrix.BeaconBlock) error
}

// SSVSigner used for all SSV specific signing
type SSVSigner interface {
	SignRoot(data Root, sigType SignatureType, pk []byte) (Signature, error)
}

type DKGSigner interface {
	SSVSigner
	// SignDKGOutput signs output according to the SIP https://docs.google.com/document/d/1TRVUHjFyxINWW2H9FYLNL2pQoLy6gmvaI62KL_4cREQ/edit
	SignDKGOutput(output Root, address common.Address) (Signature, error)
	// SignETHDepositRoot signs an ethereum deposit root
	SignETHDepositRoot(root []byte, address common.Address) (Signature, error)
}

// KeyManager is an interface responsible for all key manager functions
type KeyManager interface {
	BeaconSigner
	SSVSigner
	// AddShare saves a share key
	AddShare(shareKey *bls.SecretKey) error
	// RemoveShare removes a share key
	RemoveShare(pubKey string) error
}
