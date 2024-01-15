package core

import "time"

type ReviewerDetail struct {
	ID               int       `json:"id"`
	UnixID           string    `json:"unix_id"`
	UserReviewerID   string    `json:"user_reviewer_id"`
	CampaignDetailID string    `json:"campaign_detail_id"`
	NameReviewer     string    `json:"name_reviewer"`
	Description      string    `json:"description"`
	StatusReview     string    `json:"status_review"`
	Rating           int       `json:"rating"`
	UpdateIDAdmin    string    `json:"updateId_admin"`
	UpdateAtAdmin    time.Time `json:"updateAt_admin"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
