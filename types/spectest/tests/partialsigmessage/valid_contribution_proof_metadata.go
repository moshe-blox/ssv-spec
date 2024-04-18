package partialsigmessage

import (
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// ValidContributionProofMetaData tests a PartialSignatureMessage for contribution proof metadata valid
func ValidContributionProofMetaData() *MsgSpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)
	msg.Type = types.ContributionProofs

	return &MsgSpecTest{
		Name:     "valid meta data when type ContributionProofs",
		Messages: []*types.PartialSignatureMessages{msg},
	}
}
