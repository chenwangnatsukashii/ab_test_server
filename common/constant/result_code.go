package constant

type ResponseMsg string

// 返回状态
const (
	SuccessMsg ResponseMsg = "请求成功"
	ErrorMsg   ResponseMsg = "请求失败"
)
