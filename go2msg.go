//此包封装了第三方信息通知平台
package go2msg

import (
	"github.com/GanymedeNil/Go2Msg/Context"
)

func NewMsgContext(tpl string, params map[string]interface{}, msgdrive Context.MsgDrive) *Context.MsgContext {
	return &Context.MsgContext{
		Tpl:    tpl,
		Params: params,
		Drive:  msgdrive,
	}
}
