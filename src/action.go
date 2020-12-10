package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
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

var e *echo.Echo

func initWebFramework() {
	e = echo.New()
	e.HideBanner = true

	e.POST("/api/login", login)
	e.POST("/api/register", register)
	e.GET("/api/member", getMember)
	e.POST("/api/member", setMember)
	e.DELETE("/api/member", deleteMember)

	e.GET("/urlapi/login", loginUrl)
	e.GET("/urlapi/register", registerUrl)
	e.GET("/urlapi/member", getMemberUrl)
	e.GET("/urlapi/member/create", setMemberUrl)
	e.GET("/urlapi/member/delete", deleteMemberUrl)

}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := []byte(c.FormValue("password"))
	if err := checkLoginInfo(&User{username, password}); err != nil  {
		return c.JSON(http.StatusUnauthorized, &LoginResponse{err.Error(), ""})
	} else {
		tokenString, _ := generateToken()
		return c.JSON(http.StatusOK, &LoginResponse{"", tokenString})
	}
}

func loginUrl(c echo.Context) error {
	username := c.QueryParam("username")
	password := []byte(c.QueryParam("password"))
	if err := checkLoginInfo(&User{username, password}); err != nil  {
		return c.JSON(http.StatusUnauthorized, &PlainErr{err.Error()})
	} else {
		tokenString, expireTime := generateToken()
		setCookie(&c, "login", tokenString, &expireTime)
		return c.JSON(http.StatusOK, nil)
	}
}

func register(c echo.Context) error {
	username := c.FormValue("username")
	password := []byte(c.FormValue("password"))
	if err := addUser(&User{username, password}); err != nil {
		return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func registerUrl(c echo.Context) error  {
	username := c.QueryParam("username")
	password := []byte(c.QueryParam("password"))
	if err := addUser(&User{username, password}); err != nil {
		return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

func getMember(c echo.Context) error {
	if _, err := parseToken(c.FormValue("token")); err == nil {
		var qscers []Qscer
		db.Find(&qscers)
		return c.JSON(http.StatusOK, &qscers)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"Unauthorized"})
	}
}

func getMemberUrl(c echo.Context) error {
	if authentication(&c, "login") {
		var qscers []Qscer
		db.Find(&qscers)
		return c.JSON(http.StatusOK, &qscers)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func setMember(c echo.Context) error {
	if _, err := parseToken(c.FormValue("token")); err == nil {
		zjuid := c.FormValue("zjuid")
		name := c.FormValue("name")
		qscid := c.FormValue("qscid")
		birthday := c.FormValue("birthday")
		if err := addQSCer(&Qscer{0, zjuid, name, qscid, birthday}); err != nil {
			return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func setMemberUrl(c echo.Context) error {
	if authentication(&c, "login") {
		zjuid := c.QueryParam("zjuid")
		name := c.QueryParam("name")
		qscid := c.QueryParam("qscid")
		birthday := c.QueryParam("birthday")
		if err := addQSCer(&Qscer{0, zjuid, name, qscid, birthday}); err != nil {
			return c.JSON(http.StatusBadRequest, &PlainErr{err.Error()})
		} else {
			return c.JSON(http.StatusOK, nil)
		}
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func deleteMember(c echo.Context) error {
	if _, err := parseToken(c.FormValue("token")); err == nil {
		uid := c.FormValue("uid")
		db.Delete(&Qscer{}, uid)
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func deleteMemberUrl(c echo.Context) error {
	if authentication(&c, "login") {
		uid := c.QueryParam("uid")
		if err := db.First(&Qscer{}, uid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &PlainErr{"user not found"})
		}
		result := db.Delete(&Qscer{}, uid)
		fmt.Println(result)
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusUnauthorized, &PlainErr{"unauthorized"})
	}
}

func startServer() {
	e.Logger.Fatal(e.Start(":8080"))
}
