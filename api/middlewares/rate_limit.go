package middlewares

import (
	"context"
	"sync"
	"time"

	"hcall/api/config"
	"hcall/api/logger"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips        *sync.Map
	mu         sync.Mutex
	r          rate.Limit
	b          int
	lastGC     time.Time
	gcInterval time.Duration
	ctx        context.Context
}

func NewIPRateLimiter(ctx context.Context, r rate.Limit, b int) *IPRateLimiter {
	if b < 1 {
		b = 1
		logger.Warning("Rate Limiter Middleware: Burst ajustado para 1 (mínimo permitido)", map[string]interface{}{
			"rate":  r,
			"burst": b,
		})
	}

	return &IPRateLimiter{
		ips:        &sync.Map{},
		r:          r,
		b:          b,
		lastGC:     time.Now(),
		gcInterval: 5 * time.Minute,
		ctx:        ctx,
	}
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.runGCIfNeeded()

	limiter, loaded := i.ips.Load(ip)
	if !loaded {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips.Store(ip, limiter)
		logger.Debug("Rate Limiter Middleware: New Limiter created", map[string]interface{}{
			"ip":    ip,
			"rate":  i.r,
			"burst": i.b,
			"type":  "rate_limiter",
		})
	}
	return limiter.(*rate.Limiter)
}

func (i *IPRateLimiter) runGCIfNeeded() {
	i.mu.Lock()
	defer i.mu.Unlock()

	if time.Since(i.lastGC) < i.gcInterval {
		return
	}

	cleaned := 0
	i.ips.Range(func(key, value interface{}) bool {
		select {
		case <-i.ctx.Done():
			return false
		default:
			if value.(*rate.Limiter).Tokens() == float64(i.b) {
				i.ips.Delete(key)
				cleaned++
			}
			return true
		}
	})

	if cleaned > 0 {
		logger.Debug("Rate Limiter Middleware: GC Executed", map[string]interface{}{
			"ips_removidos": cleaned,
			"total_ips":     i.countIPs(),
		})
	}
	i.lastGC = time.Now()
}

func (i *IPRateLimiter) countIPs() int {
	count := 0
	i.ips.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

var (
	limiter *IPRateLimiter
	once    sync.Once
)

func initRateLimiter(ctx context.Context) {
	requestsPerMinute := config.AppConfig.RateLimitRequests
	if requestsPerMinute <= 0 {
		requestsPerMinute = 60
		logger.Info("Rate Limiter Middleware: Using default value", map[string]interface{}{
			"config": "RateLimitRequests",
			"value":  requestsPerMinute,
		})
	}

	burst := config.AppConfig.RateLimitWindow
	if burst <= 0 {
		burst = 5
		logger.Info("Rate Limiter Middleware: Using default value", map[string]interface{}{
			"config": "RateLimitWindow",
			"value":  burst,
		})
	}

	requestsPerSecond := rate.Limit(float64(requestsPerMinute) / 60.0)
	if requestsPerSecond < 0.0167 {
		requestsPerSecond = 0.0167
		logger.Info("Rate Limiter Middleware: Adjusted minimum rate", map[string]interface{}{
			"original": requestsPerSecond,
			"adjusted": 0.0167,
		})
	}

	logger.Info("Rate Limiter Middleware: Initial Setup", map[string]interface{}{
		"requests_per_minute": requestsPerMinute,
		"requests_per_second": requestsPerSecond,
		"burst":               burst,
	})

	limiter = NewIPRateLimiter(ctx, requestsPerSecond, burst)
}

func RateLimitMiddleware(ctx context.Context) gin.HandlerFunc {
	once.Do(func() {
		initRateLimiter(ctx)
	})

	return func(c *gin.Context) {
		if config.AppConfig.RateLimitRequests <= 0 {
			c.Next()
			return
		}

		ip := utils.GetRealIP(c)
		l := limiter.GetLimiter(ip)

		// Tenta reservar um token
		reserve := l.Reserve()
		if !reserve.OK() {
			logger.Warning("Rate Limiter Middleware: It was not possible to reserve token", map[string]interface{}{
				"ip":     ip,
				"tokens": l.Tokens(),
				"burst":  limiter.b,
			})
			utils.SendError(c, utils.CodeRateLimitError, utils.MsgRateLimitError, nil)
			c.Abort() // Garante que nenhum outro handler será executado
			return
		}

		// Verifica se precisa esperar
		if delay := reserve.Delay(); delay > 0 {
			reserve.Cancel() // Libera o token reservado pois vamos bloquear
			logger.Warning("Rate Limiter Middleware: Exceeded limit", map[string]interface{}{
				"ip":          ip,
				"tokens":      l.Tokens(),
				"burst":       limiter.b,
				"rate":        limiter.r,
				"retry_after": delay.String(),
				"path":        c.Request.URL.Path,
				"method":      c.Request.Method,
			})
			c.Header("Retry-After", delay.String())
			utils.SendError(c, utils.CodeRateLimitError, utils.MsgRateLimitError, nil)
			c.Abort() // Garante que nenhum outro handler será executado
			return
		}

		logger.Debug("Rate Limiter Middleware: Allowed request", map[string]interface{}{
			"ip":     ip,
			"tokens": l.Tokens(),
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})

		// Se chegou aqui, a requisição pode prosseguir
		c.Next()
	}
}
