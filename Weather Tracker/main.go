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
	bytes , err := ioutil.ReadFile(fileName)

	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return apiConfigData{}, err
	}
	
	return config, nil
}


func main (){
	config, err := loadApiConfig("config.json")

	if err != nil {
		panic(err)
	}


	http.HandleFunc("/ping", hello)

	http.HandleFunc("/weather/", getCityWeatherData)


	http.ListenAndServe(":8080", nil)
}

func getCityWeatherData(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
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
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=1b1c0d6f1b6f9c7e3c0f1b6f9c7e3c0", city)
	resp, err := http.Get(url)

	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}