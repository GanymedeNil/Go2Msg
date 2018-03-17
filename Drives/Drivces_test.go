package Drives

import (
	"testing"
	"fmt"
	"github.com/GanymedeNil/Go2Msg"
)

//test

func TestDing_Send(t *testing.T) {
	tpl := `{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", 
        "text": "![screenshot](@lADOpwk3K80C0M0FoA) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", 
        "hideAvatar": "0", 
        "btnOrientation": "0", 
        "singleTitle" : "阅读全文",
        "singleURL" : "https://www.dingtalk.com/"
    }, 
    "msgtype": "actionCard"
}`
	ding := &Ding{
		WebHookUrl: "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxx",
	}
	params := map[string]interface{}{
		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	gomsg := go2msg.NewMsgContext(tpl, params, ding)
	status := gomsg.Send()

	fmt.Println(status)
}

func TestEmail_Send(t *testing.T) {
	tpl := `<h1>{{.text}}</h1>`
	params := map[string]interface{}{
		"from":     "name@domain.com",
		"fromname": "name",
		"to":       []string{"name@domain.com"},
		"cc":       []string{},
		"bcc":      []string{},
		"subject":  "go2msg test mail",

		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	email := &Email{
		account:     "name@domain.com",
		password:    "xxxxxxxxxxxxxx",
		smtpUrl:     "smtp.domain.com",
		port:        465,
		contentType: "text/html",
	}
	gomsg := go2msg.NewMsgContext(tpl, params, email)
	status := gomsg.Send()
	fmt.Println(status)

}

func TestSlack_Send(t *testing.T) {
	tpl := `{"text":"{{.text}}"}`
	slack := &Slack{
		WebHookUrl: "https://hooks.slack.com/services/xxxxxxxxx/xxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxx",
	}
	params := map[string]interface{}{
		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	gomsg := go2msg.NewMsgContext(tpl, params, slack)
	status := gomsg.Send()
	fmt.Println(status)
}

//example

func ExampleDing_Send() {
	/**
	tpl 模版格式参考
	https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.K46StK&treeId=257&articleId=105735&docType=1
	 */
	tpl := `{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", 
        "text": "![screenshot](@lADOpwk3K80C0M0FoA) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", 
        "hideAvatar": "0", 
        "btnOrientation": "0", 
        "singleTitle" : "阅读全文",
        "singleURL" : "https://www.dingtalk.com/"
    }, 
    "msgtype": "actionCard"
}`
	ding := &Ding{
		WebHookUrl: "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxx",
	}
	params := map[string]interface{}{
		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	gomsg := go2msg.NewMsgContext(tpl, params, ding)
	status := gomsg.Send()
	fmt.Println(status)
	//Output:ok
}

func ExampleEmail_Send() {
	tpl := `<h1>{{.text}}</h1>`
	params := map[string]interface{}{
		"from":     "name@domain.com",
		"fromname": "name",
		"to":       []string{"name@domain.com"},
		"cc":       []string{},
		"bcc":      []string{},
		"subject":  "go2msg test mail",

		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	email := &Email{
		account:     "name@domain.com",
		password:    "xxxxxxxxxxxxxx",
		smtpUrl:     "smtp.domain.com",
		port:        465,
		contentType: "text/html",
	}
	gomsg := go2msg.NewMsgContext(tpl, params, email)
	status := gomsg.Send()
	fmt.Println(status)
	//Output:ok
}

func ExampleSlack_Send() {
	/**
	tpl 模版格式参考
	https://api.slack.com/docs/message-formatting
	https://api.slack.com/docs/message-attachments
	 */
	tpl := `{"text":"{{.text}}"}`
	slack := &Slack{
		WebHookUrl: "https://hooks.slack.com/services/xxxxxxxxx/xxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxx",
	}
	params := map[string]interface{}{
		"para": map[string]string{
			"text": "Hello,World",
		},
	}
	gomsg := go2msg.NewMsgContext(tpl, params, slack)
	status := gomsg.Send()
	fmt.Println(status)
}