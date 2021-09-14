package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driver"
	"net/http"
)

type UserController struct {
	service driver.UserService
	router  *gin.RouterGroup
}

func NewUserController(service driver.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller UserController) InitRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	router := v1.Group("/users")

	controller.router = router

	controller.create()
	controller.findOne()
}

func (controller UserController) create() {
	controller.router.POST("", func(c *gin.Context) {
		var body RequestUser
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Error(fmt.Errorf("erro ao deserializar json do request: %w", err))
			return
		}

		userSaved, err := controller.service.SaveUser(requestUserToUser(body))
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userSaved)
	})
}

func (controller UserController) findOne() {
	controller.router.GET("/email/:email", func(c *gin.Context) {
		email := c.Param("email")

		userFound, err := controller.service.FindByEmail(email)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userFound)
	})
}

func requestUserToUser(user RequestUser) domain.User {
	return domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
}
