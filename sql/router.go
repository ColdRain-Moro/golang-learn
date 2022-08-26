package sql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

var r *gin.Engine

func InitRouter(e *gin.Engine) {
	r = e
	InitAuth()
}

func InitAuth() {
	r.POST("/register", func(context *gin.Context) {
		name := context.PostForm("name")
		password := context.PostForm("password")
		spQuestion := context.PostForm("sp_question")
		spAnswer := context.PostForm("sp_answer")
		stmt, err := dB.Prepare("insert into login_plus(id, name, password_md5, password_salt, sp_question, sp_answer) values (?,?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)
			return
		}
		salt := GenerateUUIDStr()
		passwordMd5 := MD5(password + salt)
		_, err = stmt.Exec(0, name, passwordMd5, salt, spQuestion, spAnswer)
		if err != nil {
			log.Fatal(err)
			return
		}
		context.JSON(200, map[string]string{
			"code":    "0",
			"message": "注册成功",
		})
	})
	r.POST("/login", func(context *gin.Context) {
		name := context.PostForm("name")
		password := context.PostForm("password")
		var passwordMd5 string
		var passwordSalt string
		stmt, err := dB.Prepare("select password_md5, password_salt from login_plus where name = ?")
		if err != nil {
			log.Fatal(err)
			return
		}
		rows, err := stmt.Query(name)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()
		rows.Next()
		err = rows.Scan(&passwordMd5, &passwordSalt)
		if err != nil {
			log.Fatal(err)
			return
		}
		if passwordMd5 == MD5(password+passwordSalt) {
			context.SetCookie("gin_cookie", "", 3600, "/", "", false, true)
			context.JSON(200, map[string]string{
				"code":    "0",
				"message": "登录成功",
			})
		} else {
			context.JSON(200, map[string]string{
				"code":    "-1",
				"message": "账号或密码错误",
			})
		}
	})
}

func GenerateUUIDStr() string {
	rand.Seed(time.Now().Unix())
	var randomBytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		randomBytes[i] = byte(rand.Intn(128))
	}
	// 摘自 java uuid 的生成
	randomBytes[6] &= 0x0f /* clear version        */
	randomBytes[6] |= 0x40 /* set to version 4     */
	randomBytes[8] &= 0x3f /* clear variant        */
	randomBytes[8] |= 0x80 /* set to IETF variant  */
	return hex.EncodeToString(randomBytes)
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
