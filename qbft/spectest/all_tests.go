package spectest

import (
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/commit"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/decided"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/futuremsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/latemsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/processmsg"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/prepare"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposal"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposer"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/roundchange"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/startinstance"
	"testing"
)

type SpecTest interface {
	TestName() string
	Run(t *testing.T)
}

var AllTests = []SpecTest{
	// sanity tests to be completed
	decided.Valid(),
	decided.HasQuorum(),
	decided.LateDecidedBiggerQuorum(),
	decided.LateDecidedSmallerQuorum(),
	decided.DuplicateMsg(),
	decided.WrongSignature(),
	decided.CurrentInstancePastRound(),
	decided.CurrentInstanceFutureRound(),
	processmsg.NoInstanceRunning(),

	latemsg.LateCommit(),
	latemsg.LateCommitPastRound(),

	futuremsg.Cleanup(),
	futuremsg.DuplicateSigner(),
	futuremsg.F1FutureMsgs(),

	startinstance.Valid(),
	startinstance.EmptyValue(),
	startinstance.NilValue(),

	proposer.FourOperators(),
	proposer.SevenOperators(),
	proposer.TenOperators(),
	proposer.ThirteenOperators(),

	proposal.PreparedPreviouslyJustification(),
	proposal.FirstRoundJustification(),
	proposal.FutureRoundPrevNotPrepared(),
	proposal.FutureRound(),
	proposal.NoRCJustification(),
	proposal.WrongProposer(),
	proposal.WrongSignature(),
	proposal.InvalidFullData(),

	prepare.HappyFlow(),

	commit.HappyFlow(),

	roundchange.HappyFlow(),
	roundchange.F1Speedup(),
	roundchange.F1SpeedupPrevPrepared(),
	roundchange.NotProposer(),
	roundchange.ValidJustification(),

	// sanity tests (first version)
	proposal.NotPreparedPreviouslyJustification(),
	decided.LateDecided(),
	processmsg.FullDecided(),
	tests.HappyFlow(),
	tests.SevenOperators(),

	//timeout.Round1(),
	//timeout.Round2(),
	//timeout.Round3(),
	//timeout.Round5(),
	//timeout.Round20(),
	//
	//decided.Valid(),
	//decided.HasQuorum(),
	//decided.LateDecided(),
	//decided.LateDecidedBiggerQuorum(),
	//decided.LateDecidedSmallerQuorum(),
	//decided.NoQuorum(),
	//decided.DuplicateMsg(),
	//decided.DuplicateSigners(),
	//decided.Invalid(),
	////decided.InvalidFullData(), // TODO: implement
	//decided.InvalidValCheckData(), // TODO: fix
	//decided.PastInstance(),
	//decided.UnknownSigner(),
	//decided.WrongMsgType(),
	//decided.WrongSignature(),
	//decided.MultiDecidedInstances(),
	//decided.FutureInstance(),
	//decided.CurrentInstance(), // TODO: fix
	//decided.CurrentInstancePastRound(), // TODO: fix
	//decided.CurrentInstanceFutureRound(), // TODO: fix
	//
	//processmsg.MsgError(),
	//processmsg.BroadcastedDecided(),
	//processmsg.SingleConsensusMsg(),
	//processmsg.FullDecided(),
	//processmsg.InvalidIdentifier(),
	//processmsg.NoInstanceRunning(),
	//
	//latemsg.LateCommit(),
	//latemsg.LateCommitPastRound(),
	//latemsg.LateCommitPastInstance(),
	//latemsg.LatePrepare(),
	//latemsg.LatePreparePastInstance(),
	//latemsg.LatePreparePastRound(),
	//latemsg.LateProposal(),
	//latemsg.LateProposalPastInstance(),
	//latemsg.LateProposalPastRound(),
	//latemsg.LateRoundChange(),
	//latemsg.LateRoundChangePastInstance(),
	//latemsg.LateRoundChangePastRound(),
	//latemsg.FullFlowAfterDecided(),
	//
	//futuremsg.NoSigners(),
	//futuremsg.MultiSigners(),
	//futuremsg.Cleanup(),
	//futuremsg.DuplicateSigner(),
	//futuremsg.F1FutureMsgs(),
	//futuremsg.InvalidMsg(),
	//futuremsg.UnknownSigner(),
	//futuremsg.WrongSig(),
	//
	//startinstance.Valid(),
	//startinstance.EmptyValue(),
	//startinstance.NilValue(),
	//startinstance.PostFutureDecided(),
	//startinstance.FirstHeight(),
	//startinstance.PreviousDecided(),
	//startinstance.PreviousNotDecided(),
	//startinstance.InvalidValue(),
	//
	//proposer.FourOperators(),
	//proposer.SevenOperators(),
	//proposer.TenOperators(),
	//proposer.ThirteenOperators(),
	//
	//messages.RoundChangeDataInvalidJustifications(),
	//messages.RoundChangePrePreparedJustifications(),
	//messages.RoundChangeNotPreparedJustifications(),
	//messages.CommitDataEncoding(),
	//messages.MsgNilIdentifier(),
	//messages.MsgNonZeroIdentifier(),
	//messages.MsgTypeUnknown(),
	//messages.PrepareDataEncoding(),
	//messages.ProposeDataEncoding(),
	//messages.SignedMsgNoSigners(),
	//messages.SignedMsgDuplicateSigners(),
	//messages.SignedMsgMultiSigners(),
	//messages.GetRoot(),
	//messages.SignedMessageEncoding(),
	//messages.CreateProposal(),
	//messages.CreateProposalPreviouslyPrepared(),
	//messages.CreateProposalNotPreviouslyPrepared(),
	//messages.CreatePrepare(),
	//messages.CreateCommit(),
	//messages.CreateRoundChange(),
	//messages.CreateRoundChangePreviouslyPrepared(),
	//messages.RoundChangeDataEncoding(),
	//messages.SignedMessageSigner0(),
	//messages.MarshaJustifications(),
	//messages.MarshaJustificationsWithFullData(),
	//messages.UnMarshaJustifications(),
	//
	//tests.HappyFlow(),
	//tests.SevenOperators(),
	//tests.TenOperators(),
	//tests.ThirteenOperators(),
	//

	//proposal.InvalidFullData(),
	//proposal.PastRoundProposalPrevPrepared(),
	//proposal.NotPreparedPreviouslyJustification(),
	//proposal.PreparedPreviouslyJustification(),
	//proposal.DifferentJustifications(),
	//proposal.JustificationsNotHeighest(),
	//proposal.JustificationsValueNotJustified(), // TODO: fix
	//proposal.DuplicateMsg(),
	//proposal.DuplicateMsgDifferentRoot(),
	//proposal.FirstRoundJustification(),
	//proposal.FutureRoundPrevNotPrepared(),
	//proposal.FutureRound(),
	//proposal.InvalidRoundChangeJustificationPrepared(),
	//proposal.InvalidRoundChangeJustification(),
	//proposal.PreparedPreviouslyNoRCJustificationQuorum(),
	//proposal.NoRCJustification(),
	proposal.PreparedPreviouslyNoPrepareJustificationQuorum(), // TODO: fix
	//proposal.PreparedPreviouslyDuplicatePrepareMsg(),
	//proposal.PreparedPreviouslyDuplicatePrepareQuorum(),
	//proposal.PreparedPreviouslyDuplicateRCMsg(),
	//proposal.PreparedPreviouslyDuplicateRCQuorum(),
	//proposal.DuplicateRCMsg(),
	proposal.InvalidPrepareJustificationValue(), // TODO: fix
	proposal.InvalidPrepareJustificationRound(), // TODO: fix
	//proposal.InvalidValueCheck(), // TODO: implement
	//proposal.MultiSigner(),
	//proposal.PostDecided(),
	//proposal.PostPrepared(),
	//proposal.SecondProposalForRound(),
	//proposal.WrongHeight(),
	//proposal.WrongProposer(),
	//proposal.WrongSignature(),
	//proposal.UnknownSigner(),
	//
	//prepare.DuplicateMsg(),
	//prepare.HappyFlow(),
	//prepare.InvalidPrepareData(), // TODO: fix
	//prepare.MultiSigner(),
	//prepare.NoPreviousProposal(),
	//prepare.OldRound(),
	//prepare.FutureRound(),
	//prepare.PostDecided(),
	//prepare.WrongData(),
	//prepare.WrongHeight(),
	//prepare.WrongSignature(),
	//prepare.UnknownSigner(),
	//
	//commit.CurrentRound(),
	//commit.FutureRound(),
	//commit.PastRound(),
	//commit.DuplicateMsg(),
	//commit.HappyFlow(),
	//commit.InvalidCommitData(), // TODO: fix
	//commit.PostDecided(),
	//commit.WrongData1(),
	//commit.WrongData2(),
	//commit.MultiSignerWithOverlap(),
	//commit.MultiSignerNoOverlap(),
	//commit.DuplicateSigners(),
	//commit.NoPrevAcceptedProposal(),
	//commit.WrongHeight(),
	//commit.WrongSignature(),
	//commit.UnknownSigner(),
	//commit.InvalidValCheck(),
	//commit.NoPrepareQuorum(),
	//
	//roundchange.HappyFlow(),
	//roundchange.WrongHeight(),
	//roundchange.WrongSig(),
	//roundchange.UnknownSigner(),
	//roundchange.MultiSigner(),
	//roundchange.QuorumNotPrepared(),
	//roundchange.QuorumPrepared(),
	//roundchange.PeerPrepared(),
	//roundchange.PeerPreparedDifferentHeights(),
	//roundchange.JustificationWrongValue(), // TODO: fix
	//roundchange.JustificationWrongRound(),
	//roundchange.JustificationNoQuorum(),
	//roundchange.JustificationMultiSigners(),
	//roundchange.JustificationInvalidSig(),
	//roundchange.JustificationInvalidRound(),
	//roundchange.JustificationInvalidPrepareData(), // TODO: fix
	//roundchange.JustificationDuplicateMsg(),
	//roundchange.InvalidRoundChangeData(), // TODO: fix
	//roundchange.F1DifferentFutureRounds(), // TODO: fix
	//roundchange.F1DifferentFutureRoundsNotPrepared(), // TODO: fix
	//roundchange.PastRound(),
	//roundchange.DuplicateMsgQuorum(),
	//roundchange.DuplicateMsgQuorumPreparedRCFirst(), // TODO: fix
	//roundchange.DuplicateMsg(),
	//roundchange.NotProposer(),
	//roundchange.ValidJustification(),
	//roundchange.F1DuplicateSigner(),
	//roundchange.F1DuplicateSignerNotPrepared(),
	//roundchange.F1Speedup(),
	//roundchange.F1SpeedupPrevPrepared(),
	//roundchange.AfterProposal(),
	//roundchange.RoundChangePartialQuorum(), // TODO: fix
	//roundchange.QuorumOrder2(),
	//roundchange.QuorumOrder1(), // TODO: fix
	//roundchange.QuorumMsgNotPrepared(), // TODO: fix
	//roundchange.JustificationPastRound(),
}
