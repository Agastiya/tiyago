package contracts

type Engine interface {
	InitDatabase()
	InitCommand()
	Serve()
}
