package share

import (
	"crypto/rsa"

	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// NoQuorumDuplicate tests msg with < unique 2f+1 signers (but 2f+1 signers including duplicates)
func NoQuorumDuplicate() *ShareTest {
	ks := testingutils.Testing4SharesSet()
	share := testingutils.TestingShare(ks)

	msg := testingutils.TestingCommitMultiSignerMessage([]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[3], ks.OperatorKeys[2]}, []types.OperatorID{1, 3, 2})
	msg.OperatorIDs = []types.OperatorID{1, 1, 2}

	return &ShareTest{
		Name:                     "no quorum duplicate",
		Share:                    *share,
		Message:                  *msg,
		ExpectedHasPartialQuorum: true,
		ExpectedHasQuorum:        false,
		ExpectedFullCommittee:    false,
		ExpectedError:            "non unique signer",
	}
}
