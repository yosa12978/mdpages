package session

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/yosa12978/mdpages/types"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func GetKey(r *http.Request, key string) (any, error) {
	session, err := store.Get(r, "user_store")
	if err != nil { // without this may throw null pointer exception
		return nil, err
	}
	return session.Values[key], err
}

func SetKey(r *http.Request, w http.ResponseWriter, key string, value any) error {
	session, err := store.Get(r, "user_store")
	if err != nil {
		return err
	}
	session.Values[key] = value
	return session.Save(r, w)
}

func GetInfo(r *http.Request) (*types.SessionInfo, error) {
	session, err := store.Get(r, "user_store")
	if err != nil {
		return nil, err
	}
	userval := session.Values["account"]
	if userval == nil {
		session.Values["account"] = nil
		return nil, errors.New("user is not logged in")
	}
	var info types.SessionInfo
	err = json.Unmarshal([]byte(userval.(string)), &info)
	return &info, err
}

func SetInfo(w http.ResponseWriter, r *http.Request, account *types.Account) error {
	session, err := store.Get(r, "user_store")
	if err != nil {
		return err
	}
	if account == nil {
		session.Values["account"] = nil
		return session.Save(r, w)
	}

	sessionInfo := types.SessionInfo{
		Username:  account.Username,
		Groups:    account.Groups,
		Timestamp: time.Now().UnixNano(),
		LoggedIn:  true,
	}
	acc, err := json.Marshal(sessionInfo)
	if err != nil {
		return err
	}
	session.Values["account"] = string(acc)
	return session.Save(r, w)
}

func SetDefault(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "user_store")
	if err != nil {
		return err
	}
	s := types.SessionInfo{
		Username:  "",
		Groups:    []types.Group{},
		Timestamp: 0,
		LoggedIn:  false,
	}
	acc, err := json.Marshal(s)
	if err != nil {
		return err
	}
	session.Values["account"] = acc
	return nil
}
