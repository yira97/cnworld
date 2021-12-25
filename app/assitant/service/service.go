package service

type CreateReady interface {
	Clean4Create()
	Check4Create() (e error)
	Init4Create()
}
