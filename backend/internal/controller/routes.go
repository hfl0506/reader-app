package controller

import (
	"github.com/hfl0506/reader-app/internal/store"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	UserStore *store.UserStore
	BookStore *store.BookStore
}

func NewHandler(payload Handler) *Handler {
	return &Handler{
		UserStore: payload.UserStore,
		BookStore: payload.BookStore,
	}
}

func (h *Handler) RegisterRoutes(v1 *echo.Group) {
	bookRoutes := v1.Group("/books")
	bookRoutes.GET("", h.GetBooks)
	bookRoutes.GET("/:id", h.GetBookById)
	bookRoutes.POST("", h.CreateBook)
	bookRoutes.PUT("/:id", h.UpdateBook)
	bookRoutes.DELETE("/:id", h.DeleteBook)

	userRoutes := v1.Group("/users")
	userRoutes.GET("/:id", h.GetUserById)
	userRoutes.POST("", h.CreateUser)
	userRoutes.PUT("/:id", h.UpdateUser)
	userRoutes.DELETE("/:id", h.DeleteUser)

	authRoutes := v1.Group("/auth")
	authRoutes.POST("/login", h.Login)

	imageRoutes := v1.Group("/images")
	imageRoutes.POST("/upload", h.UploadImage)
	imageRoutes.POST("/download", h.DownloadFile)
}
