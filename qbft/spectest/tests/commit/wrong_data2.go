package commit

import (
	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// WrongData2 tests a single commit received with a different commit data than the prepared data
func WrongData2() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessage(ks.OperatorKeys[1], 1),

		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], 1),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], 2),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[3], 3),

		testingutils.TestingCommitMessageWrongRoot(ks.OperatorKeys[1], 1),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "commit data != prepared data",
		Pre:           pre,
		PostRoot:      "738fff32174f946a02813c45bfae8807d761c0987d3360d3bf00d7468c78058a",
		InputMessages: msgs,
		OutputMessages: []*types.SignedSSVMessage{
			testingutils.TestingPrepareMessage(ks.OperatorKeys[1], 1),
			testingutils.TestingCommitMessage(ks.OperatorKeys[1], 1),
		},
		ExpectedError: "invalid signed message: proposed data mistmatch",
	}
}
