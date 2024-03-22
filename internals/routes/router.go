package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thesamueliyeh/cbt-app-v1/internals/handlers"
	"github.com/thesamueliyeh/cbt-app-v1/internals/middlewares"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", handlers.IndexPageHandler, middlewares.LoggedOutUserMiddleware)
	e.GET("/auth", handlers.AuthPageHandler, middlewares.LoggedOutUserMiddleware)
	e.GET("/dashboard", handlers.DashboardPageHandler, middlewares.LoggedinUserMiddleware)
	e.GET("/profile", handlers.ProfilePageHandler, middlewares.LoggedinUserMiddleware)
	e.GET("/create-exam", handlers.CreateExamPageHandler, middlewares.LoggedinUserMiddleware)
	e.GET("/take-exam", handlers.TakeExamPageHandler)
	e.GET("/exam", handlers.ExamPageHandler)
	e.POST("/signup", handlers.SignupApiHandler, middlewares.LoggedOutUserMiddleware)
	e.POST("/login", handlers.LoginApiHandler, middlewares.LoggedOutUserMiddleware)
	e.POST("/forgotpassword", handlers.ForgotPasswordApiHandler, middlewares.LoggedOutUserMiddleware)
	e.POST("/logout", handlers.LogoutApiHandler, middlewares.LoggedinUserMiddleware)
	e.GET("/verify", handlers.VerifyOtpHandler, middlewares.LoggedOutUserMiddleware)
	e.GET("/updatepassword", handlers.UpdatePasswordPageHandler, middlewares.LoggedinUserMiddleware)
	e.POST("/updatepassword", handlers.UpdatePasswordApiHandler, middlewares.LoggedinUserMiddleware)
	// e.GET("/login", handlers.LoginFormHandler)
	// e.GET("/signup", handlers.SignupFormHandler)
}
