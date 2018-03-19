package Request

import (
	"net/http"
	"bytes"
	"unsafe"
	"io/ioutil"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
)

//
func PostByJson(url, json string) (string) {
	reader := bytes.NewReader([]byte(json))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println("请求失败", err)
		return ""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("请求失败", err)
		return ""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("请求失败", err)
		return ""
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}

func SendMail(account, password, smtpUrl string, port int, params map[string]interface{},contentType, body string) string {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", params["from"].(string), params["fromname"].(string))
	m.SetHeader("To", strings.Join(params["to"].([]string), ","))
	if params["cc"] != nil {

		m.SetHeaders(map[string][]string{
			"Cc":params["cc"].([]string),
		})
	}
	if params["bcc"]!= nil {

		m.SetHeaders(map[string][]string{
			"Bcc":params["bcc"].([]string),
		})
	}
	m.SetHeader("Subject", params["subject"].(string))
	m.SetBody(contentType, body)
	d := gomail.NewPlainDialer(smtpUrl, port, account, password)
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return "err"
	}

	log.Println("done.发送成功")
	return "ok"
}
