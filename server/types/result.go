package types

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *Result {
	return &Result{Code: 0, Data: data, Msg: "success"}
}

func OkWithMsg(data interface{}, msg string) *Result {
	return &Result{Code: 0, Data: data, Msg: msg}
}

func Error(data interface{}) *Result {
	return &Result{Code: -1, Data: data, Msg: "error"}
}
