package imap

type ErrorMap interface {
	Read(errin error) error
}
