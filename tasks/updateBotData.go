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

	var updates int = 0
	var errors int = 0
	var dberrors int = 0

	for _, bot := range bots {
		start := time.Now()

		success := UpdateBotListingData(&bot)
		if success {
			updates += 1
		} else {
			errors += 1
		}

		err := db.UpdateBot(context.Background(), bot)
		if err != nil {
			dberrors += 1
			utils.Logger.Sugar().Errorf("An error while updating bot %s: %w", bot.ListingId, err)
		}

		if updates > 5 && errors > 5 && float64(errors)/float64(updates) > 0.5 {
			utils.Logger.Sugar().Errorf("Too many errors (updated %d, api error %d, db errors, %d), stopping...", updates, errors, dberrors)
			break
		}

		if updates > 5 && dberrors > 5 && float64(dberrors)/float64(updates) > 0.5 {
			utils.Logger.Sugar().Errorf("Too many database errors (updated %d, api error %d, db errors, %d), stopping...", updates, errors, dberrors)
			break
		}

		elapsed := time.Since(start)
		remaining := (1 * time.Second) - elapsed
		if remaining > 0 {
			utils.Logger.Sugar().Debugf("Finished with %s in %dms, waiting %dms...", bot.ListingId, elapsed.Milliseconds(), remaining.Milliseconds())
			time.Sleep(remaining)
		}
	}

	utils.Logger.Sugar().Infof("Bot Data - Updated %d/%d - Errored %d - DB Errors %d", updates, len(bots), errors, dberrors)
}

func UpdateBotListingData(bot *db.BotListing) bool {
	req, err := utils.FetchApplication(bot.AppId)
	bot.Fetched = time.Now().UTC()
	if err != nil {
		utils.Logger.Sugar().Errorf("failed to fetch japi application: %w", err)
		bot.Flags = append(bot.Flags, "FRESH_DATA_ERROR")
		return false
	}

	if req.Data.Message == "Unknown Application" {
		utils.Logger.Sugar().Errorf("Unknown application for %s (%s)", bot.ListingId, bot.AppId)
		bot.Flags = append(bot.Flags, "UNKNOWN_APPLICATION", "FRESH_DATA_BLOCKED", "VISIBILITY_REDUCED")
		return false
	}

	bot.Username = req.Data.Bot.Username
	bot.Discriminator = req.Data.Bot.Discriminator
	bot.PrivacyPolicy = req.Data.Application.PrivacyPolicyURL
	bot.TermsOfService = req.Data.Application.TermsOfServiceURL
	bot.Store = req.Data.Application.StorefrontAvailable
	bot.Servers = req.Data.Bot.ApproximateGuildCount
	bot.Flags = utils.RemoveStrings(bot.Flags, "FRESH_DATA_ERROR", "FRESH_DATA_BLOCKED")
	return true
}
