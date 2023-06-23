package er

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware global handle *gin.Context error middleware
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()

			switch err.Err.(type) {
			case *Error:
				appError := err.Err.(*Error)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				errUnknown := New(http.StatusInternalServerError, 500999, err.Error(), "Unknown error")
				c.AbortWithStatusJSON(errUnknown.Status, errUnknown)
				break
			}
		}()

		c.Next()
	}
}
