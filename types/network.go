package types

import (
	"encoding/json"
	"math/big"
	"time"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
)

// NetworkConfig describes a network.
type NetworkConfig struct {
	Name   string
	SSV    SSVConfig
	Beacon BeaconConfig
}

type SSVConfig struct {
	Domain                 DomainType
	ForkVersion            [4]byte
	GenesisEpoch           spec.Epoch
	DefaultSyncOffset      *big.Int // prod contract genesis block
	DepositContractAddress string
	GenesisValidatorsRoot  string
	Bootnodes              []string
}

type BeaconConfig struct {
	Network          string
	SlotDuration     time.Duration
	SlotsPerEpoch    uint64
	MinGenesisTime   uint64
	CapellaForkEpoch spec.Epoch
}

var TestNetwork = NetworkConfig{
	Name: "now_test_network",
	SSV: SSVConfig{
		DefaultSyncOffset:      new(big.Int).SetInt64(8661727),
		ForkVersion:            [4]byte{0x99, 0x99, 0x99, 0x99},
		Domain:                 V3Testnet,
		DepositContractAddress: "0xff50ed3d0ec03ac01d4c79aad74928bff48a7b2b",
		GenesisValidatorsRoot:  "043db0d9a83813551ee2f33450d23797757d430911a9320530ad8a0eabc43efb",
		GenesisEpoch:           152834, // TODO: another value?
	},
	Beacon: BeaconConfig{
		Network:          "prater",
		MinGenesisTime:   1616508000,
		SlotDuration:     12 * time.Second,
		SlotsPerEpoch:    32,
		CapellaForkEpoch: 162304, // Goerli taken from https://github.com/ethereum/execution-specs/blob/37a8f892341eb000e56e962a051a87e05a2e4443/network-upgrades/mainnet-upgrades/shanghai.md?plain=1#L18
	},
}

// ForkVersion returns the fork version of the network.
func (n NetworkConfig) ForkVersion() [4]byte {
	return n.SSV.ForkVersion
}

// MinGenesisTime returns min genesis time value
func (n NetworkConfig) MinGenesisTime() uint64 {
	return n.Beacon.MinGenesisTime
}

// SlotDuration returns slot duration
func (n NetworkConfig) SlotDuration() time.Duration {
	return n.Beacon.SlotDuration
}

// SlotsPerEpoch returns number of slots per one epoch
func (n NetworkConfig) SlotsPerEpoch() uint64 {
	return n.Beacon.SlotsPerEpoch
}

// EstimatedCurrentSlot returns the estimation of the current slot
func (n NetworkConfig) EstimatedCurrentSlot() spec.Slot {
	return n.EstimatedSlotAtTime(time.Now().Unix())
}

// EstimatedSlotAtTime estimates slot at the given time
func (n NetworkConfig) EstimatedSlotAtTime(time int64) spec.Slot {
	genesis := int64(n.MinGenesisTime())
	if time < genesis {
		return 0
	}
	return spec.Slot(uint64(time-genesis) / uint64(n.SlotDuration().Seconds()))
}

func (n NetworkConfig) EstimatedTimeAtSlot(slot spec.Slot) int64 {
	d := int64(slot) * int64(n.SlotDuration().Seconds())
	return int64(n.MinGenesisTime()) + d
}

// EstimatedCurrentEpoch estimates the current epoch
// https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#compute_start_slot_at_epoch
func (n NetworkConfig) EstimatedCurrentEpoch() spec.Epoch {
	return n.EstimatedEpochAtSlot(n.EstimatedCurrentSlot())
}

// EstimatedEpochAtSlot estimates epoch at the given slot
func (n NetworkConfig) EstimatedEpochAtSlot(slot spec.Slot) spec.Epoch {
	return spec.Epoch(slot / spec.Slot(n.SlotsPerEpoch()))
}

func (n NetworkConfig) FirstSlotAtEpoch(epoch spec.Epoch) spec.Slot {
	return spec.Slot(uint64(epoch) * n.SlotsPerEpoch())
}

func (n NetworkConfig) EpochStartTime(epoch spec.Epoch) time.Time {
	firstSlot := n.FirstSlotAtEpoch(epoch)
	t := n.EstimatedTimeAtSlot(firstSlot)
	return time.Unix(t, 0)
}

func (n NetworkConfig) String() string {
	b, err := json.MarshalIndent(n, "", "\t")
	if err != nil {
		return "<malformed>"
	}

	return string(b)
}
