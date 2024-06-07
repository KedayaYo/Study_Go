package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"log"
	"main/constant"
	"net"
	"net/http"
	"strings"
	"time"
)

type GeoIPInfo struct {
	City            string  `json:"city"`
	Subdivision     string  `json:"subdivision"`
	Country         string  `json:"country"`
	ISOCode         string  `json:"iso_code"`
	TimeZone        string  `json:"time_zone"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	IsPublic        bool    `json:"is_public"`
	IP              string  `json:"ip"`
	Browser         string  `json:"browser"`
	OperatingSystem string  `json:"operatingSystem"`
}

// GeoIP 函数用于获取地理位置信息
func GeoIP(rc *gin.Context) (*GeoIPInfo, error) {
	r := rc.Request
	userAgent := rc.GetHeader("User-Agent")
	os, browser := GetOSAndBrowser(userAgent)
	ipAddress := GetClientIP(r)

	// 打开 GeoIP 数据库
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Println("无法打开 GeoIP 数据库:", err)
		return nil, err
	}
	defer db.Close()

	// 解析IP地址
	ip := net.ParseIP(ipAddress)
	isPub := IsPublicIP(ip)

	// 获取城市信息
	record, err := db.City(ip)
	if err != nil {
		log.Println("获取城市信息失败:", err)
		return nil, err
	}

	// 创建 GeoIPInfo 实例
	info := &GeoIPInfo{
		Country:         record.Country.Names["zh-CN"],
		City:            record.City.Names["zh-CN"],
		ISOCode:         record.Country.IsoCode,
		TimeZone:        record.Location.TimeZone,
		Latitude:        record.Location.Latitude,
		Longitude:       record.Location.Longitude,
		IsPublic:        isPub,
		IP:              ip.String(),
		Browser:         browser,
		OperatingSystem: os,
	}

	// 如果有分区信息，则添加分区信息
	if len(record.Subdivisions) > 0 {
		info.Subdivision = record.Subdivisions[0].Names["zh-CN"]
	}

	return info, nil
}

// GetClientIP 获取客户端的公网IP地址
func GetClientIP(r *http.Request) string {
	ipAddress := ""

	// 根据常见的代理服务器转发的请求ip存放协议，从请求头获取原始请求ip
	ipAddress = r.Header.Get(constant.HEADER_X_FORWARDED_FOR)
	if ipAddress == "" || strings.EqualFold(ipAddress, constant.UNKNOWN) {
		ipAddress = r.Header.Get(constant.HEADER_PROXY_CLIENT_IP)
	}
	if ipAddress == "" || strings.EqualFold(ipAddress, constant.UNKNOWN) {
		ipAddress = r.Header.Get(constant.HEADER_WL_PROXY_CLIENT_IP)
	}

	// 如果没有转发的ip，则取当前通信的请求端的ip
	if ipAddress == "" || strings.EqualFold(ipAddress, constant.UNKNOWN) {
		ipAddress, _, _ = net.SplitHostPort(r.RemoteAddr)
		if ipAddress == constant.LOCALHOST {
			interfaces, err := net.Interfaces()
			if err == nil {
				for _, iface := range interfaces {
					if addrs, err := iface.Addrs(); err == nil {
						for _, addr := range addrs {
							if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
								if ipNet.IP.To4() != nil {
									ipAddress = ipNet.IP.String()
									break
								}
							}
						}
					}
				}
			}
		}
	}

	// 对于通过多个代理的情况，第一个IP为客户端真实IP,多个IP按照','分割
	if ipAddress != "" {
		ipAddress = strings.TrimSpace(strings.Split(ipAddress, constant.SEPARATOR)[0])
	}

	return ipAddress
}

// IsPublicIP 判断一个 IP 是否为公网 IP
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}

	ip4 := IP.To4()
	if ip4 == nil {
		// 检查IPv6特殊地址
		return !isPrivateIPv6(IP)
	}

	// 检查IPv4私有地址
	switch {
	case ip4[0] == 10:
		return false
	case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
		return false
	case ip4[0] == 192 && ip4[1] == 168:
		return false
	case ip4[0] == 100 && (ip4[1] >= 64 && ip4[1] <= 127): // 100.64.0.0/10
		return false
	case ip4[0] == 127:
		return false
	case ip4[0] == 198 && (ip4[1] == 18 || ip4[1] == 19): // 198.18.0.0/15
		return false
	default:
		return true
	}
}

// isPrivateIPv6 检查 IPv6 地址是否为私有地址
func isPrivateIPv6(ip net.IP) bool {
	return ip.IsLoopback() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() ||
		(ip[0] == 0xfd) || // 唯一本地地址 (ULA)
		(ip[0] == 0xfe && ip[1] == 0xc0) // 站点本地地址 (已废弃，但仍然有效)
}

// GetOSAndBrowser 解析User-Agent头并获取操作系统和浏览器信息
func GetOSAndBrowser(userAgent string) (string, string) {
	ua := user_agent.New(userAgent)
	name, version := ua.Browser()
	os := ua.OS()

	return os, fmt.Sprintf("%s %s", name, version)
}

// geoIP 用于测试GeoIP数据库的功能
func geoIP(ipAddress string) {
	db, err := geoip2.Open("ip_learn/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(ipAddress)
	isPub := IsPublicIP(ip)
	if !isPub {
		log.Println("内网IP")
	}

	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

	fmt.Println("中文结果")
	fmt.Printf("中文城市名: %v\n", record.City.Names["zh-CN"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("中文分区名: %v\n", record.Subdivisions[0].Names["zh-CN"])
	}
	fmt.Printf("中文国家名: %v\n", record.Country.Names["zh-CN"])
	fmt.Printf("ISO 国家代码: %v\n", record.Country.IsoCode)
	fmt.Printf("时区: %v\n", record.Location.TimeZone)
	fmt.Printf("坐标: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}

// ip2region 用于测试IP2Region数据库的功能
func ip2region(ip string) {
	var dbPath = "ip_learn/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("创建 searcher 失败: %s\n", err.Error())
		return
	}
	defer searcher.Close()

	var tStart = time.Now()
	ips, err := net.LookupIP("www.github.com")
	if err != nil {
		fmt.Printf("域名解析失败: %s\n", err.Error())
		return
	}
	ipres := ips[0].String()
	fmt.Printf("域名的IP：%s\n", ipres)

	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("IP搜索失败(%s): %s\n", ip, err)
		return
	}
	fmt.Printf("{region: %s, 耗时: %s}\n\n", region, time.Since(tStart))
	// 备注：并发使用时，每个 goroutine 需要创建一个独立的 searcher 对象。
}
