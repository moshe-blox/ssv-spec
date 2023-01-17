package testingutils

import (
	"encoding/hex"
	altair "github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/types"
	ssz "github.com/ferranbt/fastssz"
	"github.com/prysmaticlabs/go-bitfield"
)

var signBeaconObject = func(
	obj ssz.HashRoot,
	domainType spec.DomainType,
	ks *TestKeySet,
) spec.BLSSignature {
	domain, _ := NewTestingBeaconNode().DomainData(1, domainType)
	ret, _, _ := NewTestingKeyManager().SignBeaconObject(obj, domain, ks.ValidatorPK.Serialize(), domainType)

	blsSig := spec.BLSSignature{}
	copy(blsSig[:], ret)

	return blsSig
}

var TestingAttestationData = &spec.AttestationData{
	Slot:            12,
	Index:           3,
	BeaconBlockRoot: spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	Source: &spec.Checkpoint{
		Epoch: 0,
		Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
	Target: &spec.Checkpoint{
		Epoch: 1,
		Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
}
var TestingWrongAttestationData = func() *spec.AttestationData {
	byts, _ := TestingAttestationData.MarshalSSZ()
	ret := &spec.AttestationData{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.Slot = 100
	return ret
}()

var TestingSignedAttestation = func(ks *TestKeySet) *spec.Attestation {
	aggregationBitfield := bitfield.NewBitlist(TestingAttesterDuty.CommitteeLength)
	aggregationBitfield.SetBitAt(TestingAttesterDuty.ValidatorCommitteeIndex, true)
	return &spec.Attestation{
		Data:            TestingAttestationData,
		Signature:       signBeaconObject(TestingAttestationData, types.DomainAttester, ks),
		AggregationBits: aggregationBitfield,
	}
}

var TestingBeaconBlock = &bellatrix.BeaconBlock{
	Slot:          12,
	ProposerIndex: 10,
	ParentRoot:    spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	StateRoot:     spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	Body: &bellatrix.BeaconBlockBody{
		RANDAOReveal: spec.BLSSignature{},
		ETH1Data: &spec.ETH1Data{
			DepositRoot:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			DepositCount: 100,
			BlockHash:    []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
		},
		Graffiti:          []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
		ProposerSlashings: []*spec.ProposerSlashing{},
		AttesterSlashings: []*spec.AttesterSlashing{},
		Attestations: []*spec.Attestation{
			{
				AggregationBits: bitfield.NewBitlist(122),
				Data:            TestingAttestationData,
				Signature:       spec.BLSSignature{},
			},
		},
		Deposits:       []*spec.Deposit{},
		VoluntaryExits: []*spec.SignedVoluntaryExit{},
		SyncAggregate: &altair.SyncAggregate{
			SyncCommitteeBits:      bitfield.NewBitvector512(),
			SyncCommitteeSignature: spec.BLSSignature{},
		},
		ExecutionPayload: &bellatrix.ExecutionPayload{
			ParentHash:    spec.Hash32{},
			FeeRecipient:  bellatrix.ExecutionAddress{},
			StateRoot:     spec.Hash32{},
			ReceiptsRoot:  spec.Hash32{},
			LogsBloom:     [256]byte{},
			PrevRandao:    [32]byte{},
			BlockNumber:   100,
			GasLimit:      1000000,
			GasUsed:       800000,
			Timestamp:     123456789,
			BaseFeePerGas: [32]byte{},
			BlockHash:     spec.Hash32{},
			Transactions:  []bellatrix.Transaction{},
		},
	},
}
var TestingWrongBeaconBlock = func() *bellatrix.BeaconBlock {
	byts, err := TestingBeaconBlock.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}
	ret := &bellatrix.BeaconBlock{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.Slot = 100
	return ret
}()

var TestingSignedBeaconBlock = func(ks *TestKeySet) *bellatrix.SignedBeaconBlock {
	return &bellatrix.SignedBeaconBlock{
		Message:   TestingBeaconBlock,
		Signature: signBeaconObject(TestingBeaconBlock, types.DomainProposer, ks),
	}
}

var TestingAggregateAndProof = &spec.AggregateAndProof{
	AggregatorIndex: 1,
	SelectionProof:  spec.BLSSignature{},
	Aggregate: &spec.Attestation{
		AggregationBits: bitfield.NewBitlist(128),
		Signature:       spec.BLSSignature{},
		Data:            TestingAttestationData,
	},
}
var TestingWrongAggregateAndProof = func() *spec.AggregateAndProof {
	byts, err := TestingAggregateAndProof.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}
	ret := &spec.AggregateAndProof{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.AggregatorIndex = 100
	return ret
}()

var TestingSignedAggregateAndProof = func(ks *TestKeySet) *spec.SignedAggregateAndProof {
	return &spec.SignedAggregateAndProof{
		Message:   TestingAggregateAndProof,
		Signature: signBeaconObject(TestingAggregateAndProof, types.DomainAggregateAndProof, ks),
	}
}

const (
	TestingDutySlot       = 12
	TestingDutySlot2      = 50
	TestingDutyEpoch      = 0
	TestingDutyEpoch2     = 1
	TestingValidatorIndex = 1

	UnknownDutyType = 100
)

var TestingSyncCommitteeBlockRoot = spec.Root{}
var TestingSyncCommitteeWrongBlockRoot = spec.Root{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var TestingSignedSyncCommitteeBlockRoot = func(ks *TestKeySet) *altair.SyncCommitteeMessage {
	return &altair.SyncCommitteeMessage{
		Slot:            TestingDutySlot,
		BeaconBlockRoot: TestingSyncCommitteeBlockRoot,
		ValidatorIndex:  TestingValidatorIndex,
		Signature:       signBeaconObject(types.SSZBytes(TestingSyncCommitteeBlockRoot[:]), types.DomainSyncCommittee, ks),
	}
}

var TestingContributionProofIndexes = []spec.CommitteeIndex{0, 1, 2}
var TestingContributionProofsSigned = func() []spec.BLSSignature {
	// signed with 3515c7d08e5affd729e9579f7588d30f2342ee6f6a9334acf006345262162c6f
	byts1, _ := hex.DecodeString("b18833bb7549ec33e8ac414ba002fd45bb094ca300bd24596f04a434a89beea462401da7c6b92fb3991bd17163eb603604a40e8dd6781266c990023446776ff42a9313df26a0a34184a590e57fa4003d610c2fa214db4e7dec468592010298bc")
	byts2, _ := hex.DecodeString("9094342c95146554df849dc20f7425fca692dacee7cb45258ddd264a8e5929861469fda3d1567b9521cba83188ffd61a0dbe6d7180c7a96f5810d18db305e9143772b766d368aa96d3751f98d0ce2db9f9e6f26325702088d87f0de500c67c68")
	byts3, _ := hex.DecodeString("a7f88ce43eff3aa8cdd2e3957c5bead4e21353fbecac6079a5398d03019bc45ff7c951785172deee70e9bc5abbc8ca6a0f0441e9d4cc9da74c31121357f7d7c7de9533f6f457da493e3314e22d554ab76613e469b050e246aff539a33807197c")

	ret := make([]spec.BLSSignature, 0)
	for _, byts := range [][]byte{byts1, byts2, byts3} {
		b := spec.BLSSignature{}
		copy(b[:], byts)
		ret = append(ret, b)
	}
	return ret
}()

var TestingSyncCommitteeContributions = []*altair.SyncCommitteeContribution{
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 0,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 1,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 2,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
}

var TestingSignedSyncCommitteeContributions = func(
	contrib *altair.SyncCommitteeContribution,
	proof spec.BLSSignature,
	ks *TestKeySet) *altair.SignedContributionAndProof {
	msg := &altair.ContributionAndProof{
		AggregatorIndex: TestingValidatorIndex,
		Contribution:    contrib,
		SelectionProof:  proof,
	}
	return &altair.SignedContributionAndProof{
		Message:   msg,
		Signature: signBeaconObject(msg, types.DomainContributionAndProof, ks),
	}
}

var TestingAttesterDuty = &types.Duty{
	Type:                    types.BNRoleAttester,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingProposerDuty = &types.Duty{
	Type:                    types.BNRoleProposer,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

// TestingProposerDutyNextEpoch testing for a second duty start
var TestingProposerDutyNextEpoch = &types.Duty{
	Type:                    types.BNRoleProposer,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot2,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingAggregatorDuty = &types.Duty{
	Type:                    types.BNRoleAggregator,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

// TestingAggregatorDutyNextEpoch testing for a second duty start
var TestingAggregatorDutyNextEpoch = &types.Duty{
	Type:                    types.BNRoleAggregator,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot2,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingSyncCommitteeDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommittee,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

var TestingSyncCommitteeContributionDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommitteeContribution,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

// TestingSyncCommitteeContributionNexEpochDuty testing for a second duty start
var TestingSyncCommitteeContributionNexEpochDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommitteeContribution,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot2,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

var TestingUnknownDutyType = &types.Duty{
	Type:                    UnknownDutyType,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    12,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingWrongDutyPK = &types.Duty{
	Type:                    types.BNRoleAttester,
	PubKey:                  TestingWrongValidatorPubKey,
	Slot:                    12,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

//func blsSigFromHex(str string) spec.BLSSignature {
//	byts, _ := hex.DecodeString(str)
//	ret := spec.BLSSignature{}
//	copy(ret[:], byts)
//	return ret
//}

type TestingBeaconNode struct {
	BroadcastedRoots             []spec.Root
	syncCommitteeAggregatorRoots map[string]bool
}

func NewTestingBeaconNode() *TestingBeaconNode {
	return &TestingBeaconNode{
		BroadcastedRoots: []spec.Root{},
	}
}

// SetSyncCommitteeAggregatorRootHexes FOR TESTING ONLY!! sets which sync committee aggregator roots will return true for aggregator
func (bn *TestingBeaconNode) SetSyncCommitteeAggregatorRootHexes(roots map[string]bool) {
	bn.syncCommitteeAggregatorRoots = roots
}

// GetBeaconNetwork returns the beacon network the node is on
func (bn *TestingBeaconNode) GetBeaconNetwork() types.BeaconNetwork {
	return types.NowTestNetwork
}

// GetAttestationData returns attestation data by the given slot and committee index
func (bn *TestingBeaconNode) GetAttestationData(slot spec.Slot, committeeIndex spec.CommitteeIndex) (*spec.AttestationData, error) {
	return TestingAttestationData, nil
}

// SubmitAttestation submit the attestation to the node
func (bn *TestingBeaconNode) SubmitAttestation(attestation *spec.Attestation) error {
	r, _ := attestation.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// GetBeaconBlock returns beacon block by the given slot and committee index
func (bn *TestingBeaconNode) GetBeaconBlock(slot spec.Slot, graffiti, randao []byte) (*bellatrix.BeaconBlock, error) {
	return TestingBeaconBlock, nil
}

// SubmitBeaconBlock submit the block to the node
func (bn *TestingBeaconNode) SubmitBeaconBlock(block *bellatrix.SignedBeaconBlock) error {
	r, _ := block.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// SubmitAggregateSelectionProof returns an AggregateAndProof object
func (bn *TestingBeaconNode) SubmitAggregateSelectionProof(slot spec.Slot, committeeIndex spec.CommitteeIndex, committeeLength uint64, index spec.ValidatorIndex, slotSig []byte) (*spec.AggregateAndProof, error) {
	return TestingAggregateAndProof, nil
}

// SubmitSignedAggregateSelectionProof broadcasts a signed aggregator msg
func (bn *TestingBeaconNode) SubmitSignedAggregateSelectionProof(msg *spec.SignedAggregateAndProof) error {
	r, _ := msg.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// GetSyncMessageBlockRoot returns beacon block root for sync committee
func (bn *TestingBeaconNode) GetSyncMessageBlockRoot(slot spec.Slot) (spec.Root, error) {
	return TestingSyncCommitteeBlockRoot, nil
}

// SubmitSyncMessage submits a signed sync committee msg
func (bn *TestingBeaconNode) SubmitSyncMessage(msg *altair.SyncCommitteeMessage) error {
	r, _ := msg.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// IsSyncCommitteeAggregator returns tru if aggregator
func (bn *TestingBeaconNode) IsSyncCommitteeAggregator(proof []byte) (bool, error) {
	if len(bn.syncCommitteeAggregatorRoots) != 0 {
		if val, found := bn.syncCommitteeAggregatorRoots[hex.EncodeToString(proof)]; found {
			return val, nil
		}
		return false, nil
	}
	return true, nil
}

// SyncCommitteeSubnetID returns sync committee subnet ID from subcommittee index
func (bn *TestingBeaconNode) SyncCommitteeSubnetID(index spec.CommitteeIndex) (uint64, error) {
	// each subcommittee index correlates to TestingContributionProofRoots by index
	return uint64(index), nil
}

// GetSyncCommitteeContribution returns
func (bn *TestingBeaconNode) GetSyncCommitteeContribution(slot spec.Slot, subnetID uint64) (*altair.SyncCommitteeContribution, error) {
	return TestingSyncCommitteeContributions[subnetID], nil
}

// SubmitSignedContributionAndProof broadcasts to the network
func (bn *TestingBeaconNode) SubmitSignedContributionAndProof(contribution *altair.SignedContributionAndProof) error {
	r, _ := contribution.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

func (bn *TestingBeaconNode) DomainData(epoch spec.Epoch, domain spec.DomainType) (spec.Domain, error) {
	// epoch is used to calculate fork version, here we hard code it
	return types.ComputeETHDomain(domain, types.GenesisForkVersion, types.GenesisValidatorsRoot)
}
