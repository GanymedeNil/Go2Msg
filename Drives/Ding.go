package Drives

import (
	"github.com/GanymedeNil/Go2Msg/Context"
	"encoding/json"
	"github.com/GanymedeNil/Go2Msg/Tmpl"
	"github.com/GanymedeNil/Go2Msg/Request"
)

type Ding struct {
	//dingding 中频道绑定的webhookurl
	WebHookUrl string
}
type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (d *Ding) Send(ctx *Context.MsgContext) string {
	data := Tmpl.Tmpl(ctx.Params["para"].(map[string]string), ctx.Tpl)
	res := Request.PostByJson(d.WebHookUrl, data)
	var result Result
	json.Unmarshal([]byte(res), &result)
	return result.Errmsg
}
