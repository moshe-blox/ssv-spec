package validation

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mosheblox/ssv-spec/qbft"
	"github.com/mosheblox/ssv-spec/ssv"
	"github.com/mosheblox/ssv-spec/types"
	"github.com/pkg/errors"
)

// MsgValidatorFunc represents a message validator
type MsgValidatorFunc = func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult

func MsgValidation(runner ssv.Runner) MsgValidatorFunc {
	return func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
		signedSSVMsg, err := DecodePubsubMsg(msg)
		if err != nil {
			return pubsub.ValidationReject
		}

		switch signedSSVMsg.SSVMessage.GetType() {
		case types.SSVConsensusMsgType:
			if validateConsensusMsg(runner, signedSSVMsg) != nil {
				return pubsub.ValidationReject
			}
		case types.SSVPartialSignatureMsgType:
			if validatePartialSigMsg(runner, signedSSVMsg.SSVMessage.Data) != nil {
				return pubsub.ValidationReject
			}
		default:
			return pubsub.ValidationReject
		}

		return pubsub.ValidationAccept
	}
}

func DecodePubsubMsg(msg *pubsub.Message) (*types.SignedSSVMessage, error) {
	byts := msg.GetData()
	ret := &types.SignedSSVMessage{}
	if err := ret.Decode(byts); err != nil {
		return nil, err
	}
	return ret, nil
}

func validateConsensusMsg(runner ssv.Runner, signedMsg *types.SignedSSVMessage) error {

	contr := runner.GetBaseRunner().QBFTController

	if err := contr.BaseMsgValidation(signedMsg); err != nil {
		return err
	}

	/**
	Main controller processing flow
	_______________________________
	All decided msgs are processed the same, out of instance
	All valid future msgs are saved in a container and can trigger highest decided futuremsg
	All other msgs (not future or decided) are processed normally by an existing instance (if found)
	*/
	isDecided, err := qbft.IsDecidedMsg(contr.Share, signedMsg)
	if err != nil {
		return err
	}
	if isDecided {
		return qbft.ValidateDecided(contr.GetConfig(), signedMsg, contr.Share)
	}

	msg, err := qbft.DecodeMessage(signedMsg.SSVMessage.Data)
	if err != nil {
		return err
	}

	if msg.Height > contr.Height {
		return validateFutureMsg(contr.GetConfig(), signedMsg, contr.Share.Committee)
	}

	if inst := contr.StoredInstances.FindInstance(msg.Height); inst != nil {
		return inst.BaseMsgValidation(signedMsg)
	}
	return errors.New("unknown instance")
}

func validatePartialSigMsg(runner ssv.Runner, data []byte) error {
	signedMsg := &types.PartialSignatureMessages{}
	if err := signedMsg.Decode(data); err != nil {
		return err
	}

	if signedMsg.Type == types.PostConsensusPartialSig {
		return runner.GetBaseRunner().ValidatePostConsensusMsg(runner, signedMsg)
	}
	return runner.GetBaseRunner().ValidatePreConsensusMsg(runner, signedMsg)
}

func validateFutureMsg(
	config qbft.IConfig,
	msg *types.SignedSSVMessage,
	operators []*types.Operator,
) error {
	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "invalid decided msg")
	}

	if len(msg.GetOperatorIDs()) != 1 {
		return errors.New("allows 1 signer")
	}

	// verify signature
	if err := config.GetSignatureVerifier().Verify(msg, operators); err != nil {
		return errors.Wrap(err, "msg signature invalid")
	}

	return nil
}
