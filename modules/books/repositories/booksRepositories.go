package booksRepositories

import (
	"log"

	booksEntities "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/entities"
	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	"github.com/jmoiron/sqlx"
)

type IBooksRepository interface {
	InsertBook(req *booksModels.Book) (*booksEntities.Book, error)
	RetrieveAllBooks() ([]*booksEntities.Book, error)
}

type booksRepository struct {
	db *sqlx.DB
}

func NewBooksRepository(db *sqlx.DB) IBooksRepository {
	return &booksRepository{
		db: db,
	}
}

func (b *booksRepository) InsertBook(req *booksModels.Book) (*booksEntities.Book, error) {
	query := `INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id, title, author`
	var book booksEntities.Book
	err := b.db.QueryRow(query, req.Title, req.Author).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		log.Printf("Failed to insert: %v", err)
		return nil, err
	}
	return &book, nil
}

func (b *booksRepository) RetrieveAllBooks() ([]*booksEntities.Book, error) {
	query := `SELECT id, title, author FROM books`
	var books []*booksEntities.Book
	err := b.db.Select(&books, query)
	if err != nil {
		log.Printf("Failed to retrieve books: %v", err)
		return nil, err
	}
	return books, nil
}
