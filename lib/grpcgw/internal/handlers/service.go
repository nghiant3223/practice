package handlers

import (
	"context"
	"errors"
	"math/rand"

	"github.com/nghiant3223/gopractice/grpcgw/protogens/protos"
)

type BookServiceHandler struct{}

var books []*book.Book

var authors = []*book.Author{
	{
		Id:        1,
		FirstName: "1st",
		LastName:  "author",
	}, {
		Id:        2,
		FirstName: "2nd",
		LastName:  "author",
	},
}

var (
	ErrAuthorNotFound = errors.New("author not found")
)

func (s *BookServiceHandler) CreateBook(c context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	var author *book.Author
	for _, a := range authors {
		if a.Id == req.AuthorId {
			author = a
		}
	}
	if author == nil {
		return nil, ErrAuthorNotFound
	}

	newBook := &book.Book{
		Id:     rand.Int31(),
		Name:   req.Name,
		Author: author,
	}
	books = append(books, newBook)

	res := &book.CreateBookResponse{
		Book: newBook,
	}
	return res, nil
}

func (s *BookServiceHandler) GetBooks(context.Context, *book.GetBooksRequest) (*book.GetBooksResponse, error) {
	res := &book.GetBooksResponse{
		Books: books,
	}
	return res, nil
}
