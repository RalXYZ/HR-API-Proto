package controllers

import (
	"HR-API-proto/database"
	"HR-API-proto/database/models"
	"HR-API-proto/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

type PlainErr struct {
	Error string `json:"error"`
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := []byte(c.FormValue("password"))
	if err := database.VerifyLoginCredential(&models.User{username, password}); err != nil {
		return c.JSON(http.StatusUnauthorized, &LoginResponse{err.Error(), ""})
	} else {
		tokenString, _ := utils.GenerateJWT()
		return c.JSON(http.StatusOK, &LoginResponse{"", tokenString})
	}
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := []byte(c.FormValue("password"))
	if err := database.AddUser(&models.User{username, password}); err != nil {
		return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func GetMember(c echo.Context) error {
	if _, err := utils.ParseToken(c.FormValue("token")); err == nil {
		var qscers []models.Qscer
		database.DB.Find(&qscers)
		return c.JSON(http.StatusOK, &qscers)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"Unauthorized"})
	}
}

func SetMember(c echo.Context) error {
	if _, err := utils.ParseToken(c.FormValue("token")); err == nil {
		zjuid := c.FormValue("zjuid")
		name := c.FormValue("name")
		qscid := c.FormValue("qscid")
		birthday := c.FormValue("birthday")
		if err := database.AddQSCer(&models.Qscer{0, zjuid, name, qscid, birthday}); err != nil {
			return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func DeleteMember(c echo.Context) error {
	if _, err := utils.ParseToken(c.FormValue("token")); err == nil {
		uid := c.FormValue("uid")
		database.DB.Delete(&models.Qscer{}, uid)
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

