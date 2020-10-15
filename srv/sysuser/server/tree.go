package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ITree interface {
	EditTree(req *sysuser.TreeReq) error
	DelTree(req *sysuser.DelReq) error
}

func NewTree() ITree {
	return &Tree{}
}

type Tree struct{}

func (*Tree) DelTree(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	return db.Delete(models.SysTree{}, req.Id).Error
}

func (*Tree) EditTree(req *sysuser.TreeReq) error {
	db := Conf.DbConfig.New()
	Tree := &models.SysTree{}
	if req.Id > 0 { //修改0
		if err := db.First(Tree, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Tree)
		return db.Updates(Tree).Error
	} else { //添加
		Tree.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, Tree)
		return db.Create(Tree).Error
	}
}
