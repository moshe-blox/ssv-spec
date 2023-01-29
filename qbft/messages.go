package qbft

import (
	"bytes"
	"crypto/sha256"

	json "github.com/bytedance/sonic"

	"github.com/bloxapp/ssv-spec/types"
	"github.com/pkg/errors"
)

// HasQuorum returns true if a unique set of signers has quorum
func HasQuorum(share *types.Share, msgs []*SignedMessage) bool {
	uniqueSigners := make(map[types.OperatorID]bool)
	for _, msg := range msgs {
		for _, signer := range msg.GetSigners() {
			uniqueSigners[signer] = true
		}
	}
	return share.HasQuorum(len(uniqueSigners))
}

// HasPartialQuorum returns true if a unique set of signers has partial quorum
func HasPartialQuorum(share *types.Share, msgs []*SignedMessage) bool {
	uniqueSigners := make(map[types.OperatorID]bool)
	for _, msg := range msgs {
		for _, signer := range msg.GetSigners() {
			uniqueSigners[signer] = true
		}
	}
	return share.HasPartialQuorum(len(uniqueSigners))
}

type MessageType int

const (
	ProposalMsgType MessageType = iota
	PrepareMsgType
	CommitMsgType
	RoundChangeMsgType
)

type ProposalData struct {
	Data                     []byte
	RoundChangeJustification []*SignedMessage
	PrepareJustification     []*SignedMessage
}

// Encode returns a msg encoded bytes or error
func (d *ProposalData) Encode() ([]byte, error) {
	return json.Marshal(d)
}

