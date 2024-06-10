package booksHandlers

import (
	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	booksUsecases "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/usecases"
	"github.com/gofiber/fiber/v2"
)

type IBooksHandlers interface {
	InsertBook(c *fiber.Ctx) error
}
type booksHandlers struct {
	booksUsecase booksUsecases.IBooksUsecase
}

func NewBooksHandlers(booksUsecase booksUsecases.IBooksUsecase) IBooksHandlers {
	return &booksHandlers{
		booksUsecase: booksUsecase,
	}
}

func (b *booksHandlers) InsertBook(c *fiber.Ctx) error {
	var book booksModels.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	insertedBook, err := b.booksUsecase.InsertBook(&book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(insertedBook)
}
