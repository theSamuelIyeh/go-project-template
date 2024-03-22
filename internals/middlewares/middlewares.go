package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/thesamueliyeh/cbt-app-v1/internals/services"
	"github.com/thesamueliyeh/cbt-app-v1/internals/utils"
)

func LoggedinUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check if user is logged in
		// if not redirect to login page
		ctx := c.Request().Context()
		jwtSecret := os.Getenv("JWT_SECRET")

		// get access and refresh token from cookies
		accessCookie, accessToken, accessErr := utils.GetTokenFromCookie(c, "access_token")
		_, refreshToken, refreshErr := utils.GetTokenFromCookie(c, "refresh_token")

		if accessErr != nil || refreshErr != nil {
			utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
			utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
			utils.CreateCookie(c, "Bearer Fake token", "password_change", -1)
			return utils.RedirectToUrl(c, "/auth?type=login")
		}

		// parse and validate jwt
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			return utils.RedirectToUrl(c, "/auth?type=login")
		}

		// check if token is valid
		if token.Valid {
			// check access cookie expiry for refresh
			if accessCookie.Expires.Before(time.Now().Add(5 * time.Minute)) {
				user, err := services.Supabase.Auth.RefreshUser(ctx, accessToken, refreshToken)
				if err != nil {
					return utils.RedirectToUrl(c, "/auth?type=login")
				}

				// Refresh the cookie with new access token
				newAccessToken := user.AccessToken
				newRefreshToken := user.RefreshToken
				utils.CreateCookie(c, newAccessToken, "access_token", user.ExpiresIn)
				utils.CreateCookie(c, newRefreshToken, "refresh_token", 0)

				// set access token and user to context
				c.Set("access_token", newAccessToken)
				c.Set("user", user.User)
			} else {
				// get the user using access token
				user, err := services.Supabase.Auth.User(ctx, accessToken)
				if err != nil {
					return utils.RedirectToUrl(c, "/auth?type=login")
				}

				// set access token and user to context
				c.Set("access_token", accessToken)
				c.Set("user", user)
			}
			_, _, passwordChangeErr := utils.GetTokenFromCookie(c, "password_change")
			if passwordChangeErr == nil {
				if c.Path() == "/updatepassword" {
					return next(c)
				} else {
					return utils.RedirectToUrl(c, "/updatepassword")
				}
			} else {
				if c.Path() == "/updatepassword" {
					return utils.RedirectToUrl(c, "/dashboard")
				} else {
					return next(c)
				}
			}
		}
		return utils.RedirectToUrl(c, "/auth?type=login")
	}
}

// middleware for unprotected routes
func LoggedOutUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("access_token")
		if err != nil {
			utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
			utils.CreateCookie(c, "Bearer Fake token", "password_change", -1)
			return next(c)
		}
		return utils.RedirectToUrl(c, "/dashboard")
	}
}
