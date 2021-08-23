package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	gorilla_context "github.com/gorilla/context"
)

type UserInfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AuthMiddleware struct {
	es *ElasticHelper
	db *DBHelper
}

func initAuthMiddleware(es *ElasticHelper, db *DBHelper) *AuthMiddleware {

	au := AuthMiddleware{es, db}

	return &au
}

func TimeTrack(name string, start time.Time) {
	elapsed := time.Since(start)

	log.Trace(fmt.Sprintf("%s took %s", name, elapsed))
}

func (auth *AuthMiddleware) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer TimeTrack("Processing "+r.URL.Path, time.Now())

		err := auth.validateSession(w, r)
		if err == nil {
			next.ServeHTTP(w, r)
		}
	})
}

func (auth *AuthMiddleware) validateSession(w http.ResponseWriter, r *http.Request) error {

	sessionCookie, err := r.Cookie("sessionid")
	if err != nil {
		err = fmt.Errorf("Failed to retrieve sessionid cookie: %v", err)
		log.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	sessionKey := sessionCookie.Value

	log.Trace("Getting user info for sessionid ", sessionKey)

	var encodedSessionData string
	err = auth.db.conn.QueryRow(context.Background(), "select session_data from django_session where session_key=$1", sessionKey).Scan(&encodedSessionData)
	if err != nil {
		err = fmt.Errorf("Session does not exist: %v", err)
		log.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	decodedSessionData, err := base64.StdEncoding.DecodeString(encodedSessionData)
	if err != nil {
		err = fmt.Errorf("Failed to decode session data: %v", err)
		log.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	re := regexp.MustCompile(`[^{]*({.*})$`)
	jsonSessionData := re.ReplaceAllString(string(decodedSessionData), "${1}")

	log.Printf("SessionData: %v\n", jsonSessionData)
	var sessionData struct {
		UserID int `json:"_auth_user_id"`
	}
	json.Unmarshal([]byte(jsonSessionData), &sessionData)

	log.Printf("UserID: %v\n", sessionData.UserID)

	ui := UserInfo{Id: sessionData.UserID}
	udi, err := auth.db.getUserInfo(ui.Id)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return err
	}

	ui.Name = udi.first_name + " " + udi.last_name
	ui.Email = udi.email

	if !udi.verified {
		err = fmt.Errorf("User %s (%s) is not verified yet", ui.Name, ui.Email)
		http.Error(w, err.Error(), http.StatusForbidden)
		return err
	}

	if udi.blocked {
		err = fmt.Errorf("User %s (%s) is blocked", ui.Name, ui.Email)
		http.Error(w, err.Error(), http.StatusForbidden)
		return err
	}

	log.Info("User ", ui)

	gorilla_context.Set(r, "UserInfo", ui)

	return nil
}

func (auth *AuthMiddleware) getUserInfo(r *http.Request) UserInfo {
	return gorilla_context.Get(r, "UserInfo").(UserInfo)
}