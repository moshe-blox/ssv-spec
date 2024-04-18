package share

import (
	reflect2 "reflect"
	"testing"

	"github.com/mosheblox/ssv-spec/types"
	comparable2 "github.com/mosheblox/ssv-spec/types/testingutils/comparable"
	"github.com/stretchr/testify/require"
)

type EncodingTest struct {
	Name         string
	Data         []byte
	ExpectedRoot [32]byte
}

func (test *EncodingTest) TestName() string {
	return test.Name
}

func (test *EncodingTest) Run(t *testing.T) {
	// decode
	decodedShare := &types.Share{}
	require.NoError(t, decodedShare.Decode(test.Data))

	// encode
	byts, err := decodedShare.Encode()
	require.NoError(t, err)
	require.EqualValues(t, test.Data, byts)

	// test root
	r, err := decodedShare.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, test.ExpectedRoot, r)

	comparable2.CompareWithJson(t, test, test.TestName(), reflect2.TypeOf(test).String())
}
