package messages

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// CommitDataEncoding tests encoding CommitData
func CommitDataEncoding() *tests.MsgSpecTest {
	ks := testingutils.Testing4SharesSet()
	msg := testingutils.TestingCommitMessage(ks.Shares[1], types.OperatorID(1))

	r, _ := msg.GetRoot()
	b, _ := msg.Encode()

	return &tests.MsgSpecTest{
		Name: "commit data nil or len 0",
		Messages: []*qbft.SignedMessage{
			msg,
		},
		EncodedMessages: [][]byte{
			b,
		},
		ExpectedRoots: [][32]byte{
			r,
		},
	}
}
