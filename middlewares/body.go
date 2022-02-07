package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mowahaeser/bet/inputs"
)

func Body(input interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		switch input.(type) {
		case inputs.Login:
			input = &inputs.Login{}
		case inputs.Register:
			input = &inputs.Register{}
		}

		if ctx.BindJSON(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		}

		// make it available for next actions
		ctx.Set("input", input)
	}
}
