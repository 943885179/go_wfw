package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ITree interface {
	EditTree(req *sysuser.SysTree, resp *sysuser.EditResp) error
	DelTree(req *sysuser.DelReq, resp *sysuser.EditResp) error
	TreeList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewTree() ITree {
	return &Tree{}
}

type Tree struct{}

func (t *Tree) TreeList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var ts []models.SysTree
	db := Conf.DbConfig.New().Model(&models.SysTree{}).Where("p_id=0")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&ts)
	for _, role := range ts {
		var r sysuser.SysTree
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Tree) DelTree(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysTree{}, req.Id).Error
}

func (*Tree) EditTree(req *sysuser.SysTree, resp *sysuser.EditResp) error {
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
