package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const OPEN_WEATHER_MAP = "https://api.openweathermap.org/data/2.5";

type CloudsType struct {
  All float64 `json:"all"`
}

type CoordinateType struct {
  Longitude float64 `json:"lon"`
  Latitude float64 `json:"lat"`
}

type MainType struct {
  Temperature float64 `json:"temp"`
  FeelsLike float64 `json:"feels_like"`
  TemperatureMin float64 `json:"temp_min"`
  TemperatureMax float64 `json:"temp_max"`
  Pressure float64 `json:"pressure"`
  Humidity float64 `json:"humidity"`
  SeaLevel float64 `json:"sea_level"`
  GrandLevel float64 `json:"grnd_level"`
}

type SystemType struct {
  Type int `json:"type"`
  Id int `json:"id"`
  Country string `json:"country"`
  Sunrise float64 `json:"sunrise"`
  Sunset float64 `json:"sunset"`
}

type WeatherType struct {
  Id int `json:"id"`
  Main string `json:"main"`
  Description string `json:"description"`
  Icon string `json:"icon"`
}

type WindType struct {
  Speed float64 `json:"speed"`
  Degree float64 `json:"deg"`
  Gust float64 `json:"gust"`
}

type OpenWeatherMap struct {
  Base string `json:"base"`
  Clouds CloudsType `json:"clouds"`
  Coordinate CoordinateType `json:"coord"`
  DateTime int `json:"dt"`
  Id int `json:"id"`
  Main MainType `json:"main"`
  Name string `json:"name"`
  System SystemType `json:"sys"`
  TimeZone int `json:"timezone"`
  Weather []WeatherType `json:"weather"`
  Wind WindType `json:"wind"`
}

func GetWeather(appid string, q string, units string) (*OpenWeatherMap, error) {
  // Assign default value
  if q == "" {
    q = "hanoi"
  }

  // Build query parameters
  var queryParameters []string = []string{}
  if appid != "" {
    queryParameters = append(queryParameters, "appid=" + appid)
  }
  if q != "" {
    queryParameters = append(queryParameters, "q=" + q)
  }
  if units != "" {
    queryParameters = append(queryParameters, "units=" + units)
  }

  // Build url
  var url = fmt.Sprintf(
    "%s/weather?%s",
    OPEN_WEATHER_MAP,
    strings.Join(queryParameters, "&"),
  )

  body, getError := Get(url)
  if getError != nil {
    return nil, getError
  }

  var openWeatherMap *OpenWeatherMap
  jsonUnmarshalError := json.Unmarshal(body, &openWeatherMap)
  if jsonUnmarshalError != nil {
    return nil, jsonUnmarshalError
  }

  return openWeatherMap, nil
}

func Get(url string) ([]byte, error){
  response, httpGetError := http.Get(url)
  if httpGetError != nil {
    return nil, httpGetError
  }
  defer response.Body.Close()

  body, readBodyError := ioutil.ReadAll(response.Body)
  if readBodyError != nil {
    return nil, readBodyError
  }

  return body, nil
}
