package main

import (
	"encoding/json"
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

// CORS middleware
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loadApiConfig(fileName string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return apiConfigData{}, err
	}

	fileContent := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	var config apiConfigData
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		return apiConfigData{}, err
	}

	return config, nil
}

func main() {
	// Create a new mux (router)
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", hello)
	mux.HandleFunc("/weather/", getCityWeatherData)

	// Wrap the mux with CORS middleware
	http.ListenAndServe(":8080", enableCors(mux))
}

func getCityWeatherData(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := query(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
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
