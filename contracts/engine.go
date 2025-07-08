package contracts

type Engine interface {
	InitDatabase()
	Migrate()
	InitPackage()
	Serve()
}
