package gin_learn

import "github.com/gin-gonic/gin"

func Run() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()
	// 初始化登录模块
	InitLogin(r)
	//在中间放入鉴权中间件
	r.GET("/hello", auth, func(c *gin.Context) {
		cookie, _ := c.Get("cookie")
		str := cookie.(string)
		c.String(200, "hello world"+str)
		//测试next函数
		c.Set("next", "test next")
	})
	// By default, it serves on :8080 unless a
	// PORT environment variable was defined.
	err := r.Run()
	if err != nil {
		return
	}
	// router.Run(":3000") for a hard coded port
}
