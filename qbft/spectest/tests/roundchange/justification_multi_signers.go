package roundchange

import (
	"crypto/rsa"

	"github.com/mosheblox/ssv-spec/qbft/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// JustificationMultiSigners tests a single prepare justification msg with multiple signers
func JustificationMultiSigners() tests.SpecTest {
	pre := testingutils.BaseInstance()
	pre.State.Round = 2
	ks := testingutils.Testing4SharesSet()

	prepareMsgs := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingPrepareMultiSignerMessage(
			[]*rsa.PrivateKey{ks.OperatorKeys[2], ks.OperatorKeys[3]},
			[]types.OperatorID{2, 3},
		),
	}
	msgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.OperatorKeys[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(prepareMsgs)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:           "justification multi signer",
		Pre:            pre,
		PostRoot:       "96e6d7bdbb98a2d9937f3d97d6aa096bd3a58f923b61012048ac70ad52765919",
		InputMessages:  msgs,
		OutputMessages: []*types.SignedSSVMessage{},
		ExpectedError:  "invalid signed message: round change justification invalid: msg allows 1 signer",
	}
}
