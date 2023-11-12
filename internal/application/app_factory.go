package application

import (
	"github.com/thumperq/wms/user-clean-up-job/internal/infrastructure/db"
)

type AppFactory struct {
	UserApp UserApp
}

func NewApplicationFactory(DbFactory db.DbFactory) AppFactory {
	return AppFactory{
		UserApp: NewUserApp(DbFactory),
	}
}
