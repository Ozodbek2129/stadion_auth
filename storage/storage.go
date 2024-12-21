package storage

type IStorage interface {
	User() IUserStorage
	Close()
}

type IUserStorage interface {
}
