package types

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(data interface{}) *Result {
	return &Result{Code: 0, Data: data, Message: "success"}
}

func OkWithMsg(data interface{}, msg string) *Result {
	return &Result{Code: 0, Data: data, Message: msg}
}

func Error(data interface{}) *Result {
	return &Result{Code: -1, Data: data, Message: "error"}
}
