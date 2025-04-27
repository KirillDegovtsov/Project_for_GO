package http

import (
	"fmt"
	"log"
	"net/http"
	"project_university/api/http/types"
	"project_university/domain"
	"project_university/usecases"
	"project_university/usecases/service"

	"github.com/go-chi/chi/v5"
)

// вызвать в маин гарпиш колектор для всех сессий по времени
// сделать свэгер
// добавить логирование
// написать тесты
// прикрутить нджиникс

type Handler struct {
	noteSrv    usecases.Note
	userSrv    usecases.User
	sessionMng *service.Manager
}

func NewHandler(noteSrv usecases.Note, userSrv usecases.User, sessionMng *service.Manager) *Handler {
	return &Handler{
		noteSrv:    noteSrv,
		userSrv:    userSrv,
		sessionMng: sessionMng,
	}
}

func (h *Handler) getNoteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetDeleteNoteHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, _ := h.sessionMng.SessionStart(w, r)
	req.UserId, _ = session.Get("user_id")
	resp, err := h.noteSrv.Get(req)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	log.Printf("GET")
	types.CreateHandlerRespose(w, nil, types.NoteHandler{Title: resp.Title, Text: resp.Text})
}

func (h *Handler) postNoteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostPutNoteHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, _ := h.sessionMng.SessionStart(w, r)
	req.UserId, _ = session.Get("user_id")
	if err = h.noteSrv.Post(req); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	fmt.Printf("POST")
}

func (h *Handler) putTitleNoteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePutTitleNoteHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, _ := h.sessionMng.SessionStart(w, r)
	var newData domain.Note
	newData.UserId, _ = session.Get("user_id")
	newData.Title = req.OldTitle
	oldData, err := h.noteSrv.Get(&newData)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	newData.Title = req.NewTitle
	newData.Text = oldData.Text
	if err = h.noteSrv.Delete(oldData); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	if err = h.noteSrv.Post(&newData); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	fmt.Printf("PUT")
}

func (h *Handler) putTextNoteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostPutNoteHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, _ := h.sessionMng.SessionStart(w, r)
	req.UserId, _ = session.Get("user_id")
	if err = h.noteSrv.Put(req); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	fmt.Printf("PUT")
}

func (h *Handler) deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetDeleteNoteHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, _ := h.sessionMng.SessionStart(w, r)
	req.UserId, _ = session.Get("user_id")
	if err := h.noteSrv.Delete(req); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	fmt.Printf("DELETE")
}

func (h *Handler) registerHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostUserHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	if err = h.userSrv.Post(req); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostUserHandlerRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	user, err := h.userSrv.Get(req)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	session, err := h.sessionMng.SessionStart(w, r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	err = session.Set("user_id", user.Id)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) putUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePutUserPasswordRequest(r)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	newData := &domain.User{
		Login:    req.Login,
		Password: req.OldPassword,
	}
	oldData, err := h.userSrv.Get(newData)
	if err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	newData.Id = oldData.Id
	newData.Password = req.NewPassword
	if err = h.userSrv.Put(newData); err != nil {
		types.CreateHandlerRespose(w, err, nil)
		return
	}
	log.Printf("PUT")
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := h.sessionMng.SessionStart(w, r)
		if err != nil {
			types.CreateHandlerRespose(w, err, nil)
			return
		}
		_, err = session.Get("user_id")
		if err != nil {
			types.CreateHandlerRespose(w, err, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) WithObjectHandlers(r chi.Router) {
	r.Post("/register", h.registerHandler)
	r.Post("/login", h.loginHandler)

	r.With(h.AuthMiddleware).Get("/my_note", h.getNoteHandler)
	r.With(h.AuthMiddleware).Post("/add", h.postNoteHandler)
	r.With(h.AuthMiddleware).Put("/change_title", h.putTitleNoteHandler)
	r.With(h.AuthMiddleware).Put("/change_text", h.putTextNoteHandler)
	r.With(h.AuthMiddleware).Delete("/delete", h.deleteNoteHandler)
	r.With(h.AuthMiddleware).Put("/change_password", h.putUserPasswordHandler)
}
