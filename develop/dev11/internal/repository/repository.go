package repository

import "WB-L2/develop/dev11/internal/repository/inmemory"

// Набор всех хранилищ
type Repository struct {
	inmemory.Event
}

func NewRepository() Repository {
	return Repository{
		Event: inmemory.NewEvent(),
	}
}
