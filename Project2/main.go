package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Result struct{
	Adition int
	Subtraction int
	Multiplication int
	Division int
}

func Calculation(c echo.Context)error{
	num1:= c.QueryParam("num1")
	a, err := strconv.Atoi(num1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	num2:= c.QueryParam("num2")
	b, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	result := Result{
		Adition: a + b,
		Subtraction: a - b,
		Multiplication: a * b,
		Division: a / b,
	}
	return c.JSON(http.StatusOK, result)
}

func main() {
	/* New echo instance */
	e := echo.New()
	/* Route with GET method */
	e.GET("/calculate", Calculation)

	/* Start server */
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("Server started at port 8080")

}