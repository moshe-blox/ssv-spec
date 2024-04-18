package messages

import (
	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// CreateProposalNotPreviouslyPrepared tests creating a proposal msg, non-first round and not previously prepared
func CreateProposalNotPreviouslyPrepared() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.CreateMsgSpecTest{
		CreateType: tests.CreateProposal,
		Name:       "create proposal not previously prepared",
		Value:      [32]byte{1, 2, 3, 4},
		RoundChangeJustifications: []*types.SignedSSVMessage{
			testingutils.TestingProposalMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
			testingutils.TestingProposalMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
			testingutils.TestingProposalMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),
		},
		ExpectedRoot: "310725e97a6c1566433ca70d4ae3ff4b91f8d32809a753ea2c64125a1bb98db2",
	}
}
