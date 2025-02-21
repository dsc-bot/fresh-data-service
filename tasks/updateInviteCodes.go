package tasks

import "github.com/dsc-bot/fresh-data-service/utils"

func UpdateInviteCodes() {
	req, err := utils.FetchInvite("NzUYWsfe")
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to fetch japi application: %w", err)
		return
	}

	utils.Logger.Debug(req.Data.Guild.Name)
}
