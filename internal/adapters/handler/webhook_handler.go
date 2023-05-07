package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/LordMoMA/Hexagonal-Architecture/internal/adapters/repository"
	"github.com/LordMoMA/Hexagonal-Architecture/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateMembershipStatus(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// get api key from config
	apiKey := apiCfg.APIKey

	// check if api key is valid
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		HandleError(ctx, http.StatusBadRequest, errors.New("no api key provided"))
		return
	}
	apiString := strings.TrimPrefix(authHeader, "ApiKey " )

	if apiString != apiKey {
		HandleError(ctx, http.StatusBadRequest, errors.New("invalid api key"))
		return
	}

	// get user_id from request body
	userID, err := ctx.Request.Body.Read([]byte("user_id"))
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	
	err = h.svc.UpdateMembershipStatus(strconv.Itoa(userID), true)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User's membership status updated successfully",
	})
}
