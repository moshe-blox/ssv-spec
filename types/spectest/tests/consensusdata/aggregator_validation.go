package consensusdata

import "github.com/moshe-blox/ssv-spec/types/testingutils"

// AggregatorValidation tests a valid consensus data with AggregateAndProof
func AggregatorValidation() *ConsensusDataTest {
	ks := testingutils.Testing4SharesSet()

	return &ConsensusDataTest{
		Name:          "aggregator valid",
		ConsensusData: *testingutils.TestAggregatorWithJustificationsConsensusData(ks),
	}
}
