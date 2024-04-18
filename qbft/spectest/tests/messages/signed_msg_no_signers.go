package messages

import (
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// SignedMsgNoSigners tests SignedMessage len(signers) == 0
func SignedMsgNoSigners() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	msg := testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1))
	msg.OperatorIDs = nil

	return &tests.MsgSpecTest{
		Name: "no signers",
		Messages: []*types.SignedSSVMessage{
			msg,
		},
		ExpectedError: "no signers",
	}
}
