package gosuapiclient_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	gosuapiclient "github.com/maxkruse/gosu-api-client"
	"golang.org/x/oauth2"
)

var token oauth2.Token
var err error

func init() {
	token, err = ReadTokenFromFile("token.json")
	if err != nil {
		panic(err)
	}
}

func TestGetBeatmap(t *testing.T) {
	client := gosuapiclient.NewClient(token, context.Background())

	res, err := client.GetBeatmap(3322364)
	if err != nil {
		t.Error(err)
		return
	}

	if res.ID != 3322364 {
		t.Error("ID is not 3322364, but " + strconv.Itoa(res.ID))
	}

	if res.Beatmapset.Title != "Shinjuku Jack" {
		t.Error("Title is not Shinjuku Jack, but " + res.Beatmapset.Title)
	}

	if res.DifficultyRating <= 7.09 || res.DifficultyRating >= 7.11 {
		t.Error("Difficulty rating is not 7.1, but " + strconv.FormatFloat(res.DifficultyRating, 'f', 1, 64))
	}
}

func TestGetBeatmapWithMods(t *testing.T) {
	client := gosuapiclient.NewClient(token, context.Background())

	res, err := client.GetBeatmapWithMods(3322364, "osu", gosuapiclient.HardRock)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Attributes.StarRating <= 7.69 || res.Attributes.StarRating >= 7.71 {
		t.Error("Star rating is not 7.7, but " + strconv.FormatFloat(res.Attributes.StarRating, 'f', 1, 64))
	}

	if res.Attributes.ApproachRate != 10.0 {
		t.Error("Approach rate is not 10.0, but " + strconv.FormatFloat(res.Attributes.ApproachRate, 'f', 1, 64))
	}

	if res.Attributes.MaxCombo != 2702 {
		t.Error("Max combo is not 2702, but " + strconv.Itoa(res.Attributes.MaxCombo))
	}
}

func TestGetUserFromUsername(t *testing.T) {
	client := gosuapiclient.NewClient(token, context.Background())

	res, err := client.GetUserFromUsername("peppy")
	if err != nil {
		t.Error(err)
		return
	}

	if res.ID != 2 {
		t.Error("ID is not 2, but " + strconv.Itoa(res.ID))
	}
}

func TestGetUserFromI(t *testing.T) {
	client := gosuapiclient.NewClient(token, context.Background())

	res, err := client.GetUserFromID(2)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Username != "peppy" {
		t.Error("Username is not peppy, but " + res.Username)
	}
}

func ReadTokenFromFile(file string) (oauth2.Token, error) {
	// open the file
	f, err := os.Open(file)
	if err != nil {
		return oauth2.Token{}, err
	}

	// read the file
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return oauth2.Token{}, err
	}

	// unmarshal the json
	var token oauth2.Token
	err = json.Unmarshal(b, &token)
	if err != nil {
		return oauth2.Token{}, err
	}

	return token, nil
}
