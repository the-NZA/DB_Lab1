package services

type Servicer interface {
	BookService() BookServicer
	// AuthorService() AuthorServicer
	// GenreService() GenreServicer
}

type BookServicer interface {
	Get() string
	Add()
	Update()
	Delete()
	FindAll()
}

type AuthorServicer interface {
	Get()
	Add()
	Update()
	Delete()
	FindAll()
}

type GenreServicer interface {
	Get()
	Add()
	Update()
	Delete()
	FindAll()
}
