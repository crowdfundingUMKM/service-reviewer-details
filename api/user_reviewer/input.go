package api

type ReviewerIdInput struct {
	UnixID string `uri:"reviewer_id" binding:"required"`
}

type VerifyTokenUserReviewerInput struct {
	Token string `json:"token" binding:"required"`
}

type ReviewerId struct {
	UnixReviewer string
}
