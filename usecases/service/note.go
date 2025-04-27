package service

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"project_university/domain"
	"project_university/repositoty"

	"golang.org/x/crypto/chacha20poly1305"
)

type Note struct {
	repo repositoty.Note
	key  []byte
}

func NewNote(repo repositoty.Note) (*Note, error) {
	key, err := MakeKey()
	if err != nil {
		return nil, err
	}
	return &Note{
		repo: repo,
		key:  key,
	}, nil
}

func (n *Note) Get(note *domain.Note) (*domain.Note, error) {
	data, err := n.repo.Get(note)
	if err != nil {
		return nil, err
	}
	decrypted, err := Decrypt(data.Text, n.key)
	if err != nil {
		return nil, err
	}
	note.Text = decrypted
	return note, nil
}

func (n *Note) Post(note *domain.Note) error {
	encrypted, err := Encrypt(note.Text, n.key)
	if err != nil {
		return err
	}
	note.Text = encrypted
	return n.repo.Post(note)
}

func (n *Note) Put(note *domain.Note) error {
	encrypted, err := Encrypt(note.Text, n.key)
	if err != nil {
		return err
	}
	note.Text = encrypted
	return n.repo.Put(note)
}

func (n *Note) Delete(note *domain.Note) error {
	return n.repo.Delete(note)
}

func Encrypt(data string, key []byte) (string, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", domain.InternalError
	}
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", domain.InternalError
	}
	ciphertext := aead.Seal(nil, nonce, []byte(data), nil)
	fullData := append(nonce, ciphertext...)
	return base64.StdEncoding.EncodeToString(fullData), nil
}

func Decrypt(encrypted string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", domain.InternalError
	}
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", domain.InternalError
	}
	if len(data) < chacha20poly1305.NonceSizeX {
		return "", domain.InvalidData
	}
	nonce := data[:chacha20poly1305.NonceSizeX]
	ciphertext := data[chacha20poly1305.NonceSizeX:]
	decripted, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", domain.InternalError
	}
	return string(decripted), nil
}

func MakeKey() ([]byte, error) {
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, domain.InternalError
	}
	return key, nil
}
