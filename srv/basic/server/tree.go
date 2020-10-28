package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ITree interface {
	EditTree(req *dbmodel.SysTree, resp *dbmodel.Id) error
	DelTree(req *dbmodel.Id, resp *dbmodel.Id) error
	TreeList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	TreeById(id *dbmodel.Id, tree *dbmodel.SysTree) error
	TreeTree(empty *empty.Empty, resp *dbmodel.TreeResp) error
}

func NewTree() ITree {
	return &Tree{}
}

type Tree struct{}

func (r *Tree) TreeTree(empty *empty.Empty, resp *dbmodel.TreeResp) error {
	var data []models.SysTree
	db := Conf.DbConfig.New().Model(&models.SysMenu{}).Where("p_id=0")
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	if err := db.Find(&data).Error; err != nil {
		return err
	}

	for _, m := range data {
		var r dbmodel.Tree
		mzjstruct.CopyStruct(&m, &r)
		resp.Data = append(resp.Data, &r)
	}
	return nil
}

func (g *Tree) TreeById(id *dbmodel.Id, tree *dbmodel.SysTree) error {
	return Conf.DbConfig.New().Model(&models.SysTree{}).First(tree, id.Id).Error
}
func (t *Tree) TreeList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
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
		var r dbmodel.SysTree
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Tree) DelTree(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysTree{}, req.Id).Error
}

func (*Tree) EditTree(req *dbmodel.SysTree, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	Tree := &models.SysTree{}
	if len(req.Id) > 0 { //修改0
		if err := db.First(Tree, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Tree)
		resp.Id = Tree.Id
		Tree.Title = Tree.Text
		return db.Updates(Tree).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Tree)
		Tree.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		Tree.Key = Tree.Id
		resp.Id = Tree.Id
		Tree.Title = Tree.Text
		return db.Create(Tree).Error
	}
}
