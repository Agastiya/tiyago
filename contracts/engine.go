package contracts

type Engine interface {
	InitDatabase()
	Migrate()
	Serve()
}
