package gin_learn

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var auth = func(c *gin.Context) {
	//获取cookie
	value, err := c.Cookie("gin_cookie")
	if err != nil {
		c.JSON(403, gin.H{
			"message": "认证失败,没有cookie",
		})
		//认证失败
		//终止的意思，也就是说执行该函数，会终止后面所有的该请求下的函数
		c.Abort()
	} else {
		//将获取到的cookie的值写入上下文
		c.Set("cookie", value)
		//这里随便介绍一下next方法
		//挂起继续向下走，然后执行完成下面的函数，会反过来最后执行该中间件
		c.Next()
		v, _ := c.Get("next")
		fmt.Println(v)
	}
}

func InitLogin(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "123" && password == "321" {
			//对着函数摁一下ctrl+b看看对应参数
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			c.JSON(200, gin.H{
				"msg": "login successfully",
			})
		} else {
			c.JSON(403, gin.H{
				"message": "认证失败,账号密码错误",
			})
		}
	})
}
