package consensus

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/mosheblox/ssv-spec/ssv/spectest/tests"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// ValidDecided13Operators tests a valid decided value (13 operators)
func ValidDecided13Operators() tests.SpecTest {

	panic("implement me")

	ks := testingutils.Testing13SharesSet()
	return &tests.MultiMsgProcessingSpecTest{
		Name: "consensus valid decided 13 operators",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:                    "sync committee contribution",
				Runner:                  testingutils.SyncCommitteeContributionRunner(ks),
				Duty:                    &testingutils.TestingSyncCommitteeContributionDuty,
				Messages:                testingutils.SSVDecidingMsgsV(testingutils.TestSyncCommitteeContributionConsensusData, ks, types.BNRoleSyncCommitteeContribution),
				PostDutyRunnerStateRoot: "2feec534f4ef14ea96e83b17d4feaf3f40e4739a72fc7118c7da1ef2c1331937",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
					testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[1], 1, ks),
				},
			},
			{
				Name:                    "aggregator",
				Runner:                  testingutils.AggregatorRunner(ks),
				Duty:                    &testingutils.TestingAggregatorDuty,
				Messages:                testingutils.SSVDecidingMsgsV(testingutils.TestAggregatorConsensusData, ks, types.BNRoleAggregator),
				PostDutyRunnerStateRoot: "dd79a9da2d025a61d4e4ebdebbc9a331a876a43bb8c8e7c274132ae1c4a35175",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
					testingutils.PostConsensusAggregatorMsg(ks.Shares[1], 1),
				},
			},
			{
				Name:                    "proposer",
				Runner:                  testingutils.ProposerRunner(ks),
				Duty:                    testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages:                testingutils.SSVDecidingMsgsV(testingutils.TestProposerConsensusDataV(spec.DataVersionDeneb), ks, types.BNRoleProposer),
				PostDutyRunnerStateRoot: "942d09a2aaaddf2ef7b2388fd1a9ffb94d590c645fc90c2c6b27fa5414afd61e",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb),
					testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionDeneb),
				},
			},
			{
				Name:                    "proposer (blinded block)",
				Runner:                  testingutils.ProposerBlindedBlockRunner(ks),
				Duty:                    testingutils.TestingProposerDutyV(spec.DataVersionDeneb),
				Messages:                testingutils.SSVDecidingMsgsV(testingutils.TestProposerBlindedBlockConsensusDataV(spec.DataVersionDeneb), ks, types.BNRoleProposer),
				PostDutyRunnerStateRoot: "abf4eaf532eda91fd6db48136938985a655174ab5101e698d71737c0ed48c8c3",
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, spec.DataVersionDeneb),
					testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, spec.DataVersionDeneb),
				},
			},
			{
				Name: "attester and sync committee",
			},
		},
	}
}
