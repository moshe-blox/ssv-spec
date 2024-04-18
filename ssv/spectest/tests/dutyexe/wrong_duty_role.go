package dutyexe

import (
	"crypto/rsa"
	"fmt"

	"github.com/attestantio/go-eth2-client/spec"
	"github.com/moshe-blox/ssv-spec/ssv/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// WrongDutyRole tests decided value duty with wrong duty role (!= duty runner role)
func WrongDutyRole() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	// Correct ID for SSVMessage
	getID := func(role types.BeaconRole) types.MessageID {
		ret := types.NewMsgID(testingutils.TestingSSVDomainType, testingutils.TestingValidatorPubKey[:], role)
		return ret
	}
	// Wrong ID for SignedMessage
	getWrongID := func(role types.BeaconRole) []byte {
		ret := types.NewMsgID(testingutils.TestingSSVDomainType, testingutils.TestingValidatorPubKey[:], role+1)
		return ret[:]
	}

	// Function to get decided message with wrong ID for role
	decidedMessage := func(role types.BeaconRole) *types.SignedSSVMessage {
		signedMessage := testingutils.TestingCommitMultiSignerMessageWithHeightAndIdentifier(
			[]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[2], ks.OperatorKeys[3]},
			[]types.OperatorID{1, 2, 3},
			testingutils.TestingDutySlot,
			getWrongID(role))

		signedMessage.SSVMessage.MsgID = getID(role)

		sig1 := testingutils.SignedSSVMessageWithSigner(1, ks.OperatorKeys[1], signedMessage.SSVMessage).Signatures[0]
		sig2 := testingutils.SignedSSVMessageWithSigner(2, ks.OperatorKeys[2], signedMessage.SSVMessage).Signatures[0]
		sig3 := testingutils.SignedSSVMessageWithSigner(3, ks.OperatorKeys[3], signedMessage.SSVMessage).Signatures[0]

		signedMessage.Signatures = [][]byte{sig1, sig2, sig3}

		return signedMessage
	}

	expectedError := "failed processing consensus message: invalid msg: message doesn't belong to Identifier"

	multiSpecTest := &tests.MultiMsgProcessingSpecTest{
		Name: "wrong duty role",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:     "sync committee contribution",
				Runner:   testingutils.SyncCommitteeContributionRunner(ks),
				Duty:     &testingutils.TestingSyncCommitteeContributionDuty,
				Messages: []*types.SignedSSVMessage{decidedMessage(types.BNRoleSyncCommitteeContribution)},
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusContributionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
				},
				ExpectedError: expectedError,
			},
			{
				Name:           "sync committee",
				Runner:         testingutils.SyncCommitteeRunner(ks),
				Duty:           &testingutils.TestingSyncCommitteeDuty,
				Messages:       []*types.SignedSSVMessage{decidedMessage(types.BNRoleSyncCommittee)},
				OutputMessages: []*types.PartialSignatureMessages{},
				ExpectedError:  expectedError,
			},
			{
				Name:     "aggregator",
				Runner:   testingutils.AggregatorRunner(ks),
				Duty:     &testingutils.TestingAggregatorDuty,
				Messages: []*types.SignedSSVMessage{decidedMessage(types.BNRoleAggregator)},
				OutputMessages: []*types.PartialSignatureMessages{
					testingutils.PreConsensusSelectionProofMsg(ks.Shares[1], ks.Shares[1], 1, 1),
				},
				ExpectedError: expectedError,
			},
			{
				Name:           "attester",
				Runner:         testingutils.CommitteeRunner(ks),
				Duty:           &testingutils.TestingAttesterDuty,
				Messages:       []*types.SignedSSVMessage{decidedMessage(types.BNRoleAttester)},
				OutputMessages: []*types.PartialSignatureMessages{},
				ExpectedError:  expectedError,
			},
		},
	}

	// proposerV creates a test specification for versioned proposer.
	proposerV := func(version spec.DataVersion) *tests.MsgProcessingSpecTest {
		return &tests.MsgProcessingSpecTest{
			Name:     fmt.Sprintf("proposer (%s)", version.String()),
			Runner:   testingutils.ProposerRunner(ks),
			Duty:     testingutils.TestingProposerDutyV(version),
			Messages: []*types.SignedSSVMessage{decidedMessage(types.BNRoleProposer)},
			OutputMessages: []*types.PartialSignatureMessages{
				testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, version),
			},
			ExpectedError: expectedError,
		}
	}

	// proposerBlindedV creates a test specification for versioned proposer with blinded block.
	proposerBlindedV := func(version spec.DataVersion) *tests.MsgProcessingSpecTest {
		return &tests.MsgProcessingSpecTest{
			Name:     fmt.Sprintf("proposer blinded block (%s)", version.String()),
			Runner:   testingutils.ProposerBlindedBlockRunner(ks),
			Duty:     testingutils.TestingProposerDutyV(version),
			Messages: []*types.SignedSSVMessage{decidedMessage(types.BNRoleProposer)},
			OutputMessages: []*types.PartialSignatureMessages{
				testingutils.PreConsensusRandaoMsgV(ks.Shares[1], 1, version),
			},
			ExpectedError: expectedError,
		}
	}

	for _, v := range testingutils.SupportedBlockVersions {
		multiSpecTest.Tests = append(multiSpecTest.Tests, []*tests.MsgProcessingSpecTest{proposerV(v), proposerBlindedV(v)}...)
	}

	return multiSpecTest
}
