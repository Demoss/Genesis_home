package handlers

import (
	"encoding/json"
	"fmt"
	"genesis/global"
	"genesis/pkg"
	"net/http"
)

func GetBTC(w http.ResponseWriter, r *http.Request)  {
	 con := pkg.NewConnector()
	 if(global.Logged == 1) {
		 resp := con.GetBTC()
		 w.Write([]byte(fmt.Sprintf("1 %v = %v %v ", resp.Ticker.Base, resp.Ticker.Price, resp.Ticker.Target)))
	 } else {
	 	json.NewEncoder(w).Encode("You have to log in")
	 }
}

