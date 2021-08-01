package main

import (
	"flag"
	"genesis/global"
	"genesis/internal/api"
	"genesis/internal/server"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

var LoggedIn int

var (
	configPath string
)

func InitApp() (int){
	return 0
}

func init()  {
	flag.StringVar(&configPath, "config-path","configs/server.toml","path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath,config)
	if err !=nil{
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
	r := api.NewRouter()

	if(global.Logged == 0){
		r = api.NewRouter()
	} else {
		r = api.NewRouter()
	}

	http.ListenAndServe(":8000", r)


}























//r := api.NewRouter()
//http.ListenAndServe(":8000", r)