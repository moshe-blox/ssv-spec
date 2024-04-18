package messages

import (
	"crypto/rsa"

	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// SignedMsgDuplicateSigners tests SignedMessage with duplicate signers
func SignedMsgDuplicateSigners() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.TestingCommitMultiSignerMessage(
		[]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[1], ks.OperatorKeys[2]},
		[]types.OperatorID{1, 2, 3},
	)
	msg.OperatorIDs = []types.OperatorID{1, 1, 2}

	return &tests.MsgSpecTest{
		Name: "duplicate signers",
		Messages: []*types.SignedSSVMessage{
			msg,
		},
		ExpectedError: "non unique signer",
	}
}
