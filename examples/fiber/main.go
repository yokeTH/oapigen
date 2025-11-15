package main

import (
	"fiber/book"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// app.Get("healthz", func(ctx fiber.Ctx) error {
	// 	return ctx.SendString("OK")
	// })
	// Book routes
	bookDB := book.NewBookDB()
	bookRepo := book.NewBookRepository(bookDB)
	bookUsecase := book.NewBookUsecase(bookRepo)
	bookHandler := book.NewBookHandler(bookUsecase)

	app.Post("/books", bookHandler.CreateBook)
	app.Get("/books/:id", bookHandler.GetBook)
	app.Put("/books/:id", bookHandler.UpdateBook)
	app.Delete("/books/:id", bookHandler.DeleteBook)
}
