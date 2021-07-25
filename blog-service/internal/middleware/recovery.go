package middleware

import (
	"fmt"
	"playground/blog-service/global"
	"playground/blog-service/pkg/app"
	"playground/blog-service/pkg/email"
	"playground/blog-service/pkg/errcode"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Errorf(s, err)

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("例外拋出，發生時間: %d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
