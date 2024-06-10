package booksRepositories

import (
	booksEntities "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/entities"
	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	"github.com/jmoiron/sqlx"
)

type IBooksRepository interface {
	InsertBook(req *booksModels.Book) (*booksEntities.Book, error)
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
	query := `INSERT INTO books (title, author) VALUES (:title, :author) RETURNING id, title, author`
	statement, err := b.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	var book booksEntities.Book
	err = statement.Get(&book, req)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
