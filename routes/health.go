package routes

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tirlochanarora16/url_shortner/database"
	"github.com/tirlochanarora16/url_shortner/pkg"
)

func checkHealth(c *gin.Context) {
	databaseStatus := "ok"
	redisStatus := "ok"

	if err := database.DB.Ping(); err != nil {
		databaseStatus = "unreachable"
	}

	if _, err := pkg.Rdb.Ping(pkg.Ctx).Result(); err != nil {
		redisStatus = "unreachable"
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(200, gin.H{
		"status":          "ok",
		"timestamp":       time.Now().Format(time.RFC3339),
		"database_status": databaseStatus,
		"redis_status":    redisStatus,
		"goroutines":      runtime.NumGoroutine(),
		"alloc":           m.Alloc / 1024, // KB
		"total_alloc":     m.TotalAlloc / 1024,
		"sys":             m.Sys / 1024,
		"num_gc":          m.NumGC,
		"go_version":      runtime.Version(),
		"os":              runtime.GOOS,
		"architecture":    runtime.GOARCH,
	})
}
