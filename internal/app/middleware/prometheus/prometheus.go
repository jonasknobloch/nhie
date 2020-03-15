package prometheus

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"strings"
)

func UseWithAuth(e *gin.Engine, accounts gin.Accounts) {
	prometheus := ginprometheus.NewPrometheus("gin")

	// preserving low label cardinality for the request counter
	prometheus.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "id" && p.Value != "random" {
				url = strings.Replace(url, p.Value, ":id", 1)
				break
			}
		}
		return url
	}

	prometheus.UseWithAuth(e, accounts)
}
