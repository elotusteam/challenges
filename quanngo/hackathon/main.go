package main

import (
	"TEST/quanngo/hackathon/config"
	authen "TEST/quanngo/hackathon/jwt"
	"TEST/quanngo/hackathon/services"
	"TEST/quanngo/hackathon/storage/postgres"
	"TEST/quanngo/hackathon/types"
	"TEST/quanngo/hackathon/utilities"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func main() {
	_, err := postgres.New("localhost", "5432", "postgres", "secret", "hackathonDB")
	if err != nil {
		fmt.Println("postgres error!")
		return
	}

	config.LoadConfig()

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/callback", handleCallBack)
	http.HandleFunc("/upload", handleUpload)

	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	err := authen.Validate(w, r)
	if err != nil {
		return
	}

}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	args := types.Register{}
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := services.CreateUser(&args)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {

	err := authen.Validate(w, r)
	if err != nil {
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil && err.Error() != "http: no such file" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	args := types.FileInfo{}

	if file != nil {
		fileType, fileName := utilities.CheckingFileType(handler.Filename)
		if fileName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer file.Close()

		args.FileName = fileName
		args.Size = handler.Size
		args.Type = fileType
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.CreateFile(file, &args)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(config.RandomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallBack(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != config.RandomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := config.AppConfig.GoogleLoginConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if content == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 50)

	user := &types.UserInfo{}
	if err := json.Unmarshal(content, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := &authen.Claims{
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	svToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := svToken.SignedString(authen.JwtKey)
	if err != nil {
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	err = services.CheckUser(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	utilities.BuildResponse(w, r, "succeed", tokenString)
}
