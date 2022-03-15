package pkg

type Storage interface {
	SaveUser(*User) error
}
