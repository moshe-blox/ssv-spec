package types

import (
	"encoding/hex"

	json "github.com/bytedance/sonic"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type ContributionsMap map[phase0.BLSSignature]*altair.SyncCommitteeContribution

func (cm *ContributionsMap) MarshalJSON() ([]byte, error) {
	m := make(map[string]*altair.SyncCommitteeContribution)
	for k, v := range *cm {
		m[hex.EncodeToString(k[:])] = v
	}
	return json.ConfigFastest.Marshal(m)
}

func (cm *ContributionsMap) UnmarshalJSON(input []byte) error {
	m := make(map[string]*altair.SyncCommitteeContribution)
	if err := json.ConfigFastest.Unmarshal(input, &m); err != nil {
		return err
	}

	if *cm == nil {
		*cm = ContributionsMap{}
	}

	for k, v := range m {
		byts, err := hex.DecodeString(k)
		if err != nil {
			return err
		}

		blSig := phase0.BLSSignature{}
		copy(blSig[:], byts)
		(*cm)[blSig] = v
	}
	return nil
}

// ConsensusData holds all relevant duty and data Decided on by consensus
type ConsensusData struct {
	Duty                   *Duty
	AttestationData        *phase0.AttestationData
	BlockData              *bellatrix.BeaconBlock
	AggregateAndProof      *phase0.AggregateAndProof
	SyncCommitteeBlockRoot phase0.Root
	// SyncCommitteeContribution map holds as key the selection proof for the contribution
	SyncCommitteeContribution ContributionsMap
}

func (cid *ConsensusData) Encode() ([]byte, error) {
	return json.ConfigFastest.Marshal(cid)
}

func (cid *ConsensusData) Decode(data []byte) error {
	return json.ConfigFastest.Unmarshal(data, &cid)
}
