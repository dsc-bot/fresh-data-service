package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type BotData struct {
	ID                    string    `json:"id"`
	Username              string    `json:"username"`
	Avatar                string    `json:"avatar"`
	Discriminator         string    `json:"discriminator"`
	PublicFlags           int       `json:"public_flags"`
	Bot                   bool      `json:"bot"`
	ApproximateGuildCount int       `json:"approximate_guild_count"`
	CreatedAt             time.Time `json:"createdAt"`
	CreatedTimestamp      int64     `json:"createdTimestamp"`
	PublicFlagsArray      []string  `json:"public_flags_array"`
	DefaultAvatarURL      string    `json:"defaultAvatarURL"`
	AvatarURL             string    `json:"avatarURL"`
}

type AppData struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Icon                string `json:"icon"`
	Description         string `json:"description"`
	Summary             string `json:"summary"`
	IsMonetized         bool   `json:"is_monetized"`
	IsVerified          bool   `json:"is_verified"`
	IsDiscoverable      bool   `json:"is_discoverable"`
	StorefrontAvailable bool   `json:"storefront_available"`
	BotPublic           bool   `json:"bot_public"`
	BotRequireCodeGrant bool   `json:"bot_require_code_grant"`
	TermsOfServiceURL   string `json:"terms_of_service_url"`
	PrivacyPolicyURL    string `json:"privacy_policy_url"`
}

type JAPIApplicationResponse struct {
	Data struct {
		Application AppData `json:"application"`
		Bot         BotData `json:"bot"`
	} `json:"data"`
}

func FetchApplication(application_id string) (*JAPIApplicationResponse, error) {
	resp, err := http.Get("https://japi.rest/discord/v1/application/" + application_id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response JAPIApplicationResponse
	res := json.NewDecoder(resp.Body)
	res.DisallowUnknownFields()
	res.Decode(&response)

	return &response, nil
}

type GuildData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Splash      string `json:"splash"`
	Banner      string `json:"banner"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type JAPIInviteResponse struct {
	Data struct {
		Code      string    `json:"code"`
		ExpiresAt time.Time `json:"expires_at"`
		Guild     GuildData `json:"guild"`
	} `json:"data"`
}

func FetchInvite(invite_code string) (*JAPIInviteResponse, error) {
	resp, err := http.Get("https://japi.rest/discord/v1/invite/" + invite_code)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response JAPIInviteResponse
	res := json.NewDecoder(resp.Body)
	res.DisallowUnknownFields()
	res.Decode(&response)

	return &response, nil
}
