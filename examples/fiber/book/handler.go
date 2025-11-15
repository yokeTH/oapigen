package book

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type bookHandler struct {
	bookUseCase *bookUsecase
}

func NewBookHandler(bookUseCase *bookUsecase) *bookHandler {
	return &bookHandler{bookUseCase: bookUseCase}
}

// CreateBook handles POST /books
func (h *bookHandler) CreateBook(ctx fiber.Ctx) error {
	var req CreateBookRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request"})
	}
	if req.Title == "" || req.Author == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Title and Author are required"})
	}

	id := uuid.New().String()
	book := bookModel{
		ID:     id,
		Title:  req.Title,
		Author: req.Author,
	}
	if err := h.bookUseCase.Create(book); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	resp := BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}
	return ctx.Status(fiber.StatusCreated).JSON(Success(resp))
}

func (h *bookHandler) GetBook(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	book, ok := h.bookUseCase.Read(id)
	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "Book not found"})
	}
	resp := BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}
	return ctx.JSON(Success(resp))
}

func (h *bookHandler) UpdateBook(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	var req UpdateBookRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request"})
	}

	book, ok := h.bookUseCase.Read(id)
	if !ok {
		return ctx.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "Book not found"})
	}

	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}

	if err := h.bookUseCase.Update(id, book); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	resp := BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}
	return ctx.JSON(Success(resp))
}

func (h *bookHandler) DeleteBook(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.bookUseCase.Delete(id); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "Book not found"})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
