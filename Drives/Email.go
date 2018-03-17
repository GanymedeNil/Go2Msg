package Drives

import (
	"github.com/GanymedeNil/Go2Msg/Context"
	"github.com/GanymedeNil/Go2Msg/Request"
	"github.com/GanymedeNil/Go2Msg/Tmpl"
)

type Email struct {
	/**
	account, password 	发送邮件的账号密码
	smtpUrl, port       邮件服务商smtp接口与端口
	contentType
		text/html       邮件内容类型为html
		text/plain      邮件内容类型为纯文本
	 */
	account, password, smtpUrl, contentType string
	port                                    int
}

func (e *Email) Send(ctx *Context.MsgContext) string {
	body := Tmpl.Tmpl(ctx.Params["para"].(map[string]string), ctx.Tpl)
	res := Request.SendMail(e.account, e.password, e.smtpUrl, e.port, ctx.Params, e.contentType, body)
	return res

}
