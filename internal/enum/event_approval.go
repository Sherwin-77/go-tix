package enum

type EventApprovalStatus string

const (
	EventApprovalStatusPending  EventApprovalStatus = "pending"
	EventApprovalStatusApproved EventApprovalStatus = "approved"
	EventApprovalStatusRejected EventApprovalStatus = "rejected"
)
