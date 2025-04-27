package ramstorage

import (
	"project_university/domain"
)

type Titles struct {
	data map[string]string
}

type Note struct {
	data map[string]*Titles
}

func NewNote() *Note {
	return &Note{
		data: make(map[string]*Titles),
	}
}

func (n *Note) Get(note *domain.Note) (*domain.Note, error) {
	if _, err := n.data[note.UserId]; !err {
		return nil, domain.NoteNotFound
	}
	if _, err := n.data[note.UserId].data[note.Title]; !err {
		return nil, domain.NoteNotFound
	}
	return &domain.Note{Title: note.Title, Text: n.data[note.UserId].data[note.Title], UserId: note.UserId}, nil
}

func (n *Note) Post(note *domain.Note) error {
	if _, err := n.data[note.UserId]; !err {
		n.data[note.UserId] = &Titles{
			data: make(map[string]string),
		}
	}
	if _, err := n.data[note.UserId].data[note.Title]; err {
		return domain.NoteAlreadyExists
	}
	n.data[note.UserId].data[note.Title] = note.Text
	return nil
}

func (n *Note) Put(note *domain.Note) error {
	if _, err := n.data[note.UserId]; !err {
		return domain.NoteNotFound
	}
	if _, err := n.data[note.UserId].data[note.Title]; !err {
		return domain.NoteNotFound
	}
	n.data[note.UserId].data[note.Title] = note.Text
	return nil
}

func (n *Note) Delete(note *domain.Note) error {
	if _, err := n.data[note.UserId]; !err {
		return domain.NoteNotFound
	}
	if _, err := n.data[note.UserId].data[note.Title]; !err {
		return domain.NoteNotFound
	}
	delete(n.data[note.UserId].data, note.Title)
	return nil
}
