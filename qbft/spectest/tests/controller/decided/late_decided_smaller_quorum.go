package decided

import (
	"crypto/rsa"

	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// LateDecidedSmallerQuorum tests processing a decided msg for a just decided instance (with a smaller quorum)
func LateDecidedSmallerQuorum() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	return &tests.ControllerSpecTest{
		Name: "decide late decided smaller quorum",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue: []byte{1, 2, 3, 4},
				InputMessages: []*types.SignedSSVMessage{
					testingutils.TestingCommitMultiSignerMessage([]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[2], ks.OperatorKeys[3], ks.OperatorKeys[4]}, []types.OperatorID{1, 2, 3, 4}),
					testingutils.TestingCommitMultiSignerMessage([]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[2], ks.OperatorKeys[3]}, []types.OperatorID{1, 2, 3}),
				},
				ExpectedDecidedState: tests.DecidedState{
					DecidedCnt: 1,
					DecidedVal: testingutils.TestingQBFTFullData,
				},
				ControllerPostRoot: "c8ea9638aa8aa332763f5b79fe26d88bbcb173837b64545d871eec89df6b7ee5",
			},
		},
	}
}
