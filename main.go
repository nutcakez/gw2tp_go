package main

import (
	"encoding/json"
	"fmt"
	//"image/gif"
	"io/ioutil"
	"net/http"
	//"time"
)

func main() {
	giefMeMaterialz()

	
}


func giefMeMaterialz() {
	resp, err := http.Get("https://api.guildwars2.com/v2/materials/6")
		if err != nil {
			fmt.Println("oh boy")
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		stuff := materials{}
		json.Unmarshal(body, &stuff)

		fmt.Println(stuff)
}


type item_response struct {
	Id          int
	Whitelisted bool
	Buys        struct {
		Quantity   int
		Unit_price int
	}
	Sells struct {
		Quantity   int
		Unit_price int
	}
}


