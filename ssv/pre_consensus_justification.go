package ssv

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/pkg/errors"
)

// correctQBFTState returns true if QBFT controller state requires pre-consensus justification
func (b *BaseRunner) correctQBFTState(msg *qbft.Message) bool {
	inst := b.QBFTController.InstanceForHeight(b.QBFTController.Height)
	decidedInstance := inst != nil && inst.State != nil && inst.State.Decided

	// firstHeightNotDecided is true if height == 0 (special case) and did not start yet
	firstHeightNotDecided := inst == nil && b.QBFTController.Height == msg.Height && msg.Height == qbft.FirstHeight

	// notFirstHeightDecided returns true if height != 0, height decided and the message is for next height
	notFirstHeightDecided := decidedInstance && msg.Height > qbft.FirstHeight && b.QBFTController.Height+1 == msg.Height

	return firstHeightNotDecided || notFirstHeightDecided
}

// shouldProcessingJustificationsForHeight returns true if pre-consensus justification should be processed, false otherwise
func (b *BaseRunner) shouldProcessingJustificationsForHeight(signedMsg *types.SignedSSVMessage) (bool, error) {

	msg, err := qbft.DecodeMessage(signedMsg.SSVMessage.Data)
	if err != nil {
		return false, err
	}

	correctMsgTYpe := msg.MsgType == qbft.ProposalMsgType || msg.MsgType == qbft.RoundChangeMsgType
	correctBeaconRole := b.RunnerRoleType == RoleProposer || b.RunnerRoleType == RoleAggregator || b.
		RunnerRoleType == RoleSyncCommitteeContribution
	return b.correctQBFTState(msg) && correctMsgTYpe && correctBeaconRole, nil
}

// validatePreConsensusJustifications returns an error if pre-consensus justification is invalid, nil otherwise
func (b *BaseRunner) validatePreConsensusJustifications(data *types.ConsensusData, highestDecidedDutySlot phase0.Slot) error {
	//test invalid consensus data
	if err := data.Validate(); err != nil {
		return err
	}

	if b.RunnerRoleType != types.MapDutyToRunnerRole(data.Duty.Type) {
		return errors.New("wrong beacon role")
	}

	if data.Duty.Slot <= highestDecidedDutySlot {
		return errors.New("duty.slot <= highest decided slot")
	}

	// validate justification quorum
	if !b.Share[data.Duty.ValidatorIndex].HasQuorum(len(data.PreConsensusJustifications)) {
		return errors.New("no quorum")
	}

	signers := make(map[types.OperatorID]bool)
	roots := make(map[[32]byte]bool)
	rootCount := 0
	partialSigContainer := NewPartialSigContainer(b.Share[data.Duty.ValidatorIndex].Quorum)
	for i, msg := range data.PreConsensusJustifications {
		if err := msg.Validate(); err != nil {
			return err
		}

		signer := msg.Messages[0].Signer

		// check unique signers
		if !signers[signer] {
			signers[signer] = true
		} else {
			return errors.New("duplicate signer")
		}

		// verify all justifications have the same root count
		if i == 0 {
			rootCount = len(msg.Messages)
		} else {
			if rootCount != len(msg.Messages) {
				return errors.New("inconsistent root count")
			}
		}

		// validate roots
		for _, partialSigMessage := range msg.Messages {
			// validate roots
			if i == 0 {
				// check signer did not sign duplicate root
				if roots[partialSigMessage.SigningRoot] {
					return errors.New("duplicate signed root")
				}

				// record roots
				roots[partialSigMessage.SigningRoot] = true
			} else {
				// compare roots
				if !roots[partialSigMessage.SigningRoot] {
					return errors.New("inconsistent roots")
				}
			}
			partialSigContainer.AddSignature(partialSigMessage)
		}

		// verify duty.slot == msg.slot
		if err := b.validatePartialSigMsgForSlot(msg, data.Duty.Slot); err != nil {
			return err
		}
	}

	// Verify the reconstructed signature for each root
	for root := range roots {
		_, err := b.State.ReconstructBeaconSig(partialSigContainer, root, b.Share[data.Duty.ValidatorIndex].ValidatorPubKey[:], r.GetShare().ValidatorIndex)
		if err != nil {
			return errors.Wrap(err, "wrong pre-consensus partial signature")
		}
	}

	return nil
}

// processPreConsensusJustification processes pre-consensus justification
// highestDecidedDutySlot is the highest decided duty slot known
// is the qbft message carrying  the pre-consensus justification
/** Flow:
1) needs to process justifications
2) validate data
3) validate message
4) if no running instance, run instance with consensus data duty
5) add pre-consensus sigs to container
6) decided on duty
*/
func (b *BaseRunner) processPreConsensusJustification(runner Runner, highestDecidedDutySlot phase0.Slot, msg *types.SignedSSVMessage) error {

	shouldProcess, err := b.shouldProcessingJustificationsForHeight(msg)
	if err != nil {
		return err
	}
	if !shouldProcess {
		return nil
	}

	cd := &types.ConsensusData{}
	if err := cd.Decode(msg.FullData); err != nil {
		return errors.Wrap(err, "could not decoded ConsensusData")
	}

	if err := b.validatePreConsensusJustifications(cd, highestDecidedDutySlot); err != nil {
		return err
	}

	// if no duty is running start one
	if !b.hasRunningDuty() {
		b.baseSetupForNewDuty(&cd.Duty)
	}

	// add pre-consensus sigs to state container
	var r [][32]byte
	for _, signedMsg := range cd.PreConsensusJustifications {
		quorum, roots, err := b.basePartialSigMsgProcessing(signedMsg, b.State.PreConsensusContainer)
		if err != nil {
			return errors.Wrap(err, "invalid partial sig processing")
		}

		if quorum {
			r = roots
			break
		}
	}
	if len(r) == 0 {
		return errors.New("invalid pre-consensus justification quorum")
	}

	return b.decide(runner, cd)
}
