package usecases

type Session interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	SessionID() string
}
