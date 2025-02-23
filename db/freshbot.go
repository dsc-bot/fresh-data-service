package db

import (
	"context"
	"time"
)

type BotListing struct {
	ListingId      string
	AppId          string
	ClientId       string
	Flags          []string
	Username       string
	Discriminator  string
	Avatar         *string
	Banner         *string
	PrivacyPolicy  string
	TermsOfService string
	Servers        int
	Store          bool
	Fetched        time.Time
}

func GetBots(ctx context.Context) ([]BotListing, error) {
	query := `SELECT 
			"listingId", "appId", "clientId", "flags",
			"username", "discriminator", "avatar", "banner", 
			"privacyPolicy", "termsOfService", "servers", 
			"store", "fetched"
		FROM "Bot" 
		WHERE NOT ('FRESH_DATA_BLOCKED' = ANY(flags)) 
		ORDER BY "fetched" ASC 
		LIMIT 1800;`

	rows, err := Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []BotListing
	for rows.Next() {
		var listing BotListing
		err := rows.Scan(
			&listing.ListingId, &listing.AppId, &listing.ClientId, &listing.Flags,
			&listing.Username, &listing.Discriminator, &listing.Avatar, &listing.Banner,
			&listing.PrivacyPolicy, &listing.TermsOfService, &listing.Servers,
			&listing.Store, &listing.Fetched,
		)
		if err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listings, nil
}

func UpdateBot(ctx context.Context, bot BotListing) (err error) {
	query := `UPDATE "Bot" 
		SET 
			"username" = $2, "discriminator" = $3, "avatar" = $4,
			"banner" = $5, "privacyPolicy" = $6, "termsOfService" = $7,
			"servers" = $8, "store" = $9, "fetched" = $10,
			"flags" = $11
		WHERE "appId" = $1;` //"listingId" = $2

	_, err = Pool.Exec(ctx, query,
		bot.AppId,
		bot.Username, bot.Discriminator, bot.Avatar, bot.Banner, bot.PrivacyPolicy, bot.TermsOfService, bot.Servers, bot.Store, bot.Fetched, bot.Flags)
	return err
}
