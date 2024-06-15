package main

import (
	"log"

	booksHandlers "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/handlers"
	booksRepositories "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/repositories"
	booksUsecases "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/usecases"
	"github.com/MarkTBSS/go-Initial-clean-architect/pkg/databases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := databases.DatabaseConnect()
	defer db.Close()

	booksRepo := booksRepositories.NewBooksRepository(db)
	booksUsecase := booksUsecases.NewBooksUsecase(booksRepo)
	booksHandler := booksHandlers.NewBooksHandlers(booksUsecase)

	app := fiber.New()
	app.Post("/books", booksHandler.InsertBook)
	app.Get("/books", booksHandler.RetrieveAllBooks)
	app.Get("/books/search", booksHandler.RetrieveBookByField)
	app.Get("/books/dynamic-search", booksHandler.RetrieveBookByDynamicField)
	app.Patch("/books/update/:id", booksHandler.UpdateBook)
	log.Fatal(app.Listen(":8000"))
}