// Decode returns error if decoding failed
func (d *ProposalData) Decode(data []byte) error {
	return json.Unmarshal(data, &d)
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (d *ProposalData) Validate() error {
	if len(d.Data) == 0 {
		return errors.New("ProposalData data is invalid")
	}
	return nil
}

type PrepareData struct {
	Data []byte
}

// Encode returns a msg encoded bytes or error
func (d *PrepareData) Encode() ([]byte, error) {
	return json.Marshal(d)
}

// Decode returns error if decoding failed
func (d *PrepareData) Decode(data []byte) error {
	return json.Unmarshal(data, &d)
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (d *PrepareData) Validate() error {
	if len(d.Data) == 0 {
		return errors.New("PrepareData data is invalid")
	}
	return nil
}

type CommitData struct {
	Data []byte
}

// Encode returns a msg encoded bytes or error
func (d *CommitData) Encode() ([]byte, error) {
	return json.Marshal(d)
}

// Decode returns error if decoding failed
func (d *CommitData) Decode(data []byte) error {
	return json.Unmarshal(data, &d)
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (d *CommitData) Validate() error {
	if len(d.Data) == 0 {
		return errors.New("CommitData data is invalid")
	}
	return nil
}

type RoundChangeData struct {
	PreparedValue            []byte
	PreparedRound            Round
	RoundChangeJustification []*SignedMessage
}

func (d *RoundChangeData) Prepared() bool {
	if d.PreparedRound != NoRound || len(d.PreparedValue) != 0 {
		return true
	}
	return false
}

// Encode returns a msg encoded bytes or error
func (d *RoundChangeData) Encode() ([]byte, error) {
	return json.Marshal(d)
}

// Decode returns error if decoding failed
func (d *RoundChangeData) Decode(data []byte) error {
	return json.Unmarshal(data, &d)
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (d *RoundChangeData) Validate() error {
	if d.Prepared() {
		if len(d.PreparedValue) == 0 {
			return errors.New("round change prepared value invalid")
		}
		if len(d.RoundChangeJustification) == 0 {
			return errors.New("round change justification invalid")
		}
		// TODO - should next proposal data be equal to prepared value?
	}
	return nil
}

type Message struct {
	MsgType    MessageType
	Height     Height // QBFT instance Height
	Round      Round  // QBFT round for which the msg is for
	Identifier []byte // instance Identifier this msg belongs to
	Data       []byte
}

// GetProposalData returns proposal specific data
func (msg *Message) GetProposalData() (*ProposalData, error) {
	ret := &ProposalData{}
	if err := ret.Decode(msg.Data); err != nil {
		return nil, errors.Wrap(err, "could not decode proposal data from message")
	}
	return ret, nil
}

// GetPrepareData returns prepare specific data
func (msg *Message) GetPrepareData() (*PrepareData, error) {
	ret := &PrepareData{}
	if err := ret.Decode(msg.Data); err != nil {
		return nil, errors.Wrap(err, "could not decode prepare data from message")
	}
	return ret, nil
}

// GetCommitData returns commit specific data
func (msg *Message) GetCommitData() (*CommitData, error) {
	ret := &CommitData{}
	if err := ret.Decode(msg.Data); err != nil {
		return nil, errors.Wrap(err, "could not decode commit data from message")
	}
	return ret, nil
}

// GetRoundChangeData returns round change specific data
func (msg *Message) GetRoundChangeData() (*RoundChangeData, error) {
	ret := &RoundChangeData{}
	if err := ret.Decode(msg.Data); err != nil {
		return nil, errors.Wrap(err, "could not decode round change data from message")
	}
	return ret, nil
}

// Encode returns a msg encoded bytes or error
func (msg *Message) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// Decode returns error if decoding failed
func (msg *Message) Decode(data []byte) error {
	return json.Unmarshal(data, &msg)
}

// GetRoot returns the root used for signing and verification
func (msg *Message) GetRoot() ([]byte, error) {
	marshaledRoot, err := msg.Encode()
	if err != nil {
		return nil, errors.Wrap(err, "could not encode message")
	}
	ret := sha256.Sum256(marshaledRoot)
	return ret[:], nil
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (msg *Message) Validate() error {
	if len(msg.Identifier) == 0 {
		return errors.New("message identifier is invalid")
	}
	if len(msg.Data) == 0 {
		return errors.New("message data is invalid")
	}
	if msg.MsgType > 5 {
		return errors.New("message type is invalid")
	}
	return nil
}

type SignedMessage struct {
	Signature types.Signature
	Signers   []types.OperatorID
	Message   *Message // message for which this signature is for
}

func (signedMsg *SignedMessage) GetSignature() types.Signature {
	return signedMsg.Signature
}
func (signedMsg *SignedMessage) GetSigners() []types.OperatorID {
	return signedMsg.Signers
}

// MatchedSigners returns true if the provided signer ids are equal to GetSignerIds() without order significance
func (signedMsg *SignedMessage) MatchedSigners(ids []types.OperatorID) bool {
	if len(signedMsg.Signers) != len(ids) {
		return false
	}

	for _, id := range signedMsg.Signers {
		found := false
		for _, id2 := range ids {
			if id == id2 {
				found = true
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// CommonSigners returns true if there is at least 1 common signer
func (signedMsg *SignedMessage) CommonSigners(ids []types.OperatorID) bool {
	for _, id := range signedMsg.Signers {
		for _, id2 := range ids {
			if id == id2 {
				return true
			}
		}
	}
	return false
}

// Aggregate will aggregate the signed message if possible (unique signers, same digest, valid)
func (signedMsg *SignedMessage) Aggregate(sig types.MessageSignature) error {
	if signedMsg.CommonSigners(sig.GetSigners()) {
		return errors.New("duplicate signers")
	}

	r1, err := signedMsg.GetRoot()
	if err != nil {
		return errors.Wrap(err, "could not get signature root")
	}
	r2, err := sig.GetRoot()
	if err != nil {
		return errors.Wrap(err, "could not get signature root")
	}
	if !bytes.Equal(r1, r2) {
		return errors.New("can't aggregate, roots not equal")
	}

	aggregated, err := signedMsg.Signature.Aggregate(sig.GetSignature())
	if err != nil {
		return errors.Wrap(err, "could not aggregate signatures")
	}
	signedMsg.Signature = aggregated
	signedMsg.Signers = append(signedMsg.Signers, sig.GetSigners()...)

	return nil
}

// Encode returns a msg encoded bytes or error
func (signedMsg *SignedMessage) Encode() ([]byte, error) {
	return json.Marshal(signedMsg)
}

// Decode returns error if decoding failed
func (signedMsg *SignedMessage) Decode(data []byte) error {
	return json.Unmarshal(data, &signedMsg)
}

// GetRoot returns the root used for signing and verification
func (signedMsg *SignedMessage) GetRoot() ([]byte, error) {
	return signedMsg.Message.GetRoot()
}

// DeepCopy returns a new instance of SignedMessage, deep copied
func (signedMsg *SignedMessage) DeepCopy() *SignedMessage {
	ret := &SignedMessage{
		Signers:   make([]types.OperatorID, len(signedMsg.Signers)),
		Signature: make([]byte, len(signedMsg.Signature)),
	}
	copy(ret.Signers, signedMsg.Signers)
	copy(ret.Signature, signedMsg.Signature)

	ret.Message = &Message{
		MsgType:    signedMsg.Message.MsgType,
		Height:     signedMsg.Message.Height,
		Round:      signedMsg.Message.Round,
		Identifier: make([]byte, len(signedMsg.Message.Identifier)),
		Data:       make([]byte, len(signedMsg.Message.Data)),
	}
	copy(ret.Message.Identifier, signedMsg.Message.Identifier)
	copy(ret.Message.Data, signedMsg.Message.Data)
	return ret
}

// Validate returns error if msg validation doesn't pass.
// Msg validation checks the msg, it's variables for validity.
func (signedMsg *SignedMessage) Validate() error {
	if len(signedMsg.Signature) != 96 {
		return errors.New("message signature is invalid")
	}
	if len(signedMsg.Signers) == 0 {
		return errors.New("message signers is empty")
	}

	// check unique signers
	signed := make(map[types.OperatorID]bool)
	for _, signer := range signedMsg.Signers {
		if signed[signer] {
			return errors.New("non unique signer")
		}
		if signer == 0 {
			return errors.New("signer ID 0 not allowed")
		}
		signed[signer] = true
	}

	return signedMsg.Message.Validate()
}
