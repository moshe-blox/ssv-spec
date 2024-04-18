package proposal

import (
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// PastRoundProposalPrevNotPrepared tests a valid proposal for past round (not prev prepared)
func PastRoundProposalPrevNotPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 10
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingRoundChangeMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingRoundChangeMessage(ks.OperatorKeys[3], types.OperatorID(3)),
	}

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessageWithRoundAndRC(ks.OperatorKeys[1], types.OperatorID(1), qbft.FirstRound,
			testingutils.MarshalJustifications(rcMsgs)),
	}
	return &tests.MsgProcessingSpecTest{
		Name:           "proposal past round (not prev prepared)",
		Pre:            pre,
		PostRoot:       "ed0b4ac99e52e0e2be985db854913958e62d52a4424bb77fa69fc606a9060bbd",
		InputMessages:  msgs,
		OutputMessages: []*types.SignedSSVMessage{},
		ExpectedError:  "invalid signed message: past round",
	}
}
