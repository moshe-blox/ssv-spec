package consensusdata

import (
	comparable2 "github.com/bloxapp/ssv-spec/types/testingutils/comparable"
	reflect2 "reflect"
	"testing"
)

type SpecTest struct {
	Name string
	Data []byte
}

func (test *SpecTest) TestName() string {
	return test.Name
}

func (test *ConsensusDataTest) Run(t *testing.T) {

	err := test.ConsensusData.Validate()

	if len(test.ExpectedError) != 0 {
		require.EqualError(t, err, test.ExpectedError)
	} else {
		require.NoError(t, err)
	}

	comparable2.CompareWithJson(t, test, test.TestName(), reflect2.TypeOf(test).String())
}
