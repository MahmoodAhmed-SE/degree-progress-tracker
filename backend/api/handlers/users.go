package handlers

import (
	"log"
	"net/http"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api/schemas"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/database/models"
	"github.com/gin-gonic/gin"
)

func InitUsersRoute() {
	AddAuthRoute(http.MethodPost, "/v1/users", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			schemas.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if err := user.Register(); err != nil {
			schemas.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		schemas.SuccessResponse(ctx, http.StatusCreated, "User registered successfully", nil)
	})
	AddAuthRoute(http.MethodPost, "/v1/users/login", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			schemas.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		token, err := user.Login()
		if err != nil {
			schemas.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		log.Println("Login token:", token)
		schemas.SuccessResponse(ctx, http.StatusOK, "Login successful", gin.H{"token": token})
	})
	AddRoute(http.MethodGet, "/v1/users/me", func(ctx *gin.Context) {
		var user models.User
		CurrentUser, exists := ctx.Get("user")
		if !exists {
			schemas.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}
		user = CurrentUser.(models.User)
		if err := user.GetByID(); err != nil {
			schemas.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		schemas.SuccessResponse(ctx, http.StatusOK, "User fetched successfully", user)
	})
	AddRoute(http.MethodPut, "/v1/users/me", func(ctx *gin.Context) {
		var user models.User
		userInterface, exists := ctx.Get("user")
		if !exists {
			schemas.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}
		user = userInterface.(models.User)
		if err := ctx.ShouldBindJSON(&user); err != nil {
			schemas.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if err := user.Update(); err != nil {
			schemas.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		schemas.SuccessResponse(ctx, http.StatusOK, "User updated successfully", nil)
	})

	AddRoute(http.MethodGet, "/v1/users", func(ctx *gin.Context) {
		var user models.User
		users, err := user.GetAll()
		if err != nil {
			schemas.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		schemas.SuccessResponse(ctx, http.StatusOK, "Users fetched successfully", users)
	}, "USERS_READ")
	AddRoute(http.MethodPost, "/v1/users/logout", func(ctx *gin.Context) {
		var user models.User
		authUser, exists := ctx.Get("user")
		if !exists {
			schemas.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}
		user = authUser.(models.User)
		token, err := user.Logout()
		if err != nil {
			schemas.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		schemas.SuccessResponse(ctx, http.StatusOK, "Logout successful", gin.H{"token": token})

	})
}
