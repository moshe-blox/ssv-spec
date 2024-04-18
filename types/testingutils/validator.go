package testingutils

import (
	"github.com/mosheblox/ssv-spec/ssv"
	"github.com/mosheblox/ssv-spec/types"
)

var BaseValidator = func(keySet *TestKeySet) *ssv.Validator {
	return ssv.NewValidator(
		NewTestingNetwork(1, keySet.OperatorKeys[1]),
		NewTestingBeaconNode(),
		TestingShare(keySet),
		NewTestingKeyManager(),
		NewTestingOperatorSigner(keySet, 1),
		map[types.BeaconRole]ssv.Runner{
			types.BNRoleAttester:                  CommitteeRunner(keySet),
			types.BNRoleProposer:                  ProposerRunner(keySet),
			types.BNRoleAggregator:                AggregatorRunner(keySet),
			types.BNRoleSyncCommittee:             SyncCommitteeRunner(keySet),
			types.BNRoleSyncCommitteeContribution: SyncCommitteeContributionRunner(keySet),
		},
		NewTestingVerifier(),
	)
}
