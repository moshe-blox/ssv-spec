package postconsensus

import (
	"github.com/attestantio/go-eth2-client/spec"

	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/ssv"
	ssvcomparable "github.com/mosheblox/ssv-spec/ssv/spectest/comparable"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
	"github.com/mosheblox/ssv-spec/types/testingutils/comparable"
)

// quorumSyncCommitteeContributionSC returns state comparison object for the Quorum SyncCommitteeContribution versioned spec test
// runner should finish since quorum is achieved
func quorumSyncCommitteeContributionSC() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestSyncCommitteeContributionConsensusData
	cdBytes := testingutils.TestSyncCommitteeContributionConsensusDataByts

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.SyncCommitteeContributionRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[1], 1, ks))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[2], 2, ks))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommitteeContribution(nil, testingutils.PostConsensusSyncCommitteeContributionMsg(ks.Shares[3], 3, ks))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}

// quorumSyncCommitteeSC returns state comparison object for the Quorum SyncCommittee versioned spec test
// runner should finish since quorum was achieved
func quorumSyncCommitteeSC() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestSyncCommitteeConsensusData
	cdBytes := testingutils.TestSyncCommitteeConsensusDataByts

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.SyncCommitteeRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[1], 1))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[2], 2))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgSyncCommittee(nil, testingutils.PostConsensusSyncCommitteeMsg(ks.Shares[3], 3))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}

// quorumAggregatorSC returns state comparison object for the Quorum Aggregator versioned spec test
// runner should finish since quorum was achieved
func quorumAggregatorSC() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestAggregatorConsensusData
	cdBytes := testingutils.TestAggregatorConsensusDataByts

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.AggregatorRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAggregator(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[1], 1))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAggregator(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[2], 2))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAggregator(nil, testingutils.PostConsensusAggregatorMsg(ks.Shares[3], 3))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}

// quorumAttesterSC returns state comparison object for the Quorum Attester versioned spec test
// runner should finish since quorum was achieved
func quorumAttesterSC() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestAttesterConsensusData
	cdBytes := testingutils.TestAttesterConsensusDataByts

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.CommitteeRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[2], 2, qbft.FirstHeight))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgAttester(nil, testingutils.PostConsensusAttestationMsg(ks.Shares[3], 3, qbft.FirstHeight))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}

// quorumProposerSC returns state comparison object for the Quorum Proposer versioned spec test
// runner should finish since quorum was achieved
func quorumProposerSC(version spec.DataVersion) *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestProposerConsensusDataV(version)
	cdBytes := testingutils.TestProposerConsensusDataBytsV(version)

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.ProposerRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, version))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[2], 2, version))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[3], 3, version))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}

// quorumBlindedProposerSC returns state comparison object for the Quorum Blinded Proposer versioned spec test
// runner should finish since quorum was achieved
func quorumBlindedProposerSC(version spec.DataVersion) *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	cd := testingutils.TestProposerBlindedBlockConsensusDataV(version)
	cdBytes := testingutils.TestProposerBlindedBlockConsensusDataBytsV(version)

	return &comparable.StateComparison{
		ExpectedState: func() ssv.Runner {
			ret := testingutils.ProposerBlindedBlockRunner(ks)
			ret.GetBaseRunner().State = &ssv.State{
				PreConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{},
				),
				PostConsensusContainer: ssvcomparable.SetMessagesInContainer(
					ssv.NewPartialSigContainer(3),
					[]*types.SignedSSVMessage{
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[1], 1, version))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[2], 2, version))),
						testingutils.SignPartialSigSSVMessage(ks, testingutils.SSVMsgProposer(nil, testingutils.PostConsensusProposerMsgV(ks.Shares[3], 3, version))),
					},
				),
				DecidedValue: cd,
				StartingDuty: &cd.Duty,
				Finished:     true,
			}
			ret.GetBaseRunner().State.RunningInstance = &qbft.Instance{
				State: &qbft.State{
					Share:             testingutils.TestingShare(ks),
					ID:                ret.GetBaseRunner().QBFTController.Identifier,
					Round:             qbft.FirstRound,
					Height:            qbft.FirstHeight,
					LastPreparedRound: qbft.NoRound,
					Decided:           true,
					DecidedValue:      cdBytes,
				},
			}
			comparable.SetMessages(ret.GetBaseRunner().State.RunningInstance, []*types.SignedSSVMessage{})
			ret.GetBaseRunner().QBFTController.StoredInstances = append(ret.GetBaseRunner().QBFTController.StoredInstances, ret.GetBaseRunner().State.RunningInstance)

			return ret
		}(),
	}
}
