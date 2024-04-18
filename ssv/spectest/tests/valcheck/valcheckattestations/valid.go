package valcheckattestations

import (
	"github.com/mosheblox/ssv-spec/ssv/spectest/tests"
	"github.com/mosheblox/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/mosheblox/ssv-spec/types/testingutils"
)

// Valid tests valid data
func Valid() tests.SpecTest {
	return &valcheck.SpecTest{
		Name:       "attestation value check valid",
		Network:    types.PraterNetwork,
		BeaconRole: types.BNRoleAttester,
		Input:      testingutils.TestAttesterConsensusDataByts,
	}
}
