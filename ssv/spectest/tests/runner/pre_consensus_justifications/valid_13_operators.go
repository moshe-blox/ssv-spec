package pre_consensus_justifications

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/ssv/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// Valid13Operators tests a 13 operator valid pre-consensus justification
func Valid13Operators() tests.SpecTest {
	ks := testingutils.Testing13SharesSet()

	msgF := func(obj *types.ConsensusData, id []byte) *types.SignedSSVMessage {
		fullData, _ := obj.Encode()
		root, _ := qbft.HashDataRoot(fullData)
		msg := &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     1,
			Round:      qbft.FirstRound,
			Identifier: id,
			Root:       root,
		}
		signed := testingutils.SignQBFTMsg(ks.OperatorKeys[1], 1, msg)
		signed.FullData = fullData

		return signed
	}

	return &tests.MultiMsgProcessingSpecTest{
		Name: "pre consensus valid 13 operator quorum",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:   "sync committee aggregator selection proof",
				Runner: decideFirstHeight(testingutils.SyncCommitteeContributionRunner(ks)),
				Duty:   &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestContributionProofWithJustificationsConsensusData(ks), testingutils.SyncCommitteeContributionMsgID),
				},
				PostDutyRunnerStateRoot: "3fb467f07ac5c1ccfa963e0170aaf390eed37b46e827d6d65acca027191fc085",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "aggregator selection proof",
				Runner: decideFirstHeight(testingutils.AggregatorRunner(ks)),
				Duty:   &testingutils.TestingAggregatorDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestSelectionProofWithJustificationsConsensusData(ks), testingutils.AggregatorMsgID),
				},
				PostDutyRunnerStateRoot: "4b2feec74320a809de3112c15a1491d2aeeda08de63030fb8248b3fbad599c08",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "randao",
				Runner: decideFirstHeight(testingutils.ProposerRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestProposerWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "865622dc1856f64bddd7098351172932cc32f36482621d3139cacae00dec4e15",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb), // broadcasts when starting a new duty
				},
			},
			{
				Name:   "randao (blinded block)",
				Runner: decideFirstHeight(testingutils.ProposerBlindedBlockRunner(ks)),
				Duty:   testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestProposerBlindedWithJustificationsConsensusDataV(ks, spec.DataVersionDeneb), testingutils.ProposerMsgID),
				},
				PostDutyRunnerStateRoot: "94ebf94fd556e8dbeb860a01e28d3fd8eaa9ce4655c7ee71d67dcf175e646153",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb), // broadcasts when starting a new duty
				},
			},
			{

				Name:   "attester",
				Runner: decideFirstHeight(testingutils.CommitteeRunner(ks)),
				Duty:   &testingutils.TestingAttesterDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestAttesterConsensusData, testingutils.AttesterMsgID),
				},
				PostDutyRunnerStateRoot: "723a3a8b1b1e20f28ec99cae097adb85b73473a842fd86686361ba2aef3d911a",
				OutputMessages:          []*types.PartialSignatureMessages{},
			},
			{
				Name:   "sync committee",
				Runner: decideFirstHeight(testingutils.SyncCommitteeRunner(ks)),
				Duty:   &testingutils.TestingSyncCommitteeDuty,
				Messages: []*types.SignedSSVMessage{
					msgF(testingutils.TestSyncCommitteeConsensusData, testingutils.SyncCommitteeMsgID),
				},
				PostDutyRunnerStateRoot: "0459c2be68b353e1a780f8f2e17c7350e32242d9a5c50e30ceac7eab0cc4cb85",
				OutputMessages:          []*types.PartialSignatureMessages{},
			},
		},
	}
}
