package main

type ApiResponse struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
