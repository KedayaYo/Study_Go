package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	clientID     = "Ov23liC0XVKwWK8d9odN"
	clientSecret = "02ae7329fefa21ccde73c6e557ec3b1b97a2001d"
	db           *sql.DB
	rdb          *redis.Client
	ctx          = context.Background()
	jwtSecretKey = []byte("skysys")
)

var (
	githubOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8020/login/oauth2/code/github",
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}
	oauthStateString  = generateStateOauthCookie()
	ErrorInvalidToken = errors.New("verify Token Failed")
)

// 初始化数据库和 Redis 连接
func init() {
	var err error
	// 配置数据库连接字符串
	dsn := "root:root@tcp(118.25.27.160:33306)/custom_project_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("无法ping数据库: %v", err)
	}

	// 初始化 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
		DB:   0,                // 使用默认的 Redis 数据库
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}
}

// 生成随机的状态字符串以防止CSRF攻击
func generateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// 处理GitHub登录请求
func handleGitHubLogin(c *gin.Context) {
	// 生成GitHub OAuth2授权URL并重定向
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	log.Println("url:", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// 处理GitHub回调请求
func handleGitHubCallback(c *gin.Context) {
	// 验证state参数以防止CSRF攻击
	if c.Query("state") != oauthStateString {
		log.Println("无效的OAuth状态")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")

	// 通过授权码获取token
	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("oauthConf.Exchange() 失败: %v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	client := githubOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Printf("client.Get() 失败: %v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer resp.Body.Close()

	var user struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Printf("json.Decode() 失败: %v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 创建JWT
	accessToken, refreshToken, err := createJWT(user.ID, user.Login)
	if err != nil {
		log.Printf("createJWT() 失败: %v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 保存第三方登录信息到数据库
	err = saveThirdPartyLoginInfo(user.ID, "github", token.AccessToken, token.RefreshToken)
	if err != nil {
		log.Printf("saveThirdPartyLoginInfo() 失败: %v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 将令牌返回给客户端
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// 创建JWT
func createJWT(userID int, username string) (string, string, error) {
	// 创建访问令牌
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	// 创建刷新令牌
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	// 将令牌存储到 Redis
	err = storeTokens(userID, accessTokenString, refreshTokenString)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// 将令牌存储到 Redis
func storeTokens(userID int, accessToken, refreshToken string) error {
	atKey := fmt.Sprintf("access_token_%d", userID)
	rtKey := fmt.Sprintf("refresh_token_%d", userID)

	err := rdb.Set(ctx, atKey, accessToken, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, rtKey, refreshToken, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

// 验证Token
func VerifyToken(tokenID string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenID, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrorInvalidToken
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, ErrorInvalidToken
	}

	return claims, nil
}

// 通过 refresh token 刷新 access token
func RefreshToken(c *gin.Context) {
	var request struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式错误"})
		return
	}

	userID, err := verifyRefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的刷新令牌"})
		return
	}

	newAccessToken, newRefreshToken, err := createJWT(userID, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "令牌刷新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

// 验证 refresh token
func verifyRefreshToken(rtoken string) (int, error) {
	var claim jwt.MapClaims
	token, err := jwt.ParseWithClaims(rtoken, &claim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return 0, ErrorInvalidToken
	}

	userID := int(claim["user_id"].(float64))

	// 检查 redis 中是否存在该 refresh token
	rtKey := fmt.Sprintf("refresh_token_%d", userID)
	storedRToken, err := rdb.Get(ctx, rtKey).Result()
	if err != nil || storedRToken != rtoken {
		return 0, ErrorInvalidToken
	}

	return userID, nil
}

// 保存第三方登录信息到数据库
func saveThirdPartyLoginInfo(accountID int, platform, accessToken, refreshToken string) error {
	query := `INSERT INTO tb_sys_third_party (account_id, platform, open_id, access_token, refresh_token, status, creator, creator_id, create_time, modifier, modifier_id, update_time)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE access_token=?, refresh_token=?, update_time=?`
	_, err := db.Exec(query, accountID, platform, fmt.Sprintf("%d", accountID), accessToken, refreshToken, 1, "system", 0, time.Now(), "system", 0, time.Now(), accessToken, refreshToken, time.Now())
	return err
}

func main() {
	r := gin.Default()
	r.GET("/login", handleGitHubLogin)                       // GitHub登录处理
	r.GET("/login/oauth2/code/github", handleGitHubCallback) // GitHub回调处理

	r.POST("/refresh_token", RefreshToken) // 刷新令牌

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	log.Println("服务已启动在 http://localhost:8020")
	if err := r.Run(":8020"); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
