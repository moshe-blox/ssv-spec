package messages

import (
	"crypto/rsa"

	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// SignedMessageSigner0 tests SignedMessage signer == 0
func SignedMessageSigner0() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.TestingCommitMultiSignerMessage(
		[]*rsa.PrivateKey{
			ks.OperatorKeys[1],
			ks.OperatorKeys[2],
			ks.OperatorKeys[3],
		},
		[]types.OperatorID{1, 2, 0},
	)

	return &tests.MsgSpecTest{
		Name: "signer 0",
		Messages: []*types.SignedSSVMessage{
			msg,
		},
		ExpectedError: "signer ID 0 not allowed",
	}
}
