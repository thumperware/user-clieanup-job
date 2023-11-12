package application

import (
	"context"

	"github.com/thumperq/wms/user-clean-up-job/internal/infrastructure/db"
)

type UserApp struct {
	dbFactory db.DbFactory
}

func NewUserApp(DbFactory db.DbFactory) UserApp {
	app := UserApp{
		dbFactory: DbFactory,
	}
	return app
}

func (app UserApp) CleanupUsers(ctx context.Context) error {
	return app.dbFactory.UserDb.DeleteUsers(ctx)
}
