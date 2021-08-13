package controllers

import (
	"dao"
	"dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var route *gin.Engine

func StartServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	route = gin.Default()
	route.Use(cors())

	customerRoute := route.Group("/customers")
	{
		customerRoute.GET("/findAll", findAllCustomers)
		customerRoute.GET("/findById/:id", findCustomerById)
		customerRoute.POST("/add", addCustomer)
		customerRoute.PUT("/update/:id", updateCustomer)
		customerRoute.DELETE("/delete/:id", deleteCustomer)
	}
	route.Run(fmt.Sprintf(":%d", port))
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func findAllCustomers(c *gin.Context) {
	customers, err := dao.FindAllCustomers()
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"customers": customers,
		})
	}
}

func findCustomerById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	customer, err := dao.FindCustomerById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"customer": customer,
		})
	}
}

func addCustomer(c *gin.Context) {
	var customer dto.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := dao.AddCustomer(customer)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}
	if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is created",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not insert customer in database",
		})
	}
}

func updateCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var customer dto.Customer
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := dao.UpdateCustomer(customer, id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	} else if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is updated",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not update customer in database",
		})
	}
}

func deleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result, err := dao.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	} else if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is deleted",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not deleted customer in database",
		})
	}
}
