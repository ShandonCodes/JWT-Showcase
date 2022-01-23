package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

//Custom claims object
type UserTokenClaims struct {
	ID                    string   `json:"id"`
	Roles                 []string `json:"roles"`
	PasswordResetRequired bool     `json:"passwordResetRequired"`
	jwt.StandardClaims
	Name string `json:"name"`
}

func main() {

	e := echo.New()

	//Login and generate token
	e.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		if email != "test@test.io" || password != "password" {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "User does not exists or is disabled"})
		}

		claims := UserTokenClaims{
			uuid.NewString(),
			[]string{"ADMIN"},
			false,
			jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour).Unix()},
			fmt.Sprintf("%s %s", "Test", "Testing"),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, _ := token.SignedString([]byte("SECRET"))

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	e.Logger.Fatal(e.Start(":7877"))
}
