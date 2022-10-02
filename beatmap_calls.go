package gosuapiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (c Client) GetBeatmap(mapID uint) (BeatmapObject, error) {
	// Respect the ratelimit!
	c.rateLimiter.Take()

	var beatmap BeatmapObject
	resp, err := c.authenticatedClient.Get(BASE_URL + "/beatmaps/" + strconv.Itoa(int(mapID)))
	if err != nil {
		return beatmap, err
	}
	// check if the response status code is valid
	if resp.StatusCode != http.StatusOK {
		return beatmap, errors.New("invalid response code: " + resp.Status)
	}

	defer resp.Body.Close()

	// use json decoder to read the body
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&beatmap)
	if err != nil {
		return beatmap, err
	}

	return beatmap, nil
}

func (c Client) GetBeatmapWithMods(mapId uint, ruleset string, mods int64) (DifficultyAttribute, error) {
	// Respect the ratelimit!
	c.rateLimiter.Take()

	var difficulty DifficultyAttribute
	// we need to format the ruleset and mods into a json object in the format
	// {"ruleset": "osu", "mods": 72}
	// we can do this by creating a struct and then marshalling it into json
	type modRequest struct {
		Ruleset string `json:"ruleset"`
		Mods    int64  `json:"mods"`
	}
	request := modRequest{
		Ruleset: ruleset,
		Mods:    mods,
	}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return difficulty, err
	}

	jsonRequestReader := bytes.NewReader(jsonRequest)

	resp, err := c.authenticatedClient.Post(BASE_URL+"/beatmaps/"+strconv.Itoa(int(mapId))+"/attributes", "application/json", jsonRequestReader)
	if err != nil {
		return difficulty, err
	}

	// check if the response status code is valid
	if resp.StatusCode != http.StatusOK {
		return difficulty, errors.New("invalid response code: " + resp.Status)
	}

	defer resp.Body.Close()

	// use json decoder to read the body
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&difficulty)
	if err != nil {
		return difficulty, err
	}

	return difficulty, nil
}
