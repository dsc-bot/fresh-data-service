package tasks

import (
	"context"
	"time"

	"github.com/dsc-bot/fresh-data-service/db"
	"github.com/dsc-bot/fresh-data-service/utils"
)

func UpdateBotData() {
	bots, err := db.GetBots(context.Background())
	if err != nil {
		utils.Logger.Sugar().Errorf("An error while fetching bots: %w", err)
		return
	}

	for _, bot := range bots {
		start := time.Now()

		UpdateBotListingData(&bot)
		err := db.UpdateBot(context.Background(), bot)
		if err != nil {
			utils.Logger.Sugar().Errorf("An error while updating bot %s: %w", bot.ListingId, err)
		}

		elapsed := time.Since(start)
		remaining := (1 * time.Second) - elapsed
		if remaining > 0 {
			utils.Logger.Sugar().Debugf("Finished with %s in %dms, waiting %dms...", bot.ListingId, elapsed.Milliseconds(), remaining.Milliseconds())
			time.Sleep(remaining)
		}
		break
	}

	utils.Logger.Sugar().Debug("Bot Data - Completed")
}

func UpdateBotListingData(bot *db.BotListing) {
	req, err := utils.FetchApplication(bot.AppId)
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to fetch japi application: %w", err)
		bot.Flags = append(bot.Flags, "FRESH_DATA_ERROR")
		return
	}

	if req.Data.Message == "Unknown Application" {
		utils.Logger.Sugar().Errorf("Unknown application for %s (%s)", bot.ListingId, bot.AppId)
		bot.Flags = append(bot.Flags, "UNKNOWN_APPLICATION", "FRESH_DATA_BLOCKED", "VISIBILITY_REDUCED")
		bot.Fetched = time.Now()
		return
	}

	bot.Username = req.Data.Bot.Username
	bot.Discriminator = req.Data.Bot.Discriminator
	bot.PrivacyPolicy = req.Data.Application.PrivacyPolicyURL
	bot.TermsOfService = req.Data.Application.TermsOfServiceURL
	bot.Store = req.Data.Application.StorefrontAvailable
	bot.Servers = req.Data.Bot.ApproximateGuildCount
	bot.Flags = utils.RemoveStrings(bot.Flags, "FRESH_DATA_ERROR", "FRESH_DATA_BLOCKED")
	bot.Fetched = time.Now()
}
