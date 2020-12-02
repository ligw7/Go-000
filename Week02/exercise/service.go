package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

var ServiceInstance *service

func init() {
	ServiceInstance = new(service)
}

type serviceErr struct {
	err  error
	Code int
	Msg  string
}

func NewServiceErr(err error, code int, msg string) *serviceErr {
	return &serviceErr{
		err:  err,
		Code: code,
		Msg:  msg,
	}
}

type service struct {
}

func (s *service) GetUserInfo(studentId uint64) (*User, *serviceErr) {
	m, err := DaoInstance.GetUserById(studentId)
	// 隐藏下层代码逻辑错误，转译为业务层case
	if err != nil {
		// 打印错误堆栈,handle errors once
		fmt.Printf("service: %+v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewServiceErr(err, 1, "the user is not exist.")
		}
		// other else
		return nil, NewServiceErr(err, 2, "something was wrong.")
	}
	return m, nil
}
