package prepare

import (
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// UnknownSigner tests a single prepare received with an unknown signer
func UnknownSigner() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	pre := testingutils.BaseInstance()
	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1))

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(5)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "prepare unknown signer",
		Pre:           pre,
		PostRoot:      "2253eea5735c33797cd1f1a1e3ced2cb8b16ee1c78ae1747e18041b67216d622",
		InputMessages: msgs,
		ExpectedError: "invalid signed message: signer not in committee",
	}
}
