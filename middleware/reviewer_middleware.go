package middleware

import (
	"net/http"
	api_reviewer "service-reviewer-details/api/user_reviewer"
	"service-reviewer-details/auth"
	"service-reviewer-details/core"
	"service-reviewer-details/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthApiReviewerMiddleware(authService auth.Service, userService core.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized API User Reviewer", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// token, err := authService.ValidateToken(tokenString)
		// if err != nil {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// claim, ok := token.Claims.(jwt.MapClaims)

		// if !ok || !token.Valid {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// adminID := claim["admin_id"].(string)

		if tokenString == "" {
			response := helper.APIResponse("Unauthorized API User Reviewer", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		reviewerID, err := api_reviewer.VerifyTokenReviewer(tokenString)

		if err != nil { //wrong token
			response := helper.APIResponse("Unauthorized API User Reviewer", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUserReviewer", api_reviewer.ReviewerId{UnixReviewer: reviewerID})
		c.Next()
	}
}
