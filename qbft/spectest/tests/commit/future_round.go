package commit

import (
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// FutureRound tests a commit msg received with a future round, should error
func FutureRound() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.OperatorKeys[1], 1)

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingCommitMessageWithRound(ks.OperatorKeys[1], 1, 2),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "commit future round",
		Pre:           pre,
		PostRoot:      "2253eea5735c33797cd1f1a1e3ced2cb8b16ee1c78ae1747e18041b67216d622",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: wrong msg round",
	}
}
