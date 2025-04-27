package service

import (
	"project_university/domain"
	"project_university/usecases"
	"sync"
	"time"
)

type Povider struct {
	session map[string]*Session
	lock    sync.RWMutex
}

func NewProvider() *Povider {
	return &Povider{
		session: make(map[string]*Session),
	}
}

func (p *Povider) SessionInit(sid string) (usecases.Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, err := p.session[sid]; err {
		return nil, domain.KeyAlreadyExists
	}
	p.session[sid] = &Session{
		sid:        sid,
		data:       make(map[string]string),
		lastAccess: time.Now(),
	}
	return p.session[sid], nil
}

func (p *Povider) SessionRead(sid string) (usecases.Session, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if _, err := p.session[sid]; !err {
		return nil, domain.Unauthorized
	}
	return p.session[sid], nil
}

func (p *Povider) SessionDestroy(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if _, err := p.session[sid]; !err {
		return domain.Unauthorized
	}
	delete(p.session, sid)
	return nil
}

func (p *Povider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for sid, sess := range p.session {
		if time.Now().Sub(sess.lastAccess) > time.Duration(maxLifeTime)*time.Second {
			delete(p.session, sid)
		}
	}
}
