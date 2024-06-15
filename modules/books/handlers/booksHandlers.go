package booksHandlers

import (
	"fmt"

	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	booksUsecases "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/usecases"
	"github.com/gofiber/fiber/v2"
)

type IBooksHandlers interface {
	InsertBook(c *fiber.Ctx) error
	RetrieveAllBooks(c *fiber.Ctx) error
	RetrieveBookByField(c *fiber.Ctx) error
	RetrieveBookByDynamicField(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
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

func (b *booksHandlers) RetrieveAllBooks(c *fiber.Ctx) error {
	books, err := b.booksUsecase.RetrieveAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func (b *booksHandlers) RetrieveBookByField(c *fiber.Ctx) error {
	field := c.Query("field")
	value := c.Query("value")
	fmt.Println(field)
	fmt.Println(value)
	req := &booksModels.Book{}
	switch field {
	case "id":
		req.Id = value
	case "title":
		req.Title = value
	case "author":
		req.Author = value
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid field",
		})
	}
	books, err := b.booksUsecase.RetrieveBookByField(req, field)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}
func (b *booksHandlers) RetrieveBookByDynamicField(c *fiber.Ctx) error {
	queryParams := c.Queries()
	req := make(map[string]string)
	for field, value := range queryParams {
		if value != "" {
			req[field] = value
		}
	}
	books, err := b.booksUsecase.RetrieveBookByDynamicField(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func (b *booksHandlers) UpdateBook(c *fiber.Ctx) error {
	var req booksModels.Book
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	req.Id = c.Params("id")
	updatedBook, err := b.booksUsecase.UpdateBook(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedBook)
}
