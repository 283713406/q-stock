// 评论留言

package routes

import (
	"net/http"

	"github.com/283713406/q-stock/version"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Comment godoc
func Comment(c *gin.Context) {
	data := gin.H{
		"Env":       viper.GetString("env"),
		"Version":   version.Version,
		"PageTitle": "Q-STOCK | 留言",
	}
	c.HTML(http.StatusOK, "comment.html", data)
	return
}
