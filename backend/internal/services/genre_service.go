package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type GenreService struct {
	repository storer.GenreRepository
}

func (g *GenreService) Get(ID string) (models.Genre, error) {
	return g.repository.Get(ID)
}

func (g *GenreService) Add(genre models.Genre) (models.Genre, error) {
	//validation here
	return g.repository.Add(genre)
}

func (g *GenreService) Update(genre models.Genre) (models.Genre, error) {
	//validation here
	return g.repository.Update(genre)
}

func (g *GenreService) Delete(ID string) error {
	//validation here
	return g.repository.Delete(ID)
}

func (g *GenreService) GetAll() ([]models.Genre, error) {
	//validation here
	return g.repository.GetAll()
}
