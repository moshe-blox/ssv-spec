package roundchange

import (
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// JustificationDuplicateMsg tests a duplicate justification msg which shouldn't result in a justification quorum
func JustificationDuplicateMsg() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 2
	ks := testingutils.Testing4SharesSet()

	prepareMsgs := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
	}
	msgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.OperatorKeys[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(prepareMsgs)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:           "justification duplicate msg",
		Pre:            pre,
		PostRoot:       "96e6d7bdbb98a2d9937f3d97d6aa096bd3a58f923b61012048ac70ad52765919",
		InputMessages:  msgs,
		OutputMessages: []*types.SignedSSVMessage{},
		ExpectedError:  "invalid signed message: no justifications quorum",
	}
}
