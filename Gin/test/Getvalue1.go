package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"log"

	"io"
	"net/http"
)

type Cities struct {
	CityCode string
	CityName string
}

func main() {
	r := gin.Default()
	r.Static("/static", "../web/static")
	r.LoadHTMLGlob("../website/*")

	// 登录跳转天气查询界面
	r.POST("/login", func(c *gin.Context) {
		excelFileName := "../docs/AMap_adcode_citycode.xlsx"
		xlFile, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		//假设EXCEL文件只有一个sheet，且你想要读取A列的数据
		sheet := xlFile.Sheets[0]

		var cities []Cities
		//遍历sheet的每一行，读取A列的数据
		for index, row := range sheet.Rows {
			if index < 1 {
				continue
			}
			if len(row.Cells) > 0 {
				cell := row.Cells[0]
				cell1 := row.Cells[1]
				cities = append(cities, Cities{
					CityCode: cell1.String(),
					CityName: cell.String(),
				})
			}
		}

		c.HTML(http.StatusOK, "weather0.html", gin.H{
			"cities": cities,
		})
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

	})
	r.GET("/getCityCode", func(c *gin.Context) {
		city := c.DefaultQuery("city", "")
		xlFile, err := xlsx.OpenFile("../docs/AMap_adcode_citycode.xlsx")
		if err != nil {
			log.Println("打开xlsx文件错误", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"错误": "打开不了Excel文件",
			})
			return
		}
		sheet := xlFile.Sheets[0]
		var cityCode string
		for index, row := range sheet.Rows {
			if index == 1 {
				continue
			}
			if len(row.Cells) >= 2 {
				if row.Cells[0].String() == city {
					cityCode = row.Cells[1].String()
					break
				}
			}
		}
		if cityCode != "" {
			c.JSON(http.StatusOK, gin.H{
				"cityCode": cityCode,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "city not found",
			})
		}
	})

	// 启动Web服务
	r.Run(":8080")
}
