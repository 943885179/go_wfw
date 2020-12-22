package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"

	"github.com/golang/protobuf/ptypes"
)

type ITree interface {
	EditTree(req *dbmodel.SysTree, resp *dbmodel.Id) error
	DelTree(req *dbmodel.Id, resp *dbmodel.Id) error
	TreeList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	TreeById(id *dbmodel.Id, tree *dbmodel.SysTree) error
	TreeTree(resp *dbmodel.TreeResp) error
	TreeByType(req *basic.TreeType, resp *dbmodel.TreeResp) error
	TreeOneByCode(code string, resp *dbmodel.SysTree) error
}

func NewTree() ITree {
	return &Tree{}
}

type Tree struct{}

func (t *Tree) TreeOneByCode(code string, resp *dbmodel.SysTree) error {
	return Conf.DbConfig.New().Model(&models.SysTree{}).Where("code=?", code).First(resp).Error
}

func (r *Tree) TreeTree(resp *dbmodel.TreeResp) error {
	var data []models.SysTree
	db := Conf.DbConfig.New().Model(&models.SysTree{}).Where("p_id=''")
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
func (g *Tree) TreeByType(req *basic.TreeType, resp *dbmodel.TreeResp) error {
	var data []models.SysTree
	db := Conf.DbConfig.New().Model(&models.SysTree{}).Where("type=? and p_id=''", int32(req.TreeType))
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
func (t *Tree) TreeList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var ts []models.SysTree
	db := Conf.DbConfig.New().Model(&models.SysTree{}) //.Where("p_id=''")
	if len(req.Type) > 0 {
		db = db.Where("type=?", req.Type)
	}
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db = db.Preload("Children")
	/*db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")*/
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
	var codeTree = models.SysTree{}
	db.Model(&models.SysTree{}).Where("code=?", req.Code).First(&codeTree)
	if len(req.Id) > 0 { //修改0
		if req.Id != codeTree.Id {
			return errors.New("编码已存在")
		}
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
		if len(codeTree.Id) > 0 {
			return errors.New("编码已存在")
		}
		mzjstruct.CopyStruct(req, Tree)
		Tree.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		Tree.Key = Tree.Id
		resp.Id = Tree.Id
		Tree.Title = Tree.Text
		return db.Create(Tree).Error
	}
}
