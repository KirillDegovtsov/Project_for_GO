package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"project_university/domain"
	"project_university/usecases"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string
	maxLifeTime int64
	provider    usecases.Provider
	lock        sync.RWMutex
}

func NewManager(provider usecases.Provider, cookieName string, maxLifeTime int64) *Manager {
	return &Manager{
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		provider:    provider,
	}
}

func (m *Manager) SessionID() string {
	newID := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, newID); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(newID)
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session usecases.Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.SessionID()
		session, err = m.provider.SessionInit(sid)
		if err != nil {

			return nil, err
		}
		cookie := http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifeTime),
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			return nil, errors.New("internal error")
		}

		session, err = m.provider.SessionRead(sid)
		if err != nil {
			return nil, err
		}

	}
	return
}

func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return domain.InternalError
	} else {
		err := m.provider.SessionDestroy(cookie.Value)
		if err != nil {
			return err
		}
		expiration := time.Now()
		cookie := http.Cookie{
			Name:     m.cookieName,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
			MaxAge:   -1,
		}
		http.SetCookie(w, &cookie)
	}
	return nil
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() { m.GC() })
}
