package messages

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// PrepareJustificationsUnmarshalling tests unmarshalling prepare justifications
func PrepareJustificationsUnmarshalling() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[1], 1, 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[2], 2, 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[3], 3, 2),
	}

	prepareMsgs := []*types.SignedSSVMessage{
		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[3], types.OperatorID(3)),
	}

	rcMarshalled := testingutils.MarshalJustifications(rcMsgs)

	prepareMarshalled := testingutils.MarshalJustifications(prepareMsgs)

	msg := testingutils.TestingProposalMessageWithParams(
		ks.OperatorKeys[1], types.OperatorID(1), 2, qbft.FirstHeight, testingutils.TestingQBFTRootData,
		rcMarshalled, prepareMarshalled)

	msgRoot, err := msg.GetRoot()
	if err != nil {
		panic(err.Error())
	}
	encodedMsg, err := msg.Encode()
	if err != nil {
		panic(err.Error())
	}

	return &tests.MsgSpecTest{
		Name: "prepare justification unmarshalling",
		Messages: []*types.SignedSSVMessage{
			msg,
		},
		EncodedMessages: [][]byte{
			encodedMsg,
		},
		ExpectedRoots: [][32]byte{
			msgRoot,
		},
	}
}
