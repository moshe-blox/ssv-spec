package roundchange

import (
	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// QuorumNotPrepared tests a round change quorum for non-prepared state
func QuorumNotPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 2
	ks := testingutils.Testing4SharesSet()

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round change not prepared",
		Pre:           pre,
		PostRoot:      "4e4bf729cf6b51c775dd8652f9cb56f2bde8564a409c4e37df3e6443b2e715d3",
		InputMessages: msgs,
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingProposalMessageWithRoundAndRC(ks.OperatorKeys[1], types.OperatorID(1), 2,
				testingutils.MarshalJustifications(msgs)),
		},
	}
}
