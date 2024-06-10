package booksUsecases

import (
	booksEntities "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/entities"
	booksModels "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/models"
	booksRepositories "github.com/MarkTBSS/go-Initial-clean-architect/modules/books/repositories"
)

type IBooksUsecase interface {
	InsertBook(req *booksModels.Book) (*booksEntities.Book, error)
}
type booksUsecase struct {
	booksRepository booksRepositories.IBooksRepository
}

func NewBooksUsecase(booksRepository booksRepositories.IBooksRepository) IBooksUsecase {
	return &booksUsecase{
		booksRepository: booksRepository,
	}
}

func (b *booksUsecase) InsertBook(req *booksModels.Book) (*booksEntities.Book, error) {
	result, err := b.booksRepository.InsertBook(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
