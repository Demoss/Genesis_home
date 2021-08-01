package pkg

import (
	"encoding/json"
	"fmt"
	"genesis/pkg/resources"
	"net/http"
)



type Connector struct {

}

func NewConnector() *Connector {
	return &Connector{}
}

func (c *Connector) GetBTC() *resources.ResponceBTC {

	res, err := http.Get("https://api.cryptonator.com/api/ticker/btc-uah")
	if err != nil {
		fmt.Println("Failed")
		return nil
	}
	var resp resources.ResponceBTC
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil{
		fmt.Println(err)
	}
	return &resp


}

