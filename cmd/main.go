package main

import (
	"context"
	"os"
	"runtime"

	"github.com/thumperq/golib/config"
	"github.com/thumperq/wms/user-clean-up-job/internal/application"
	"github.com/thumperq/wms/user-clean-up-job/internal/infrastructure/db"
)

func main() {
	ctx := context.Background()
	configManager := config.NewConfigManager()
	dbFactory := db.NewDbFactory(configManager)
	appFactory := application.NewApplicationFactory(dbFactory)
	err := appFactory.UserApp.CleanupUsers(ctx)
	if err != nil {
		os.Exit(1)
	}
	runtime.Goexit()
}
