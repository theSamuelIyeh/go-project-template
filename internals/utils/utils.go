package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func CreateCookie(c echo.Context, token, name string, expires int) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprintf("%s %s", "Bearer", token)
	cookie.HttpOnly = true
	if expires != 0 {
		cookie.Expires = time.Now().Add(time.Duration(expires) * time.Second)
	}
	if expires == -1 {
		cookie.Expires = time.Unix(0, 0)
	}
	c.SetCookie(cookie)
}

func IsHXRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func RedirectToUrl(c echo.Context, url string) error {
	if IsHXRequest(c) {
		c.Response().Header().Set("Hx-redirect", url)
		c.Response().Header().Set("hx-trigger", "swap")
		return c.String(http.StatusPermanentRedirect, "")
	}
	return c.Redirect(http.StatusPermanentRedirect, url)
}

func GetTokenFromCookie(c echo.Context, tokenName string) (*http.Cookie, string, error) {
	accessCookie, err := c.Cookie(tokenName)
	if err != nil {
		return nil, "", err
	}
	accessTokenParts := strings.Split(accessCookie.Value, " ")
	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		return nil, "", errors.New("invalid access token format")
	}
	return accessCookie, accessTokenParts[1], nil
}
