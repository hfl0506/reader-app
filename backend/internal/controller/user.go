package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hfl0506/reader-app/internal/model"
	"github.com/hfl0506/reader-app/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserById(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return err
	}

	u, err := h.UserStore.GetUserById(uuid)

	if err != nil {
		return e.JSON(http.StatusNotFound, utils.NewError(err))
	}

	if u == nil {
		return e.JSON(http.StatusNotFound, "result empty")
	}

	return e.JSON(http.StatusOK, u)
}

func (h *Handler) CreateUser(e echo.Context) error {
	var user *model.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "Hash password failed")
	}

	user.ID = uuid.New()
	user.Password = hash

	created, err := h.UserStore.CreateUser(user)

	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return e.JSON(http.StatusCreated, created)
}

func (h *Handler) UpdateUser(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	user, err := h.UserStore.GetUserById(uuid)

	if err != nil {
		return e.JSON(http.StatusNotFound, utils.NewError(err))
	}

	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if err = h.UserStore.UpdateUser(uuid, user); err != nil {
		return e.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return e.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if err := h.UserStore.DeleteUser(uuid); err != nil {
		return e.JSON(http.StatusNotFound, utils.NewError(err))
	}

	return e.JSON(http.StatusAccepted, echo.Map{
		"message": fmt.Sprintf("user %s", uuid),
	})
}
