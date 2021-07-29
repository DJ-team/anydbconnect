package driver

type DB interface {
	Connect(conn string)
}
