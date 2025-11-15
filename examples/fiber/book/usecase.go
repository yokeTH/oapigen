package book

type bookUsecase struct {
	bookRepository *bookRepository
}

func NewBookUsecase(bookRepository *bookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepository}
}

func (u *bookUsecase) Create(book bookModel) error {
	return u.bookRepository.Create(book)
}

func (u *bookUsecase) Read(id string) (bookModel, bool) {
	return u.bookRepository.Read(id)
}

func (u *bookUsecase) Update(id string, updatedBook bookModel) error {
	return u.bookRepository.Update(id, updatedBook)
}

func (u *bookUsecase) Delete(id string) error {
	return u.bookRepository.Delete(id)
}
