package consensus

type GBFT interface {
	StartConsensus(request *RequestMsg) (*PrePrepareMsg, error)
	PrePrepare(prePrepareMsg *PrePrepareMsg) (*VoteMsg, error)
	Prepare(prepareMsg *VoteMsg) (*VoteMsg, *ViewChangeMsg, error)
	Commit(commitMsg *VoteMsg) (*ReplyMsg, *RequestMsg, error)
	ViewChange(viewChangeMsg *ViewChangeMsg) (*ViewChangeMsg, error)
	ViewChangeClame(viewChangeMsgClame *ViewChangeClameMsg) error
}
