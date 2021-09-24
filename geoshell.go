package main

import (
	"bufio"
	"fmt"
	"net/http"
	"flag"
	"log"
)

var WEATHER_URL = "https://weather.cc.api.here.com/weather/1.0/report.json"
var WEATHER_ALERTS = "product=alerts"

func main() {

	apikeyPtr := flag.String("apikey", "", "Your API key for HERE Maps")

	flag.Parse()

	if *apikeyPtr == "" {
		log.Fatalf("You must supply an API Key!")
	}
	
	var city = "name=Boston"
	resp, err := http.Get(WEATHER_URL+"?"+
		WEATHER_ALERTS+"&"+
		city+"&"+
		"apiKey="+*apikeyPtr)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
