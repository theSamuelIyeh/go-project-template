package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/thesamueliyeh/cbt-app-v1/views"
)

func IndexPageHandler(c echo.Context) error {
	indexPage := views.IndexPage()
	return indexPage.Render(c.Request().Context(), c.Response().Writer)
}

func AuthPageHandler(c echo.Context) error {
	formType := c.QueryParam("type")
	authPage := views.AuthPage(formType)
	return authPage.Render(c.Request().Context(), c.Response().Writer)
}

func DashboardPageHandler(c echo.Context) error {
	dashboardPage := views.DashboardPage()
	return dashboardPage.Render(c.Request().Context(), c.Response().Writer)
}

func ProfilePageHandler(c echo.Context) error {
	profilePage := views.ProfilePage()
	return profilePage.Render(c.Request().Context(), c.Response().Writer)
}

func CreateExamPageHandler(c echo.Context) error {
	createExamPage := views.CreateExamPage()
	return createExamPage.Render(c.Request().Context(), c.Response().Writer)
}

func TakeExamPageHandler(c echo.Context) error {
	takeExamPage := views.TakeExamPage()
	return takeExamPage.Render(c.Request().Context(), c.Response().Writer)
}

func ExamPageHandler(c echo.Context) error {
	examPage := views.ExamPage("test exam")
	return examPage.Render(c.Request().Context(), c.Response().Writer)
}

func UpdatePasswordPageHandler(c echo.Context) error {
	updatePasswordPage := views.UpdatePasswordPage()
	return updatePasswordPage.Render(c.Request().Context(), c.Response().Writer)
}
