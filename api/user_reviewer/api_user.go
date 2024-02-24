package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-campaign-details/helper"
)

func GetReviewerId(input ReviewerIdInput) (string, error) {
	// check service admin

	err := CheckServiceUserReviewer()
	if err != nil {
		return "", err
	}

	reviewerID := helper.UserReviwer{}
	reviewerID.UnixReviewer = input.UnixID
	// fetch get /getAdminID from service api
	serviceReviewer := os.Getenv("SERVICE_REVIEWER")
	// if service admin is empty return error
	if serviceReviewer == "" {
		return reviewerID.UnixReviewer, errors.New("service admin is empty")
	}
	resp, err := http.Get(serviceReviewer + "/api/v1/campaign/getUserCampaignID/" + reviewerID.UnixReviewer)
	if err != nil {
		return reviewerID.UnixReviewer, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return reviewerID.UnixReviewer, errors.New("failed to get admin status or admin not found")
	}

	var response helper.UserReviewerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	} else if response.Data.StatusAccountReviewer == "deactive" {
		return "", errors.New("reviewer account is deactive")
	} else if response.Data.StatusAccountReviewer == "active" {
		return reviewerID.UnixReviewer, nil
	} else {
		return "", errors.New("invalid reviewer status")
	}
}

// verify token from service user reviewer
func VerifyTokenReviewer(input string) (string, error) {

	err := CheckServiceUserReviewer()
	if err != nil {
		return "", err
	}

	// fetch get /verifyToken from service api
	serviceAdmin := os.Getenv("SERVICE_REVIEWER")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return "", errors.New("service user reviewer is empty")
	}
	req, err := http.NewRequest("GET", serviceAdmin+"/api/v1/verifyTokenReviewer", nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+input)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid token, account deactive or token expired")
	}

	var response helper.VerifyTokenApiUserReviewerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	}

	return response.Data.UnixReviewer, nil

}
