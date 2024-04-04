package repository

type Worker interface {
	Submit(task func()) error
	Close() error
}
