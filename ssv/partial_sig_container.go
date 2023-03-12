package ssv

import (
	"encoding/hex"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/pkg/errors"
)

type PartialSigContainer struct {
	Signatures map[string]map[types.OperatorID][]byte
	// Quorum is the number of min signatures needed for quorum
	Quorum uint64
}

func NewPartialSigContainer(quorum uint64) *PartialSigContainer {
	return &PartialSigContainer{
		Quorum:     quorum,
		Signatures: make(map[string]map[types.OperatorID][]byte),
	}
}

func (ps *PartialSigContainer) AddSignature(sigMsg *types.PartialSignatureMessage) {
	if ps.Signatures[rootHex(sigMsg.SigningRoot)] == nil {
		ps.Signatures[rootHex(sigMsg.SigningRoot)] = make(map[types.OperatorID][]byte)
	}
	m := ps.Signatures[rootHex(sigMsg.SigningRoot)]

	if m[sigMsg.Signer] == nil {
		m[sigMsg.Signer] = make([]byte, 96)
		copy(m[sigMsg.Signer], sigMsg.PartialSignature)
	}
}

func (ps *PartialSigContainer) ReconstructSignature(root [32]byte, validatorPubKey []byte) ([]byte, error) {
	// Reconstruct signatures
	signature, err := types.ReconstructSignatures(ps.Signatures[rootHex(root)])
	if err != nil {
		return nil, errors.Wrap(err, "failed to reconstruct signatures")
	}
	if err := types.VerifyReconstructedSignature(signature, validatorPubKey, root); err != nil {
		return nil, errors.Wrap(err, "failed to verify reconstruct signature")
	}
	return signature.Serialize(), nil
}

func (ps *PartialSigContainer) HasQuorum(root [32]byte) bool {
	return uint64(len(ps.Signatures[rootHex(root)])) >= ps.Quorum
}

func rootHex(r [32]byte) string {
	return hex.EncodeToString(r[:])
}
