package gosuapiclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type UserObject struct {
	AvatarURL                string                 `json:"avatar_url"`
	CountryCode              string                 `json:"country_code"`
	DefaultGroup             string                 `json:"default_group"`
	ID                       int                    `json:"id"`
	IsActive                 bool                   `json:"is_active"`
	IsBot                    bool                   `json:"is_bot"`
	IsDeleted                bool                   `json:"is_deleted"`
	IsOnline                 bool                   `json:"is_online"`
	IsSupporter              bool                   `json:"is_supporter"`
	LastVisit                time.Time              `json:"last_visit"`
	PmFriendsOnly            bool                   `json:"pm_friends_only"`
	ProfileColour            string                 `json:"profile_colour"`
	Username                 string                 `json:"username"`
	CoverURL                 string                 `json:"cover_url"`
	Discord                  string                 `json:"discord"`
	HasSupported             bool                   `json:"has_supported"`
	Interests                interface{}            `json:"interests"`
	JoinDate                 time.Time              `json:"join_date"`
	Kudosu                   Kudosu                 `json:"kudosu"`
	Location                 interface{}            `json:"location"`
	MaxBlocks                int                    `json:"max_blocks"`
	MaxFriends               int                    `json:"max_friends"`
	Occupation               interface{}            `json:"occupation"`
	Playmode                 string                 `json:"playmode"`
	Playstyle                []string               `json:"playstyle"`
	PostCount                int                    `json:"post_count"`
	ProfileOrder             []string               `json:"profile_order"`
	Title                    interface{}            `json:"title"`
	Twitter                  string                 `json:"twitter"`
	Website                  string                 `json:"website"`
	Country                  Country                `json:"country"`
	Cover                    Cover                  `json:"cover"`
	IsRestricted             bool                   `json:"is_restricted"`
	AccountHistory           []interface{}          `json:"account_history"`
	ActiveTournamentBanner   interface{}            `json:"active_tournament_banner"`
	Badges                   []Badges               `json:"badges"`
	FavouriteBeatmapsetCount int                    `json:"favourite_beatmapset_count"`
	FollowerCount            int                    `json:"follower_count"`
	GraveyardBeatmapsetCount int                    `json:"graveyard_beatmapset_count"`
	Groups                   []Groups               `json:"groups"`
	LovedBeatmapsetCount     int                    `json:"loved_beatmapset_count"`
	MonthlyPlaycounts        []MonthlyPlaycounts    `json:"monthly_playcounts"`
	Page                     Page                   `json:"page"`
	PendingBeatmapsetCount   int                    `json:"pending_beatmapset_count"`
	PreviousUsernames        []interface{}          `json:"previous_usernames"`
	RankedBeatmapsetCount    int                    `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts     []ReplaysWatchedCounts `json:"replays_watched_counts"`
	ScoresFirstCount         int                    `json:"scores_first_count"`
	Statistics               Statistics             `json:"statistics"`
	SupportLevel             int                    `json:"support_level"`
	UserAchievements         []UserAchievements     `json:"user_achievements"`
	RankHistory              RankHistory            `json:"rank_history"`
}
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Cover struct {
	CustomURL string      `json:"custom_url"`
	URL       string      `json:"url"`
	ID        interface{} `json:"id"`
}
type Badges struct {
	AwardedAt   time.Time `json:"awarded_at"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	URL         string    `json:"url"`
}
type Groups struct {
	ID          int    `json:"id"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	Colour      string `json:"colour"`
}
type MonthlyPlaycounts struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}
type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}
type ReplaysWatchedCounts struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}
type Level struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}
type GradeCounts struct {
	Ss  int `json:"ss"`
	SSH int `json:"ssh"`
	S   int `json:"s"`
	Sh  int `json:"sh"`
	A   int `json:"a"`
}
type Rank struct {
	Global  int `json:"global"`
	Country int `json:"country"`
}
type Statistics struct {
	Level                  Level       `json:"level"`
	Pp                     float64     `json:"pp"`
	GlobalRank             int         `json:"global_rank"`
	RankedScore            int         `json:"ranked_score"`
	HitAccuracy            float64     `json:"hit_accuracy"`
	PlayCount              int         `json:"play_count"`
	PlayTime               int         `json:"play_time"`
	TotalScore             int         `json:"total_score"`
	TotalHits              int         `json:"total_hits"`
	MaximumCombo           int         `json:"maximum_combo"`
	ReplaysWatchedByOthers int         `json:"replays_watched_by_others"`
	IsRanked               bool        `json:"is_ranked"`
	GradeCounts            GradeCounts `json:"grade_counts"`
	Rank                   Rank        `json:"rank"`
}
type UserAchievements struct {
	AchievedAt    time.Time `json:"achieved_at"`
	AchievementID int       `json:"achievement_id"`
}
type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

func (c Client) GetUserFromUsername(username string) (UserObject, error) {
	// Respect the ratelimit!
	c.rateLimiter.Take()

	var user UserObject
	endpoint := BASE_URL + "/users/" + username + "/osu?key=username"
	resp, err := c.authenticatedClient.Get(endpoint)
	if err != nil {
		return user, err
	}
	// check if the response status code is valid
	if resp.StatusCode != http.StatusOK {
		return user, errors.New("invalid response code: " + resp.Status)
	}

	defer resp.Body.Close()

	// use json decoder to read the body
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (c Client) GetUserFromID(id int) (UserObject, error) {
	// Respect the ratelimit!
	c.rateLimiter.Take()

	var user UserObject
	resp, err := c.authenticatedClient.Get(BASE_URL + "/users/" + strconv.Itoa(id))
	if err != nil {
		return user, err
	}
	// check if the response status code is valid
	if resp.StatusCode != http.StatusOK {
		return user, errors.New("invalid response code: " + resp.Status)
	}

	defer resp.Body.Close()

	// use json decoder to read the body
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
