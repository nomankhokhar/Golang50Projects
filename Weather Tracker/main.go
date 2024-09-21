package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


type apiConfigData struct {
	OpenWeatherMapApiKey string
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
func loadApiConfig(fileName string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return apiConfigData{}, err
	}

	// Replace Windows-style line endings (\r\n) with Unix-style line endings (\n)
	fileContent := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	var config apiConfigData

	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		return apiConfigData{}, err
	}

	return config, nil
}

func main (){

	http.HandleFunc("/ping", hello)

	http.HandleFunc("/weather/", getCityWeatherData)


	http.ListenAndServe(":8080", nil)
}

func getCityWeatherData(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	fmt.Println(city)
	data , err := query(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}


func hello (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello!"))	
}

func query(city string) (weatherData, error) {
	apiConfig , err := loadApiConfig(".apiConfig")
	fmt.Println(apiConfig.OpenWeatherMapApiKey)

	fmt.Println(err)
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	fmt.Println("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	if err != nil {
		return weatherData{}, err
	}	

	defer resp.Body.Close()

	var data weatherData

	json.NewDecoder(resp.Body).Decode(&data)

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}