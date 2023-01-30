package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/hfl0506/reader-app/internal/model"
	"github.com/hfl0506/reader-app/internal/utils"
)

func (h *Handler) GetBookById(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return err
	}

	b, err := h.BookStore.GetById(uuid)

	if err != nil {
		return e.JSON(http.StatusNotFound, utils.NewError(err))
	}

	if b == nil {
		return e.JSON(http.StatusNotFound, utils.NotFound())
	}

	return e.JSON(http.StatusOK, b)
}

func (h *Handler) GetBooks(e echo.Context) error {

	books, err := h.BookStore.GetAll()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, nil)
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": books,
	})
}

func (h *Handler) CreateBook(e echo.Context) error {
	var book *model.Book

	if err := e.Bind(&book); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	book.ID = uuid.New()

	created, err := h.BookStore.CreateOne(book)

	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return e.JSON(http.StatusCreated, created)
}

func (h *Handler) UpdateBook(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	book, err := h.BookStore.GetById(uuid)

	if err != nil {
		return e.JSON(http.StatusNotFound, utils.NotFound())
	}

	if err := e.Bind(book); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err = h.BookStore.UpdateOne(uuid, book); err != nil {
		return e.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return e.JSON(http.StatusOK, book)
}

func (h *Handler) DeleteBook(e echo.Context) error {
	uuid, err := uuid.Parse(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	err = h.BookStore.DeleteOne(uuid)

	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return e.JSON(http.StatusAccepted, echo.Map{
		"result": fmt.Sprintf("book %s has been delete", uuid.String()),
	})
}
