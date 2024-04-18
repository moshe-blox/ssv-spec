package tests

import (
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
	qbftcomparable "github.com/mosheblox/ssv-spec/types/testingutils/comparable"
)

// HappyFlow tests a simple full happy flow until decided
func HappyFlow() SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	sc := happyFlowStateComparison()

	return &MsgProcessingSpecTest{
		Name:      "happy flow",
		Pre:       pre,
		PostRoot:  sc.Root(),
		PostState: sc.ExpectedState,
		InputMessages: []*types.SignedSSVMessage{
			testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1)),

			testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
			testingutils.TestingPrepareMessage(ks.OperatorKeys[3], types.OperatorID(3)),

			testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[2], types.OperatorID(2)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[3], types.OperatorID(3)),
		},
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		},
	}
}

func happyFlowStateComparison() *qbftcomparable.StateComparison {
	ks := testingutils.Testing4SharesSet()

	state := testingutils.BaseInstance().State
	state.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1))
	state.LastPreparedRound = 1
	state.LastPreparedValue = testingutils.TestingQBFTFullData
	state.Decided = true
	state.DecidedValue = testingutils.TestingQBFTFullData

	state.ProposeContainer = &qbft.MsgContainer{Msgs: map[qbft.Round][]*types.SignedSSVMessage{
		qbft.FirstRound: {
			testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		},
	}}
	state.PrepareContainer = &qbft.MsgContainer{Msgs: map[qbft.Round][]*types.SignedSSVMessage{
		qbft.FirstRound: {
			testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
			testingutils.TestingPrepareMessage(ks.OperatorKeys[3], types.OperatorID(3)),
		},
	}}
	state.CommitContainer = &qbft.MsgContainer{Msgs: map[qbft.Round][]*types.SignedSSVMessage{
		qbft.FirstRound: {
			testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[2], types.OperatorID(2)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[3], types.OperatorID(3)),
		},
	}}

	return &qbftcomparable.StateComparison{ExpectedState: state}
}
