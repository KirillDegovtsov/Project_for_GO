package repositoty

import "project_university/domain"

type Note interface {
	Get(note *domain.Note) (*domain.Note, error)
	Post(note *domain.Note) error
	Put(note *domain.Note) error
	Delete(note *domain.Note) error
}
