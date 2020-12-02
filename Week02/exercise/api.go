package main

// 对外提供 http 或 grpc 接口，路由到 service Handler;现在只通过函数调用模拟API层
type api struct {
}

func NewAPI() *api {
	return &api{}
}

func (a *api) GetUserInfo(studentId uint64) *ApiResponse {
	apiResponse := &ApiResponse{}
	m, serviceErr := ServiceInstance.GetUserInfo(studentId)
	if serviceErr != nil {
		apiResponse.ErrCode = serviceErr.Code
		apiResponse.ErrMsg = serviceErr.Msg
		apiResponse.Data = struct{}{}
		return apiResponse
	}
	apiResponse.ErrCode = 0
	apiResponse.ErrMsg = "OK"
	apiResponse.Data = m
	return apiResponse
}
