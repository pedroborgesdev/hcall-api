package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// GetRealIP extracts the real IP address from the request
func GetRealIP(c *gin.Context) string {
	// Try X-Real-IP header first
	ip := c.GetHeader("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Try X-Forwarded-For header
	ip = c.GetHeader("X-Forwarded-For")
	if ip != "" {
		// Get the first IP in the list (client's original IP)
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	// Try CF-Connecting-IP header (Cloudflare)
	ip = c.GetHeader("CF-Connecting-IP")
	if ip != "" {
		return ip
	}

	// Try True-Client-IP header
	ip = c.GetHeader("True-Client-IP")
	if ip != "" {
		return ip
	}

	// Fallback to RemoteAddr
	ip, _, _ = strings.Cut(c.Request.RemoteAddr, ":")
	return ip
}
