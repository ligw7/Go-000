package dao

import "database/sql"
import "github.com/pkg/errors"

var DaoInstance *dao

func init() {
	DaoInstance = new(dao)
}

type dao struct {
}

func (d *dao) GetUserById(id uint64) (*User, error) {
	if id == 21 {
		zhangSan := User{Name: "张三", Age: 21}
		return &zhangSan, nil
	}
	// Dao层，出现 ErrNoRows
	var err = sql.ErrNoRows
	// Wrap 基础库Error，向上抛
	return nil, errors.Wrapf(err, "Dao: not found, user id = %d", id)
}
