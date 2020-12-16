package service

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"week04/internal/dao"
	"week04/internal/model"
)

type Service struct {
	dao dao.Dao
}

func (s *Service) GetUserInfo(studentId uint64) (*model.User, error) {
	m, err := s.dao.GetUserById(studentId)
	// 隐藏下层代码逻辑错误，转译为业务层case
	if err != nil {
		// 打印错误堆栈,handle errors once
		fmt.Printf("Service: %+v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("the user is not exist.")
		}
		// other else
		return nil, errors.New("something was wrong.")
	}
	return m, nil
}
