package Drives

import (
	"github.com/GanymedeNil/Go2Msg/Context"
	"github.com/GanymedeNil/Go2Msg/Tmpl"
	"github.com/GanymedeNil/Go2Msg/Request"
)

type Slack struct {
	//slack app中频道绑定的webhookurl
	WebHookUrl string
}

func (s *Slack) Send(ctx *Context.MsgContext) string {
	data := Tmpl.Tmpl(ctx.Params["para"].(map[string]string), ctx.Tpl)
	res := Request.PostByJson(s.WebHookUrl, data)
	return res
}
