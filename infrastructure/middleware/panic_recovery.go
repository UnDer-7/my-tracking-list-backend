package middleware

import (
	"crud-test/core/app_error"
	"github.com/gin-gonic/gin"
)

func HandlePanicRecovery(c *gin.Context, _ interface{}) {
	interServerErr := app_error.ThrowInternalServerError("a panic has occurred", nil)
	c.AbortWithStatusJSON(interServerErr.StatusCode, interServerErr)
}
