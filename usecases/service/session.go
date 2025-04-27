package service

import (
	"project_university/domain"
	"sync"
	"time"
)

type Session struct {
	sid        string
	data       map[string]string
	lastAccess time.Time
	lock       sync.RWMutex
}

func (s *Session) Set(key string, value string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, err := s.data[key]; err {
		return domain.KeyAlreadyExists
	}
	s.data[key] = value
	return nil
}

func (s *Session) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if _, err := s.data[key]; !err {
		return "", domain.Unauthorized
	}
	return s.data[key], nil
}

func (s *Session) Delete(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, err := s.data[key]; !err {
		return domain.Unauthorized
	}
	delete(s.data, key)
	return nil
}

func (s *Session) SessionID() string {
	return s.sid
}
