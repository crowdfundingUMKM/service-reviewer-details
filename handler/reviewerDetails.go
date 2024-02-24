package handler

import (
	"service-reviewer-details/auth"
	"service-reviewer-details/core"
)

type reviewerDetailHandler struct {
	reviewerDetailService core.Service
	authService           auth.Service
}

func NewReviewerDetailsHandler(reviewerDetailService core.Service, authService auth.Service) *reviewerDetailHandler {
	return &reviewerDetailHandler{reviewerDetailService, authService}
}

// func (h *reviewerDetailHandler) CreateReview(c *gin.Context) {
// 	var input core.CreateReviewerDetailInput
