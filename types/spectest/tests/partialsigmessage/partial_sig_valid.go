package partialsigmessage

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// PartialSigValid tests PostConsensusMessage sig == 96 bytes
func PartialSigValid() *MsgSpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)

	return &MsgSpecTest{
		Name: "partial sig valid",
		Messages: []*types.PartialSignatureMessages{
			msg,
		},
	}
}
