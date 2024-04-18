package testingutils

import (
	"github.com/mosheblox/ssv-spec/ssv"
)

var BaseCluster = func(keySet *TestKeySet) *ssv.Committee {

	return ssv.NewCommittee(
		NewTestingNetwork(),
		NewTestingBeaconNode(),
		NewTestingKeyManager(),
		func() *ssv.ClusterRunner {
			return &ssv.ClusterRunner{}
		},
	)
}
