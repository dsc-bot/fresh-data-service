package main

import (
	"fmt"

	"github.com/dsc-bot/fresh-data-service/config"
	"github.com/dsc-bot/fresh-data-service/tasks"
	"github.com/dsc-bot/fresh-data-service/utils"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	config.Parse()

	err := utils.Configure(nil, config.Conf.JsonLogs, config.Conf.LogLevel)
	if err != nil {
		panic(fmt.Errorf("failed to create zap logger: %w", err))
	}

	s, s_err := gocron.NewScheduler()
	if s_err != nil {
		panic(fmt.Errorf("failed to create cron scheduler: %w", s_err))
	}

	// // debug job
	// s.NewJob(gocron.CronJob("* * * * *", false), gocron.NewTask(func() {
	// 	logger.Debug("Running Minute Job - Debugging")
	// }))

	// // add hourly fresh bot data
	// s.NewJob(gocron.CronJob("0 * * * *", false), gocron.NewTask(func() {
	// 	logger.Debug("Running Hourly Job - Bot Data")
	// 	tasks.UpdateBotData()
	// }))

	// add daily fresh invite data
	// s.NewJob(gocron.CronJob("0 0 * * *", false), gocron.NewTask(func() {
	// 	logger.Debug("Running Daily Job - Invite Data")
	// 	tasks.UpdateInviteCodes()
	// }))

	// start the scheduler
	utils.Logger.Info("Starting cron jobs")
	s.Start()

	tasks.UpdateBotData()

	// keep alive until shutdown signal
	// shutdownCh := make(chan os.Signal, 1)
	// signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	// <-shutdownCh

	// logger.Info("Received shutdown signal")
}
