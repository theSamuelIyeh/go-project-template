package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thesamueliyeh/cbt-app-v1/internals/services"
	"github.com/thesamueliyeh/cbt-app-v1/internals/utils"
	"github.com/thesamueliyeh/cbt-app-v1/views"
)

// login form struct
type LoginDetails struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

// signup form struct
type SignupDetails struct {
	Email     string `form:"email" validate:"required,email"`
	Password  string `form:"password" validate:"required,min=6"`
	FirstName string `form:"first-name" validate:"required"`
	LastName  string `form:"last-name" validate:"required"`
	Phone     string `form:"phone" validate:"required"`
}

// forgot password form struct
type ForgotPasswordDetails struct {
	Email string `form:"email" validate:"required,email"`
}

// verify otp query struct
type VerifyOtpDetails struct {
	TokenHash  string `query:"token_hash" validate:"required"`
	VerifyType string `query:"type" validate:"required"`
}

// update new password struct
type UpdatePasswordDetails struct {
	NewPassword        string `form:"password1" validate:"required"`
	ConfirmNewPassword string `form:"password2" validate:"required"`
}

// login api handler
func LoginApiHandler(c echo.Context) error {
	// create new user struct
	u := new(LoginDetails)

	// bind incoming data to struct
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}

	// validate incoming data
	if err := c.Validate(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}

	// call login service
	user, err := services.LoginService(c, u.Email, u.Password)
	if err != nil {
		errMsg := strings.Split(err.Error(), ":")
		errMsgPart := errMsg[1]
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusInternalServerError, "<p x-show=\"showError\">"+errMsgPart+"</p>")
	}

	// set cookies
	utils.CreateCookie(c, user.AccessToken, "access_token", user.ExpiresIn)
	utils.CreateCookie(c, user.RefreshToken, "refresh_token", 0)
	return utils.RedirectToUrl(c, "/dashboard")
}

// signup api handler
func SignupApiHandler(c echo.Context) error {
	ctx := c.Request().Context()
	u := new(SignupDetails)
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	displayName := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	_, err := services.SignupService(c, u.Email, u.Password, displayName, u.Phone)
	if err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusInternalServerError, "<p x-show=\"showError\">"+err.Error()+"</p>")
	}
	confirmEmailComponent := views.ConfirmEmailComponent()
	return confirmEmailComponent.Render(ctx, c.Response().Writer)
}

// forgot password handler
func ForgotPasswordApiHandler(c echo.Context) error {
	ctx := c.Request().Context()
	u := new(ForgotPasswordDetails)
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	err := services.ForgotPasswordService(c, u.Email)
	if err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		c.Response().Header().Set("hx-trigger", "swap")
		return c.HTML(http.StatusInternalServerError, "<p x-show=\"showError\">"+err.Error()+"</p>")
	}
	successfulPasswordResetComponent := views.SuccessfulPasswordResetComponent()
	return successfulPasswordResetComponent.Render(ctx, c.Response().Writer)
}

// logout handler
func LogoutApiHandler(c echo.Context) error {
	accessToken := c.Get("access_token").(string)
	err := services.LogoutService(c, accessToken)
	if err == nil {
		utils.CreateCookie(c, "Bearer fake_token", "access_token", -1)
		utils.CreateCookie(c, "Bearer fake_token", "refresh_token", -1)
	}
	return utils.RedirectToUrl(c, "/auth?type=login")
}

// verify otp handler
func VerifyOtpHandler(c echo.Context) error {
	u := new(VerifyOtpDetails)
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	user, err := services.VerifyOtpService(c, u.TokenHash, u.VerifyType)
	if err != nil {
		return utils.RedirectToUrl(c, "/auth?type=login")
	}
	accessToken := user.AccessToken
	refreshToken := user.RefreshToken
	expiresIn := user.ExpiresIn
	utils.CreateCookie(c, accessToken, "access_token", expiresIn)
	utils.CreateCookie(c, refreshToken, "refresh_token", 0)
	if u.VerifyType == "recovery" {
		utils.CreateCookie(c, "true", "password_change", 0)
	}
	return utils.RedirectToUrl(c, "/dashboard")
}

// update new password
func UpdatePasswordApiHandler(c echo.Context) error {
	u := new(UpdatePasswordDetails)
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(u); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if u.NewPassword != u.ConfirmNewPassword {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Passwords do not match</p>")
	}
	accessToken := c.Get("access_token").(string)
	details := map[string]interface{}{
		"password": u.ConfirmNewPassword,
	}
	_, err := services.UpdateUserDataService(c, accessToken, details)
	if err != nil {
		// fmt.Println(err)
		return utils.RedirectToUrl(c, "/auth?type=login")
	}
	utils.CreateCookie(c, "false", "password_change", -1)
	return utils.RedirectToUrl(c, "/dashboard")
}
