package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var CustomerServicePhone = map[string]string{
	"王培伦": "15874035678",
	"王奥":  "15650708414",
	"尹晨旭": "13032684686",
	"吴舒汀": "17611213883",
	"崔星林": "17717757039",
	"李斌":  "15011332744",
	"王红":  "15810025981",
	"王硕":  "18511280705",
	"吴彬":  "13381040293",
	"杨慧":  "15048008804",
	"杨学成": "13760347163",
	"李晓升": "17600642139",
	"孙凯":  "13052363610",
	"田羽":  "18951841986",
	"王杭渝": "17858226798",
}

var ServicerPhone string //定义一个工单处理人对应的手机号，用于企业微信机器人@使用

type WorkOrder struct {
	OrderID         string `json:"工单id"`
	Customer        string `json:"客户"`
	CustomerService string `json:"受理客服"`
	Subject         string `json:"主题"`
	Priority        string `json:"优先级"`
	WorkUrl         string `json:"工单链接"`
	CellPhone       string `json:"客户手机号"`
	CreateTime      string `json:"创建时间"`
}

// 企业微信机器人消息格式
type WXWorkMsg struct {
	MsgType string  `json:"msgtype"`
	Text    *WXText `json:"text"`
}

// 企业微信发送消息格式
type WXText struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_mobile_list"`
}

func ReConstructPost(w http.ResponseWriter, r *http.Request) {
	var order WorkOrder
	var builder strings.Builder //定义的企业微信机器人发送的字符串

	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid Method!")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	builder.WriteString("工单已分配，请注意及时查看!")
	builder.WriteString("\n")
	builder.WriteString("工单id: ")
	builder.WriteString(order.OrderID)
	builder.WriteString("\n")
	builder.WriteString("客户: ")
	builder.WriteString(order.Customer)
	builder.WriteString("\n")
	builder.WriteString("受理客服: ")
	builder.WriteString(order.CustomerService)
	builder.WriteString("\n")
	builder.WriteString("工单链接: ")
	builder.WriteString(order.WorkUrl)
	builder.WriteString("\n")
	builder.WriteString("工单主题: ")
	builder.WriteString(order.Subject)
	builder.WriteString("\n")
	builder.WriteString("客户手机号: ")
	builder.WriteString(order.CellPhone)
	// msg := "工单已分配,请注意及时查看!" + "工单id: " + order.OrderID + ", 客户: " + order.Customer + ", 受理客服: @" + order.CustomerService
	msg := builder.String()

	ServicerPhone = "0"
	for i, v := range CustomerServicePhone {
		if order.CustomerService == i {
			ServicerPhone = v
			break
		}
	}
	sendMsgToWxWorkRobot(msg)
}
func sendMsgToWxWorkRobot(msg string) {
	wxWorkRobotURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=5f4a82bf-ad78-4110-9b05-2bf9623690aa"

	// 创建企业微信机器人的消息
	wxMsg := &WXWorkMsg{
		MsgType: "text",
		Text: &WXText{
			Content:       msg,
			MentionedList: []string{ServicerPhone}, // 这里填入需要@的人的手机号
		},
	}

	jsonValue, _ := json.Marshal(wxMsg)

	fmt.Println(string(jsonValue))

	resp, err := http.Post(wxWorkRobotURL, "application/json", bytes.NewReader(jsonValue))
	if err != nil {
		fmt.Printf("Post data to WeChat Work robot failed: %v\n", err)
	}
	defer resp.Body.Close()
}
func main() {
	http.HandleFunc("/autoRemind", ReConstructPost) // 当有POST请求到达/autoRemind时，调用ReConstructPost函数处理请求体以及发送请求到企业微信机器人进行通知工单处理人
	err := http.ListenAndServe(":8082", nil)        // 在8082端口启动http服务器
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
