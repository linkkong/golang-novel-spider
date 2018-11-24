package wx

import (
	"fmt"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

const (
	OpenID     = "your openid"
	AppID      = "your appid"
	AppSecret  = "your appsecret"
	TemplateID = "your templateid"
	Token      = "15_LfHiyt6fR2YfApiMdIiTYRJkRpjugHC_Dv2XuTIuW5e8f02Tnnpm9ceNGw_cZlCOXzy7JlZ-GmAUXWlU9oJ5puv_1k3rMBjVyT89YPKa6owfSvbRX7KSqdLi3AVGDBGtrVd5_YMsR0KII3KYYSScAAASMM"
)

type AccessToken struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type TemplateMsg struct {
	Touser     string            `json:"touser"`
	TemplateID string            `json:"template_id"`
	Url        string            `json:"url"`
	Data       map[string]Detail `json:"data"`
}

type Detail struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

//微信客服消息，参考 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433751277
func SendTemplateMsg(token string, link string, first string, keyword1 string) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)

	color := "#173177"
	var msg TemplateMsg

	msg.Touser = OpenID
	msg.TemplateID = TemplateID
	msg.Url = link
	msg.Data = map[string]Detail{
		"first":    Detail{Value: first, Color: color},
		"keyword1": Detail{Value: keyword1, Color: color},
		"remark":   Detail{Value: "请注意查收", Color: color},
	}

	jsonStr, errs := json.Marshal(msg) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		return
	}
	//fmt.Println(string(jsonStr)) //byte[]转换成string 输出
	//return
	request := gorequest.New()
	resp, body, errs1 := request.Post(url).
		Send(string(jsonStr)).
		End()
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", body)
	fmt.Printf("%+v", errs1)
}

// 发送客服消息，只能48小时
func SendMsg(msg string, token string) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", token)
	request := gorequest.New()
	jsonData := fmt.Sprintf(`{"touser":"%s","msgtype":"text", "text":{"content":"%s"}}`, OpenID, msg)
	resp, body, errs := request.Post(url).
		Send(jsonData).
		End()
	fmt.Printf("%+v", resp)
	fmt.Printf("%+v", body)
	fmt.Printf("%+v", errs)
}

//获取微信token
func GetToken() string {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, AppSecret)
	var tk AccessToken
	gorequest.New().Get(url).EndStruct(&tk)
	return tk.AccessToken
}
