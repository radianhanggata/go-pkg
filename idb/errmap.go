package idb

type ErrorMap interface {
	Read(errin error) error
}
