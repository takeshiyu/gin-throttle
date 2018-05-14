package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Throttle used to check the rate limit of incoming request
func Throttle(maxEventsPerSec int, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(context *gin.Context) {
		if limiter.Allow() {
			context.Next()
			return
		}

		context.Error(errors.New("Limit exceeded"))
		context.AbortWithStatus(http.StatusTooManyRequests)
	}
}
