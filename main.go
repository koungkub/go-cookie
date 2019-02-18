package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://web.smartfishermans.com"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	e.GET("/cookie", func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "thisiscookie"
		cookie.Value = "thisiscookievalue"
		cookie.MaxAge = 30
		// cookie.Expires = time.Now().Add(time.Hour)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, echo.Map{
			"cookie": true,
		})
	})

	e.GET("/see", func(c echo.Context) error {
		cookie, err := c.Cookie("thisiscookie")
		if err != nil {
			return echo.NewHTTPError(422, "sdf")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"name":  cookie.Name,
			"value": cookie.Value,
		})
	})

	e.GET("/delete", func(c echo.Context) error {
		cookie, err := c.Cookie("thisiscookie")
		if err != nil {
			return echo.NewHTTPError(422, "sf")
		}
		cookie.MaxAge = -1
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, echo.Map{
			"delete": true,
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
