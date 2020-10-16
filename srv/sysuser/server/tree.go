package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ITree interface {
	EditTree(req *sysuser.TreeReq, resp *sysuser.EditResp) error
	DelTree(req *sysuser.DelReq, resp *sysuser.EditResp) error
}

func NewTree() ITree {
	return &Tree{}
}

type Tree struct{}

func (*Tree) DelTree(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysTree{}, req.Id).Error
}

func (*Tree) EditTree(req *sysuser.TreeReq, resp *sysuser.EditResp) error {
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
		resp.Id = Tree.Id
		return db.Updates(Tree).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Tree)
		Tree.Id = mzjuuid.WorkerDefault()
		resp.Id = Tree.Id
		return db.Create(Tree).Error
	}
}
