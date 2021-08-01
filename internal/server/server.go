package server

import (
	"encoding/json"
	"fmt"
	"genesis/global"
	"genesis/pkg/resources"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"

	"regexp"
)


type Server struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Server)Start() error {
	if err := s.configureLogger(); err != nil{
		return err
	}


	s.logger.Info("Starting server")

	return nil
}

func (s *Server)configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil{
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func AddNewUser(user resources.User)  {
	rawDataIn, err := ioutil.ReadFile("users.json")
	if err != nil{
		println(err)
	}

	var settings resources.Settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil{
		println(err)
	}
	newUser := resources.User{user.Email,user.Pass}
	settings.Users = append(settings.Users, newUser)
	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")

	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile("users.json", rawDataOut, 0644)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}

func valid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	keys1, ok1 := r.URL.Query()["email"]
	keys2, ok2 := r.URL.Query()["pass"]
	if !ok1 || len(keys1[0]) == 0 {
		log.Println("Url Email is missing")
		json.NewEncoder(w).Encode("Please check params spelling")
		return
	}
	if !ok2 || len(keys2[0]) == 0 {
		log.Println("Url Password is missing")
		json.NewEncoder(w).Encode("Please check params spelling")
		return
	}

	key1 := keys1[0]
	if valid(key1){
		key2 := keys2[0]

		log.Println("Email " + string(key1) + " and pass " + string(key2))


		var decoder = schema.NewDecoder()

		var user resources.User

		err := decoder.Decode(&user,r.URL.Query())
		if err!= nil{
			fmt.Println(err)
			return
		}
			AddNewUser(user)
		w.Header().Set("context-type", "application/json")
		json.NewEncoder(w).Encode("User successfully created")

	} else {
		json.NewEncoder(w).Encode("Incorrect email")
		fmt.Println("Inccorect email")
	}
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request)  {
	rawDataIn, err := ioutil.ReadFile("users.json")
	if err != nil{
		println(err)
	}

	var settings resources.Settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil{
		println(err)
	}
	fmt.Println(settings)
	keys1, ok := r.URL.Query()["email"]
	keys2, ok := r.URL.Query()["pass"]
	if !ok || len(keys1[0]) < 1 {
		log.Println("Url Email 'key' is missing")
	}
	if !ok || len(keys2[0]) < 1 {
		log.Println("Url Password 'key' is missing")
	}

	var decoder = schema.NewDecoder()

	var user resources.User

	err = decoder.Decode(&user,r.URL.Query())
	if err!= nil{
		fmt.Println(err)
	}


	for i := 0; i < len(settings.Users); i++ {
		if(user.Email != settings.Users[i].Email || user.Pass != settings.Users[i].Pass ){
		global.Logged = 0
		} else{
			global.Logged = 1
			break
		}
	}
	if(global.Logged == 1){
		json.NewEncoder(w).Encode("You are logged in")
	} else{
		json.NewEncoder(w).Encode("User doesn`t exist")
	}
}

