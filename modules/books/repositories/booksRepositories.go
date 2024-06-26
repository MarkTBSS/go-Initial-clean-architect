package booksRepositories

import (
	"fmt"
	"log"
	"strings"

	booksEntities "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/entities"
	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	"github.com/jmoiron/sqlx"
)

type IBooksRepository interface {
	InsertBook(req *booksModels.Book) (*booksEntities.Book, error)
	RetrieveAllBooks() ([]*booksEntities.Book, error)
	RetrieveBookByField(req *booksModels.Book, field string) ([]*booksEntities.Book, error)
	RetrieveBookByDynamicField(fields map[string]string) ([]*booksEntities.Book, error)
	UpdateBook(req *booksModels.Book) (*booksEntities.Book, error)
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

func (b *booksRepository) RetrieveBookByField(req *booksModels.Book, field string) ([]*booksEntities.Book, error) {
	var query string
	var args []interface{}
	switch field {
	case "id":
		query = `SELECT id, title, author FROM books WHERE id = $1`
		args = append(args, req.Id)
	case "title":
		query = `SELECT id, title, author FROM books WHERE title = $1`
		args = append(args, req.Title)
	case "author":
		query = `SELECT id, title, author FROM books WHERE author = $1`
		args = append(args, req.Author)
	default:
		return nil, nil
	}
	var books []*booksEntities.Book
	err := b.db.Select(&books, query, args...)
	if err != nil {
		log.Printf("Failed to retrieve books by field: %v", err)
		return nil, err
	}
	return books, nil
}

func (b *booksRepository) RetrieveBookByDynamicField(fields map[string]string) ([]*booksEntities.Book, error) {
	var conditions []string
	var args []interface{}
	i := 1
	for field, value := range fields {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}

	query := fmt.Sprintf("SELECT id, title, author FROM books WHERE %s", strings.Join(conditions, " OR "))
	//fmt.Println(query)
	var books []*booksEntities.Book
	err := b.db.Select(&books, query, args...)
	if err != nil {
		log.Printf("Failed to retrieve books by dynamic fields: %v", err)
		return nil, err
	}
	return books, nil
}

func (b *booksRepository) UpdateBook(req *booksModels.Book) (*booksEntities.Book, error) {
	var sets []string
	var args []interface{}
	i := 1
	if req.Title != "" {
		sets = append(sets, fmt.Sprintf("title = $%d", i))
		args = append(args, req.Title)
		i++
	}
	if req.Author != "" {
		sets = append(sets, fmt.Sprintf("author = $%d", i))
		args = append(args, req.Author)
		i++
	}
	if len(sets) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	args = append(args, req.Id)
	query := fmt.Sprintf("UPDATE books SET %s WHERE id = $%d", strings.Join(sets, ", "), i)
	_, err := b.db.Exec(query, args...)
	if err != nil {
		log.Printf("Failed to update book: %v", err)
		return nil, err
	}
	// After updating, retrieve the updated book to return it
	updatedBookQuery := "SELECT id, title, author FROM books WHERE id = $1"
	var updatedBook booksEntities.Book
	err = b.db.Get(&updatedBook, updatedBookQuery, req.Id)
	if err != nil {
		log.Printf("Failed to retrieve updated book: %v", err)
		return nil, err
	}
	return &updatedBook, nil
}
