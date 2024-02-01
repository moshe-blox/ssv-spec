package share

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

// Returns the number of unique signers in the message signers list
func (test *ShareTest) GetUniqueMessageSignersCount() int {
	uniqueSigners := make(map[uint64]bool)

	for _, element := range test.Message.Signers {
		uniqueSigners[element] = true
	}

	return len(uniqueSigners)
}

func (test *ShareTest) Run(t *testing.T) {

	// Validate message
	err := test.Message.Validate()
	if len(test.ExpectedError) != 0 {
		require.EqualError(t, err, test.ExpectedError)
	} else {
		require.NoError(t, err)
	}

	// Get unique signers
	numSigners := test.GetUniqueMessageSignersCount()

	// Test expected thresholds results
	require.Equal(t, test.ExpectedHasPartialQuorum, test.Share.HasPartialQuorum(numSigners))
	require.Equal(t, test.ExpectedHasQuorum, test.Share.HasQuorum(numSigners))
	require.Equal(t, test.ExpectedFullCommittee, (len(test.Share.Committee) == numSigners))

	comparable2.CompareWithJson(t, test, test.TestName(), reflect2.TypeOf(test).String())
}
