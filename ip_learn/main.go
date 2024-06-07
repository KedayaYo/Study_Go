package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/utils"
	"net/http"
)

func getIPWithValidatedProxyHeaders(c *gin.Context) {
	r := c.Request
	ip := utils.GetClientIP(r)
	log.Println("登录IP：" + ip)
	//ip := c.ClientIP()
	//ip = "8.8.8.8"
	//ip = "104.192.92.160"
	//ip2region(ip)
	geoIPInfo, err := utils.GeoIP(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ip": ip})
	}
	// 如果 IP 格式合法，则使用获取到的 IP，否则使用默认的 ClientIP 方法获取
	c.JSON(http.StatusOK, gin.H{"ip": ip, "geoIPInfo": geoIPInfo})
}

func main() {
	router := gin.Default()
	router.GET("/get-ip", getIPWithValidatedProxyHeaders)
	fmt.Println("Server is listening on port 8020...")
	router.Run(":8020")
}
