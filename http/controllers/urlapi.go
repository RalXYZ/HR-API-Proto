package controllers

import (
	"HR-API-proto/database"
	"HR-API-proto/database/models"
	"HR-API-proto/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginUrl(c echo.Context) error {
	username := c.QueryParam("username")
	password := []byte(c.QueryParam("password"))
	if err := database.VerifyLoginCredential(&models.User{username, password}); err != nil {
		return c.JSON(http.StatusUnauthorized, &PlainErr{err.Error()})
	} else {
		tokenString, expireTime := utils.GenerateJWT()
		utils.SetCookie(&c, "login", tokenString, &expireTime)
		return c.JSON(http.StatusOK, nil)
	}
}

func RegisterUrl(c echo.Context) error {
	username := c.QueryParam("username")
	password := []byte(c.QueryParam("password"))
	if err := database.AddUser(&models.User{username, password}); err != nil {
		return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func GetMemberUrl(c echo.Context) error {
	if utils.Authentication(&c, "login") {
		var qscers []models.Qscer
		database.DB.Find(&qscers)
		return c.JSON(http.StatusOK, &qscers)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func SetMemberUrl(c echo.Context) error {
	if utils.Authentication(&c, "login") {
		zjuid := c.QueryParam("zjuid")
		name := c.QueryParam("name")
		qscid := c.QueryParam("qscid")
		birthday := c.QueryParam("birthday")
		if err := database.AddQSCer(&models.Qscer{0, zjuid, name, qscid, birthday}); err != nil {
			return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func DeleteMemberUrl(c echo.Context) error {
	if utils.Authentication(&c, "login") {
		uid := c.QueryParam("uid")
		if err := database.DB.First(&models.Qscer{}, uid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &PlainErr{"user not found"})
		}
		result := database.DB.Delete(&models.Qscer{}, uid)
		fmt.Println(result)
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}
