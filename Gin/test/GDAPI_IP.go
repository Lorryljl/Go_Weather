package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//url := "https://restapi.amap.com/v3/ip?&output=xml&key=049a159604b07b805722006015eedf50&" //ip
	url := "https://restapi.amap.com/v3/weather/weatherInfo?city=440605&key=049a159604b07b805722006015eedf50&extensions=base" //天气
	response, err := http.Get(url)
	if err != nil {

		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
