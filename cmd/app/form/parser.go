package form

import (
	"github.com/gin-gonic/gin"
)

func GetFormValue(c *gin.Context, named_attribute string)(string) {
	/* Get id for a named attribute of the input tag */
	var value = c.Request.PostFormValue(named_attribute)
	return value
}

