package enum

type SnapPaymentStatus string

const (
	SnapPaymentStatusPending   SnapPaymentStatus = "pending"
	SnapPaymentStatusCompleted SnapPaymentStatus = "completed"
	SnapPaymentStatusExpired   SnapPaymentStatus = "expired"
	SnapPaymentStatusCanceled  SnapPaymentStatus = "canceled"
	SnapPaymentStatusRejected  SnapPaymentStatus = "rejected"
)
