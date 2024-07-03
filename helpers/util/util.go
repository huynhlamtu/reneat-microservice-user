package util

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"reneat-microservice-user/config"
	"reneat-microservice-user/helpers/crypt"
	"reneat-microservice-user/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateUUID() (s string) {
	uuidNew, _ := uuid.NewUUID()
	return uuidNew.String()
}

func ShouldBindHeader(c *gin.Context) bool {
	platform := c.Request.Header.Get("X-PLATFORM")
	lang := c.Request.Header.Get("X-LANG")

	if platform == "" || lang == "" {
		return false
	}

	return true
}

func GetNowUTC() time.Time {
	loc, _ := time.LoadLocation("UTC")
	currentTime := time.Now().In(loc)
	return currentTime
}
func DebugJson(value interface{}) {
	fmt.Println(reflect.TypeOf(value).String())
	prettyJSON, _ := json.MarshalIndent(value, "", "    ")
	fmt.Printf("%s\n", string(prettyJSON))
}

func GetKeyFromContext(ctx context.Context, key string) (interface{}, bool) {
	if v := ctx.Value(key); v != nil {
		return v, true
	}

	return nil, false
}

func LogPrint(jsonData interface{}) {
	prettyJSON, _ := json.MarshalIndent(jsonData, "", "")
	fmt.Printf("%s\n", strings.ReplaceAll(string(prettyJSON), "\n", ""))
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func GetUserByToken(token string) (*models.User, error) {
	cfg := config.GetConfig()
	jwtSecret := cfg.GetString("jwt.secret")
	payload, err := crypt.ParseToken(token, jwtSecret)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ClientUuid: payload["client_uuid"].(string),
		Uuid:       payload["uuid"].(string),
		FirstName:  payload["first_name"].(string),
		LastName:   payload["last_name"].(string),
		Username:   payload["username"].(string),
		Email:      payload["email"].(string),
	}

	return &user, nil
}
