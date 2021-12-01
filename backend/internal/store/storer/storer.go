package storer

type Storer interface {
	Books() BookReporsitory
	Authors() AuthorRepository
	Genres() GenreRepository
}
