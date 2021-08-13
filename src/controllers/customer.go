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

// @title Customer Service
// @version 1.0
// @description This is a sample server Customer Service
// @termsOfService http://swagger.io/terms/

// @contact.name Saber Azizi
// @contact.url http://www.swagger.io/support
// @contact.email saberazizi66@yahoo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9696
// @BasePath /api/v1
// @query.collection.format multi
// @x-extension-openapi {"example": "value on a json format"}
func StartServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	route = gin.Default()
	route.Use(cors())

	//docs.SwaggerInfo.Title = "Swagger Example API"
	//docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/v2"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

// findAllCustomers godoc
// @Summary findAllCustomers
// @Description findAllCustomers
// @Produce  json
// @Success 200 {object} model.Account
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /customers/findAll [get]
func findAllCustomers(c *gin.Context) {
	customers, err := dao.FindAllCustomers()
	if err != nil {
		c.AbortWithError(404, err)
	} else {
		c.JSON(200, gin.H{
			"customers": customers,
		})
	}
}

// findCustomerById godoc
// @Summary findCustomerById
// @Description get string by ID
// @ID findCustomerById
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /customers/findById/{id} [get]
func findCustomerById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	customer, err := dao.FindCustomerById(id)
	if err != nil {
		c.AbortWithError(404, err)
	} else {
		c.JSON(200, gin.H{
			"customer": customer,
		})
	}
}

// addCustomer godoc
// @Summary addCustomer
// @Description addCustomer
// @Accept  json
// @Produce  json
// @Success 200 {object}  json
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /customers/add/ [post]
func addCustomer(c *gin.Context) {
	var customer dto.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	result, err := dao.AddCustomer(customer)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	}
	if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is created",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not insert customer in database",
		})
	}
}

// updateCustomer godoc
// @Summary updateCustomer
// @Description updateCustomer
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object}  json
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /customers/update/{id} [put]
func updateCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var customer dto.Customer
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	result, err := dao.UpdateCustomer(customer, id)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	} else if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is updated",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not update customer in database",
		})
	}
}

// deleteCustomer godoc
// @Summary deleteCustomer
// @Description deleteCustomer
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object}  json
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /customers/delete/{id} [delete]
func deleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	result, err := dao.DeleteCustomer(id)
	if err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
	} else if result {
		c.JSON(http.StatusCreated, gin.H{
			"code":    0,
			"message": "customer is deleted",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"code":    -1,
			"message": "sorry can not deleted customer in database",
		})
	}
}
