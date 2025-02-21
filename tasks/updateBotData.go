package tasks

import (
	"github.com/dsc-bot/fresh-data-service/utils"
)

func UpdateBotData() {
	req, err := utils.FetchApplication("861054036328710164")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to fetch japi application: %w", err)
		return
	}

	utils.Logger.Debug(req.Data.Bot.Username)
}
