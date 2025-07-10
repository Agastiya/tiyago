package module

import (
	"github.com/agastiya/tiyago/controller"
	"github.com/agastiya/tiyago/repository/user"
	"github.com/agastiya/tiyago/service"
	"gorm.io/gorm"
)

type AppModule struct {
	Controller controller.Controller
}

func InitModules(db *gorm.DB) AppModule {
	userRepo := user.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	serviceAggregator := service.Service{
		User: userService,
	}

	return AppModule{
		Controller: controller.NewController(serviceAggregator),
	}
}
