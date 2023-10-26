package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetRateLimitMiddleware(storeClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Key is user ip, 2 requests per second
		userIP := c.ClientIP()
		var limit int64 = 2
		interval := 3 * time.Second

		// Create a unique key for each user
		key := fmt.Sprintf("ratelimit:%s", userIP)

		// Use Redis to track the number of requests
		ctx := context.Background()
		count, err := storeClient.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		// If the count exceeds the limit, reject the request
		if count > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "You're being rate limited"})
			c.Abort()
			return
		}

		// Set the expiration time for the key
		if count == 1 {
			_, err := storeClient.Expire(ctx, key, interval).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
				return
			}
		}

		// Continue with the request if it is within the rate limit
		c.Next()
	}
}
