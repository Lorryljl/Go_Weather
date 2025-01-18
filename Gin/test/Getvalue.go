package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	"io"
	"net/http"
)

type WeatherResponse struct {
	Status    string     `json:"status"`
	Count     string     `json:"count"`
	Info      string     `json:"info"`
	Infocode  string     `json:"infocode"`
	Forecasts []Forecast `json:"forecasts"`
}
type Forecast struct {
	City       string
	Adcode     string
	Province   string
	ReportTime string
	Casts      []Cast
}
type Cast struct {
	Date           string
	Week           string
	DayWeather     string
	NightWeather   string
	DayTemp        string
	NightTemp      string
	DayWind        string
	NightWind      string
	DayPower       string
	NightPower     string
	DayTempFloat   float64
	NightTempFloat float64
}

func main() {
	r := gin.Default()
	r.Static("/static", "../web/static")
	r.LoadHTMLGlob("../website/*")

	// 登录跳转天气查询界面
	r.POST("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "weather.html", nil)
	})

	// 发送给邮箱界面
	r.GET("/E-mail.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "E-mail.html", gin.H{})
	})

	// 注册界面
	r.GET("/register.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	// 主界面
	r.GET("/", func(c *gin.Context) {
		c.File("../website/index.html")
	})

	r.GET("/getWeather", func(c *gin.Context) {
		url := "https://restapi.amap.com/v3/weather/weatherInfo?city=360783&key=049a159604b07b805722006015eedf50&extensions=all&output=JSON"
		response, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to getWeather data",
			})
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))

		//解析JSON数据到结构体
		var weatherResponse WeatherResponse
		err = json.Unmarshal(body, &weatherResponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to parse JSON data",
			})
		}
		c.JSON(http.StatusOK, weatherResponse)
	})

	// 启动Web服务
	r.Run(":8080")
}
