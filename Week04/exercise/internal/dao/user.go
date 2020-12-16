package dao

import "database/sql"
import "github.com/pkg/errors"
import "week04/internal/model"

var daoInstance *Dao

func init() {
	daoInstance = new(Dao)
}

func NewDao() *Dao {
	return daoInstance
}

type Dao struct {
}

func (d *Dao) GetUserById(id uint64) (*model.User, error) {
	if id == 21 {
		zhangSan := model.User{Name: "张三", Age: 21}
		return &zhangSan, nil
	}
	// Dao层，出现 ErrNoRows
	var err = sql.ErrNoRows
	// Wrap 基础库Error，向上抛
	return nil, errors.Wrapf(err, "Dao: not found, user id = %d", id)
}
