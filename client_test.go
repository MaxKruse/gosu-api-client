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
