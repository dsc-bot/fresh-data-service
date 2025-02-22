package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dsc-bot/fresh-data-service/config"
	"github.com/dsc-bot/fresh-data-service/db"
	"github.com/dsc-bot/fresh-data-service/tasks"
	"github.com/dsc-bot/fresh-data-service/utils"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	config.Parse()

	lerr := utils.Configure(nil, config.Conf.JsonLogs, config.Conf.LogLevel)
	if lerr != nil {
		panic(fmt.Errorf("failed to create zap logger: %w", lerr))
	}
	db.Init()

	s, s_err := gocron.NewScheduler()
	if s_err != nil {
		panic(fmt.Errorf("failed to create cron scheduler: %w", s_err))
	}

	if config.Conf.OneShot {
		utils.Logger.Info("Running Job - Bot Data")
		tasks.UpdateBotData()
		utils.Logger.Info("Running Job - Invite Data")
		tasks.UpdateInviteCodes()
		return
	}

	// add hourly fresh bot data
	s.NewJob(gocron.CronJob("0 * * * *", false), gocron.NewTask(func() {
		utils.Logger.Debug("Running Hourly Job - Bot Data")
		tasks.UpdateBotData()
	}))

	// add daily fresh invite data
	s.NewJob(gocron.CronJob("0 0 * * *", false), gocron.NewTask(func() {
		utils.Logger.Debug("Running Daily Job - Invite Data")
		tasks.UpdateInviteCodes()
	}))

	// start the scheduler
	utils.Logger.Info("Starting cron jobs")
	s.Start()

	// keep alive until shutdown signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownCh

	utils.Logger.Info("Received shutdown signal")
}
