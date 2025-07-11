package module

import (
	"github.com/agastiya/tiyago/controller"
	"github.com/agastiya/tiyago/repository"
	"github.com/agastiya/tiyago/service"
	"gorm.io/gorm"
)

type AppModule struct {
	Controller controller.Controller
}

func InitModule(db *gorm.DB) AppModule {

	repos := repository.InitRepos(db)
	services := service.InitServices(repos)
	ctrl := controller.InitController(*services)

	return AppModule{
		Controller: *ctrl,
	}
}
