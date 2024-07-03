package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"reneat-microservice-user/config"
	"reneat-microservice-user/constant"
	"reneat-microservice-user/helpers/crypt"
	"reneat-microservice-user/helpers/respond"
	timeHelper "reneat-microservice-user/helpers/time"
	"reneat-microservice-user/helpers/util"
	"reneat-microservice-user/models"
	request "reneat-microservice-user/request/auth"
	"strings"
)

var userEntity = "User"

type AuthController struct {
	UserModel models.User
}

type JWTClaim struct {
	JWTPayload
	jwt.StandardClaims
}

type JWTPayload struct {
	ClientUuid string `json:"client_uuid"`
	Uuid       string `json:"uuid"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	IsAdmin    int    `json:"is_admin"`
}

// Register
//
//	@Summary	Register user
//	@Schemes
//	@Description	Register new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.GetLoginRequest	true	"query body"
//	@Success		200		{object}	respond.Respond
//	@Failure		400		{object}	respond.Respond
//	@Router			/users/register [post]
func (authCtl AuthController) Register(c *gin.Context) {
	var req request.RegisterRequest

	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingRegister())
		return
	}

	cond := bson.M{
		"$or": []bson.M{
			{"email": strings.ToLower(strings.Trim(req.Email, " "))},
			{"username": strings.ToLower(strings.Trim(req.Username, " "))},
		},
	}

	existedUser, err := new(models.User).FindOne(cond)

	if existedUser != nil {
		logrus.Error(err)
		if existedUser.Username == strings.ToLower(strings.Trim(req.Username, " ")) {
			c.JSON(http.StatusUnprocessableEntity, respond.FieldAlreadyExist("Username"))
		} else {
			c.JSON(http.StatusUnprocessableEntity, respond.FieldAlreadyExist("Email"))
		}
		return
	}

	pwd, _ := crypt.HashPassword(req.Password)

	user := models.User{}
	user.Uuid = util.GenerateUUID()
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = strings.ToLower(strings.Trim(req.Email, " "))
	user.Username = strings.ToLower(strings.Trim(req.Username, " "))
	user.Password = pwd
	user.IsActive = constant.ACTIVE
	user.IsVerified = constant.INACTIVE
	user.IsBlock = constant.INACTIVE
	user.CreatedAt = timeHelper.NowUTC()
	user.UpdatedAt = timeHelper.NowUTC()

	_, err = user.Insert()

	if err != nil {
		fmt.Println("Insert User Fail -", user.Email, err)
		logrus.Error(err)
		c.JSON(http.StatusUnprocessableEntity, respond.CanNotCreate(userEntity))
		return
	}

	c.JSON(http.StatusOK, respond.Success(bson.M{"uuid": user.Uuid}, "Created Successfully"))
}

// Login
//
//	@Summary	Login user
//	@Schemes
//	@Description	Login with user info then return the token credential
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.GetLoginRequest	true	"query body"
//	@Success		200		{object}	respond.Respond
//	@Failure		400		{object}	respond.Respond
//	@Router			/users/login [post]
func (authCtl AuthController) Login(c *gin.Context) {
	var req request.GetLoginRequest
	cfg := config.GetConfig()

	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingLogin())
		return
	}

	cond := bson.M{
		"email": strings.Trim(req.Email, " "),
	}

	user, err := authCtl.UserModel.FindOne(cond)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, respond.EmailPasswordIncorrect())
		return
	}

	ok := crypt.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, respond.EmailPasswordIncorrect())
		return
	}

	jwtExpiry := cfg.GetDuration("jwt.expires_at")
	expiresAt := timeHelper.Now().Add(jwtExpiry)

	claim := JWTClaim{
		JWTPayload{
			ClientUuid: user.ClientUuid,
			Uuid:       user.Uuid,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			IsAdmin:    user.IsAdmin,
		},
		jwt.StandardClaims{
			Issuer:    "reneat-microservice-user",
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtSecret := cfg.GetString("jwt.secret")
	token, err := crypt.CreateToken(jwtSecret, claim)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, respond.CreatedFail())
		return
	}

	userToken := models.UserToken{
		UserUuid:  user.Uuid,
		Token:     token,
		ExpiredAt: expiresAt,
		IsActive:  constant.ACTIVE,
	}
	_, _ = userToken.Insert()

	resp := request.GetLoginResponse{
		ClientUuid: user.ClientUuid,
		Uuid:       user.Uuid,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Token:      token,
	}

	c.JSON(http.StatusOK, respond.Success(resp, "Login successfully!"))
}
