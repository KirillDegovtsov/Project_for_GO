package types

import (
	"encoding/json"
	"net/http"
	"project_university/domain"
	"time"
)

type GetDeleteNoteHandlerRequest struct {
	Title string `json:"title"`
}

func CreateGetDeleteNoteHandlerRequest(r *http.Request) (*domain.Note, error) {
	var req GetDeleteNoteHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, domain.BadRequest
	}
	return &domain.Note{Title: req.Title}, nil
}

type TimeValue struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

type NoteHandler struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	CreatedTime TimeValue `json:"created_time"`
	LastChange  TimeValue `json:"last_change"`
}

func CreatePostPutNoteHandlerRequest(r *http.Request) (*domain.Note, error) {
	var req NoteHandler
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, domain.BadRequest
	}
	return &domain.Note{Title: req.Title, Text: req.Text}, nil
}

type PutTitleNoteHandlerRequest struct {
	OldTitle string `json:"old_title"`
	NewTitle string `json:"new_title"`
}

func CreatePutTitleNoteHandlerRequest(r *http.Request) (*PutTitleNoteHandlerRequest, error) {
	var req PutTitleNoteHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, domain.BadRequest
	}
	return &req, nil
}

type PostUserHandlerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func CreatePostUserHandlerRequest(r *http.Request) (*domain.User, error) {
	var req PostUserHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, domain.BadRequest
	}
	return &domain.User{Login: req.Login, Password: req.Password}, nil
}

type PutUserPasswordRequest struct {
	Login       string `json:"login"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func CreatePutUserPasswordRequest(r *http.Request) (*PutUserPasswordRequest, error) {
	var req PutUserPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, domain.BadRequest
	}
	return &req, nil
}

func CreateHandlerRespose(w http.ResponseWriter, err error, resp any) {
	if err == domain.NoteNotFound || err == domain.UserNotFound {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err == domain.BadRequest || err == domain.InvalidPassword || err == domain.UserAlreadyExists {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} else if err == domain.InternalError || err == domain.NoteAlreadyExists || err == domain.KeyAlreadyExists || err == domain.InvalidData {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else if err == domain.Unauthorized {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}
