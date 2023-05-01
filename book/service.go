package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(booksRequest BookRequest) (Book, error)
	Update(ID int, booksRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}
func (s *service) Create(booksRequest BookRequest) (Book, error) {

	price, _ := booksRequest.Price.Int64()
	rating, _ := booksRequest.Rating.Int64()
	discount, _ := booksRequest.Discount.Int64()

	buatBuku := Book{
		Title:       booksRequest.Title,
		Price:       int(price),
		Description: booksRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(buatBuku)
	return newBook, err

}

func (s *service) Update(ID int, booksRequest BookRequest) (Book, error) {

	book, err := s.repository.FindByID(ID)

	price, _ := booksRequest.Price.Int64()
	rating, _ := booksRequest.Rating.Int64()
	discount, _ := booksRequest.Discount.Int64()

	book.Title = booksRequest.Title
	book.Price = int(price)
	book.Description = booksRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newUpdate, err := s.repository.Update(book)
	return newUpdate, err

}

func (s *service) Delete(ID int) (Book, error) {

	book, err := s.repository.FindByID(ID)

	newDelete, err := s.repository.Delete(book)
	return newDelete, err

}
