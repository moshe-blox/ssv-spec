package testingutils

import (
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/types"
)

var SSVDecidingMsgsV = func(consensusData *types.ConsensusData, ks *TestKeySet, role types.BeaconRole) []*types.SignedSSVMessage {
	id := types.NewMsgID(TestingSSVDomainType, TestingValidatorPubKey[:], role)

	signedF := func(partialSigMsg *types.PartialSignatureMessages) *types.SignedSSVMessage {
		byts, _ := partialSigMsg.Encode()
		ssvMsg := &types.SSVMessage{
			MsgType: types.SSVPartialSignatureMsgType,
			MsgID:   id,
			Data:    byts,
		}
		signer := partialSigMsg.Messages[0].Signer
		sig, err := NewTestingOperatorSigner(ks, signer).SignSSVMessage(ssvMsg)
		if err != nil {
			panic(err)
		}
		return &types.SignedSSVMessage{
			OperatorIDs: []types.OperatorID{signer},
			Signatures:  [][]byte{sig},
			SSVMessage:  ssvMsg,
		}
	}

	// pre consensus msgs
	base := make([]*types.SignedSSVMessage, 0)
	if role == types.BNRoleProposer {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, signedF(PreConsensusRandaoMsgV(ks.Shares[types.OperatorID(i)], types.OperatorID(i), consensusData.Version)))
		}
	}
	if role == types.BNRoleAggregator {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, signedF(PreConsensusSelectionProofMsg(ks.Shares[types.OperatorID(i)], ks.Shares[types.OperatorID(i)], types.OperatorID(i), types.OperatorID(i))))
		}
	}
	if role == types.BNRoleSyncCommitteeContribution {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, signedF(PreConsensusContributionProofMsg(ks.Shares[types.OperatorID(i)], ks.Shares[types.OperatorID(i)], types.OperatorID(i), types.OperatorID(i))))
		}
	}

	// consensus and post consensus
	qbftMsgs := SSVDecidingMsgsForHeight(consensusData, id[:], qbft.Height(consensusData.Duty.Slot), ks)
	base = append(base, qbftMsgs...)
	return base
}

var ExpectedSSVDecidingMsgsV = func(consensusData *types.ConsensusData, ks *TestKeySet, role types.BeaconRole) []*types.SignedSSVMessage {
	id := types.NewMsgID(TestingSSVDomainType, TestingValidatorPubKey[:], role)

	ssvMsgF := func(partialSigMsg *types.PartialSignatureMessages) *types.SignedSSVMessage {
		byts, _ := partialSigMsg.Encode()

		ssvMsg := &types.SSVMessage{
			MsgType: types.SSVPartialSignatureMsgType,
			MsgID:   id,
			Data:    byts,
		}
		signer := partialSigMsg.Messages[0].Signer
		sig, err := NewTestingOperatorSigner(ks, signer).SignSSVMessage(ssvMsg)
		if err != nil {
			panic(err)
		}
		return &types.SignedSSVMessage{
			OperatorIDs: []types.OperatorID{signer},
			Signatures:  [][]byte{sig},
			SSVMessage:  ssvMsg,
		}
	}

	// pre consensus msgs
	base := make([]*types.SignedSSVMessage, 0)
	if role == types.BNRoleProposer {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, ssvMsgF(PreConsensusRandaoMsgV(ks.Shares[types.OperatorID(i)], types.OperatorID(i), consensusData.Version)))
		}
	}
	if role == types.BNRoleAggregator {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, ssvMsgF(PreConsensusSelectionProofMsg(ks.Shares[types.OperatorID(i)], ks.Shares[types.OperatorID(i)], types.OperatorID(i), types.OperatorID(i))))
		}
	}
	if role == types.BNRoleSyncCommitteeContribution {
		for i := uint64(1); i <= ks.Threshold; i++ {
			base = append(base, ssvMsgF(PreConsensusContributionProofMsg(ks.Shares[types.OperatorID(i)], ks.Shares[types.OperatorID(i)], types.OperatorID(i), types.OperatorID(i))))
		}
	}

	// consensus and post consensus
	qbftMsgs := SSVExpectedDecidingMsgsForHeight(consensusData, id[:], qbft.Height(consensusData.Duty.Slot), ks)
	base = append(base, qbftMsgs...)
	return base
}
