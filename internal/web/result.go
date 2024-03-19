package web

type Result struct {
	//code 返回的状态码 无返回 代表正确 返回 4 代表 用户错误 返回 5 代表系统错误
	Code int `json:"code"`
	//msg 返回的内容
	Mag string `json:"mag"`
	//data 需要返回的一些数据
	Date any `json:"date"`
}
