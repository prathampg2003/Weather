package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	//simple Program that does Json parsing from the weather API

	fmt.Println("Enter Your State:")

	reader := bufio.NewReader(os.Stdin)
	state, _ := reader.ReadString('\n')
	state = strings.TrimSpace(state)
	state = strings.ReplaceAll(state, " ", "%20")

	get := "https://api.weatherapi.com/v1/forecast.json?key=ec180872243c4f57a4f153631230105&q=" + state + "" + "&days=1&aqi=no&alerts=no"

	res, err := http.Get(get)
	if err != nil {
		panic("ERROR IN GET API")
	}

	if res.StatusCode != 200 {
		panic("ERROR STATUS IN API")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic("Error in reading")
	}
	var myOnlineData map[string]interface{}
	json.Unmarshal(body, &myOnlineData)

	location := myOnlineData["location"].(map[string]interface{})
	fmt.Println("Location:-", location["name"])
	fmt.Println("Region:-", location["region"])
	current := myOnlineData["current"].(map[string]interface{})
	fmt.Println("Temperature", current["temp_c"], "'C")

	defer res.Body.Close()

}
