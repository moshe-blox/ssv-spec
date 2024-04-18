package tests

import (
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// ThirteenOperators tests a simple full happy flow until decided
func ThirteenOperators() SpecTest {
	pre := testingutils.ThirteenOperatorsInstance()
	ks := testingutils.Testing13SharesSet()

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[3], types.OperatorID(3)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[4], types.OperatorID(4)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[5], types.OperatorID(5)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[6], types.OperatorID(6)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[7], types.OperatorID(7)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[8], types.OperatorID(8)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[9], types.OperatorID(9)),

		testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[3], types.OperatorID(3)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[4], types.OperatorID(4)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[5], types.OperatorID(5)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[6], types.OperatorID(6)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[7], types.OperatorID(7)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[8], types.OperatorID(8)),
		testingutils.TestingCommitMessage(ks.OperatorKeys[9], types.OperatorID(9)),
	}
	return &MsgProcessingSpecTest{
		Name:          "happy flow thirteen operators",
		Pre:           pre,
		PostRoot:      "d41fb2f41d75191828e05db297c3e61f63d83666b75559bbfa120041e860542a",
		InputMessages: msgs,
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
			testingutils.TestingCommitMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		},
	}
}
