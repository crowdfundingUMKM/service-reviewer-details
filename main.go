package main

import (
	"fmt"
	"log"
	"os"
	"service-reviewer-details/auth"
	"service-reviewer-details/config"
	"service-reviewer-details/core"
	"service-reviewer-details/database"
	"service-reviewer-details/handler"
	"service-reviewer-details/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Route Api

// User Reviewer
// Create Review with User Review Token? Active Campaign

// Update Review  with User Review Token? Active Campaign

// Get All Review  with User Review Token? Active Campaign : Deactive Campaign : Report Campaign

// // Delete Review ? Active Campaign

// Admin Service Accsess
// GET STATUS SERVICE WITH VERIFY TOKEN ADMIN SERVICE

// GET LOG SERVICE WITH VERIFY TOKEN ADMIN SERVICE

// Get All Review ? Active Campaign : Deactive Campaign : Report Campaign

// Delete Review

func main() {
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log
	// config.InitLog()
	// setup repository
	db := database.NewConnectionDB()
	reviewerDetailRepository := core.NewRepository(db)

	// SETUP SERVICE
	reviewerDetailService := core.NewService(reviewerDetailRepository)
	authService := auth.NewService()

	// setup handler
	reviewerDetailHandler := handler.NewReviewerDetailsHandler(reviewerDetailService, authService)

	// RUN SERVICE
	router := gin.Default()

	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group api
	api := router.Group("api/v1")

	// routing
	api.POST("/create_review/:campaign_id", middleware.AuthApiReviewerMiddleware(authService, reviewerDetailService), reviewerDetailHandler.CreateReview)

	// end Rounting
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
