package book

import "fmt"

type bookRepository struct {
	db *map[string]bookModel // Just database mock
}

func NewBookDB() *map[string]bookModel {
	db := make(map[string]bookModel)
	return &db
}

func NewBookRepository(db *map[string]bookModel) *bookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(book bookModel) error {
	(*r.db)[book.ID] = book
	return nil
}

func (r *bookRepository) Read(id string) (bookModel, bool) {
	book, ok := (*r.db)[id]
	return book, ok
}

func (r *bookRepository) Update(id string, updatedBook bookModel) error {
	if _, ok := (*r.db)[id]; !ok {
		return fmt.Errorf("book not found")
	}
	(*r.db)[id] = updatedBook
	return nil
}

func (r *bookRepository) Delete(id string) error {
	if _, ok := (*r.db)[id]; !ok {
		return fmt.Errorf("book not found")
	}
	delete(*r.db, id)
	return nil
}
