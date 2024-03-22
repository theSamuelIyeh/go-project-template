package services

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

type SignupData struct {
	Phone       string `json:"phone,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

var Supabase *supa.Client

func InitSupabase() {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	Supabase = supa.CreateClient(supabaseUrl, supabaseKey)
	fmt.Println("Supabase auth connected successfully")
}

// login auth service
func LoginService(c echo.Context, email, password string) (*supa.AuthenticatedDetails, error) {
	ctx := c.Request().Context()
	user, err := Supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	return user, err
}

// signup auth service
func SignupService(c echo.Context, email, password, displayName, phone string) (*supa.User, error) {
	ctx := c.Request().Context()
	user, err := Supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
		Data: SignupData{
			DisplayName: displayName,
			Phone:       phone,
		},
	})
	return user, err
}

// forgot password auth service
func ForgotPasswordService(c echo.Context, email string) error {
	ctx := c.Request().Context()
	err := Supabase.Auth.ResetPasswordForEmail(ctx, email)
	return err
}

// logout user auth service
func LogoutService(c echo.Context, accessToken string) error {
	ctx := c.Request().Context()
	err := Supabase.Auth.SignOut(ctx, accessToken)
	return err
}

// verify otp auth service
func VerifyOtpService(c echo.Context, tokenHash, verifyType string) (*supa.AuthenticatedDetails, error) {
	ctx := c.Request().Context()
	user, err := Supabase.Auth.VerifyOtp(ctx, supa.VerifyOtpCredentials{
		TokenHash:  tokenHash,
		VerifyType: verifyType,
	})
	return user, err
}

// update user data
func UpdateUserDataService(c echo.Context, accessToken string, details map[string]interface{}) (*supa.User, error) {
	ctx := c.Request().Context()
	user, err := Supabase.Auth.UpdateUser(ctx, accessToken, details)
	return user, err
}
