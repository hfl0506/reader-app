package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/hfl0506/reader-app/internal/utils"
	"github.com/labstack/echo/v4"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jwtCustomClaims struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(e echo.Context) error {
	var payload *LoginPayload

	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	isExist, err := h.UserStore.GetUserByEmail(payload.Email)

	if err != nil {
		return e.JSON(http.StatusNotFound, utils.NewError(err))
	}

	isMatchPassword := utils.ValidatePassword(payload.Password, isExist.Password)

	if !isMatchPassword {
		return e.JSON(http.StatusUnauthorized, "password invalid")
	}

	jwtPayload := &utils.JwtPayload{
		Id:    isExist.ID,
		Name:  isExist.Name,
		Email: isExist.Email,
	}

	atExpireTime := time.Minute * 30
	rtExpireTime := time.Hour * 168

	at, err := utils.GenerateToken(jwtPayload, atExpireTime)

	fmt.Println("after gen at")

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	rt, err := utils.GenerateToken(jwtPayload, rtExpireTime)

	fmt.Println("after gen rt")

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	return e.JSON(http.StatusAccepted, echo.Map{
		"accessToken":  at,
		"refreshToken": rt,
	})
}
