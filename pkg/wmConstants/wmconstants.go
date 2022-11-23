package wmConstants

import "os"

const CURRENTWEATHER = "http://api.weatherapi.com/v1/current.json"

var APIKEY string = os.Getenv("APIKEY")
