package with

import (
	"fmt"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

const DatabaseContextKey = "db"

func Database(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("with inside")
		return next(c)
	}
}

func GetDatabase(c echo.Context) *mgo.Database {
	if db, ok := c.Get(DatabaseContextKey).(*mgo.Database); ok {
		return db
	}

	return nil
}
