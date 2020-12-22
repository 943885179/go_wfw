package mzjexcelize

import (
	"encoding/json"
	"fmt"
	"qshapi/utils/mzjtime"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//go get github.com/360EntSecGroup-Skylar/excelize/v2 excel操作库
//go get github.com/bitly/go-simplejson json操作库

//CellType 列名称
type CellType int

//TimeFormat 时间格式
type TimeFormat int

const (
	a CellType = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	aa
	ab
	ac
)

func (c CellType) String() string {
	switch c {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	case 4:
		return "E"
	case 5:
		return "F"
	case 6:
		return "G"
	case 7:
		return "H"
	case 8:
		return "I"
	case 9:
		return "J"
	case 10:
		return "K"
	case 11:
		return "L"
	case 12:
		return "M"
	case 13:
		return "N"
	case 14:
		return "O"
	case 15:
		return "P"
	case 16:
		return "Q"
	case 17:
		return "R"
	case 18:
		return "S"
	case 19:
		return "T"
	case 20:
		return "U"
	case 21:
		return "V"
	case 22:
		return "W"
	case 23:
		return "X"
	case 24:
		return "Y"
	case 25:
		return "Z"
	case 26:
		return "AA"
	case 27:
		return "AB"
	case 28:
		return "AC"
	default:
		return "请先定义字段,请先定义"

	}
}

//Excel xcel操作类
type Excel struct {
	FileName    string       //文件名称
	SheetName   string       //sheet名称
	IsHasHead   bool         //是否拥有头部
	CellOptions []CellOption //列对照表
}

//CellOption 列配置
type CellOption struct {
	Name     string   //对应的字段名称
	CName    string   //表头名称
	CellType CellType //所在的行
	//IsTime     bool                 //是否为时间格式
	Type       string             //格式（读取和时间格式写入需要设定）
	TimeFormat mzjtime.TimeFormat //时间格式的话格式类型
}

//========================旧版使用反射处理=======================================================================

//OldWriteByEntitys 将Entity写入excel(需要转成[]interface{}),支持多个stuct
func (e Excel) OldWriteByEntitys(data []interface{}) {
	f := excelize.NewFile()
	index := f.NewSheet(e.SheetName)
	f.SetActiveSheet(index)
	e.headSave(f)
	e.oldSave(f, data)
	f.SaveAs(e.FileName)
}

func (e Excel) oldSave(f *excelize.File, data []interface{}) {
	var headIndex = 0
	if e.IsHasHead {
		headIndex = 1
	}
	for rowIndex, d := range data {
		v := reflect.ValueOf(d)
		t := v.Type()
		k := t.Kind()
		switch k {
		case reflect.Struct:
			for i := 0; i < t.NumField(); i++ {
				var isjb = false //是否匹配
				for _, c := range e.CellOptions {
					//匹配局部
					if strings.ToLower(c.Name) == strings.ToLower(fmt.Sprintf("%s.%s", t.Name(), t.Field(i).Name)) { //这里先匹配对应struct的值
						if c.CName == "" {
							c.CName = c.Name
						}
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+headIndex+1), v.Field(i).Interface())
						isjb = true
						break
					}
				}
				if isjb { //如果局部匹配到了那么不需要做全局匹配了
					continue
				}
				for _, c := range e.CellOptions {
					//匹配全局
					if strings.ToLower(c.Name) == strings.ToLower(t.Field(i).Name) {
						if c.CName == "" {
							c.CName = c.Name
						}
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+headIndex+1), v.Field(i).Interface())
						break
					}
				}
			}
			break
		case reflect.Map:
			break
		}
	}
}

//保存头部
func (e Excel) headSave(f *excelize.File) int {
	var headIndex = 0
	if e.IsHasHead {
		headIndex = 1
		for _, c := range e.CellOptions {
			if c.CName == "" {
				c.CName = c.Name
			}
			f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), headIndex), c.CName)
		}
		f.AutoFilter(e.SheetName, fmt.Sprintf("%s%d", e.CellOptions[0].CellType.String(), 1), fmt.Sprintf("%s%d", e.CellOptions[len(e.CellOptions)-1].CellType.String(), 1), "") //创建条件筛选,必须保证设定的首个必须是开头，尾部必须是结束，例：有3列，ABC,那么必须A是celloptions的第一个值，C为最后一个
	}
	return headIndex
}

//=========================新版使用json处理======================================================================

//Write 新版写入不支持同名多列
func (e Excel) Write(data interface{}) {
	f := excelize.NewFile()
	index := f.NewSheet(e.SheetName)
	f.SetActiveSheet(index)
	e.writeByEntitys(f, "", data)
	f.SaveAs(e.FileName)
}
func (e Excel) writeByEntitys(f *excelize.File, qz string, data interface{}) {

	headIndex := e.headSave(f)
	v := reflect.ValueOf(data)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Slice:
		bt, _ := json.Marshal(v.Interface())
		e.jsonBtSave(f, qz, bt)
		break
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			ck := v.Field(i).Kind()
			var n = fmt.Sprintf("%s.%s", t.Name(), t.Field(i).Name)
			if qz != "" {
				n = fmt.Sprintf("%s.%s.%s", qz, t.Name(), t.Field(i).Name)
			}
			switch ck {
			case reflect.Slice:
				bt, _ := json.Marshal(v.Field(i).Interface())
				e.jsonBtSave(f, n, bt)
				break
			/*case reflect.Struct:
			e.writeByEntitys(t.Name(), v.Field(i).Interface())
			break*/
			case reflect.Struct:
				if v.Field(i).Type().String() == "time.Time" {
					var isjb = false //是否匹配
					for _, co := range e.CellOptions {
						if strings.ToLower(co.Name) == strings.ToLower(n) { //这里先匹配对应struct的值
							timeStr := v.Field(i).Interface().(time.Time).Format(co.TimeFormat.String())
							f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", co.CellType.String(), headIndex+1), timeStr)
							isjb = true
							break
						}
					}
					if isjb { //如果局部匹配到了那么不需要做全局匹配了
						continue
					}
					for _, co := range e.CellOptions {
						//匹配全局
						if strings.ToLower(co.Name) == strings.ToLower(t.Field(i).Name) {
							timeStr := v.Field(i).Interface().(time.Time).Format(co.TimeFormat.String())
							fmt.Println(timeStr)
							f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", co.CellType.String(), headIndex+1), timeStr)
							break
						}
					}
				} else {
					ls := []interface{}{v.Field(i).Interface()}
					e.writeByEntitys(f, n, ls)
				}
				break
			default:
				var isjb = false //是否匹配
				for _, co := range e.CellOptions {
					if strings.ToLower(co.Name) == strings.ToLower(n) { //这里先匹配对应struct的值
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", co.CellType.String(), headIndex+1), v.Field(i).Interface())
						isjb = true
						break
					}
				}
				if isjb { //如果局部匹配到了那么不需要做全局匹配了
					continue
				}
				for _, co := range e.CellOptions {
					//匹配全局
					if strings.ToLower(co.Name) == strings.ToLower(t.Field(i).Name) {
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", co.CellType.String(), headIndex+1), v.Field(i).Interface())
						break
					}
				}
				break
			}
		}
		break
	}
}

//NewWriteByJSONStr json保存为excel
func (e Excel) NewWriteByJSONStr(str string) {
	f := excelize.NewFile()
	index := f.NewSheet(e.SheetName)
	f.SetActiveSheet(index)
	e.headSave(f)
	bt, _ := json.Marshal([]byte(str))
	e.jsonBtSave(f, "", bt)
	f.SaveAs(e.FileName)
}
func (e Excel) jsonBtSave(f *excelize.File, qz string, jsonBt []byte) {
	jstr, _ := simplejson.NewJson(jsonBt)
	rows, _ := jstr.Array()
	var offsetRow = 0 //默认没有头
	if e.IsHasHead {
		offsetRow = 1
	}
	e.rowSave(f, qz, offsetRow, rows)
}
func (e Excel) rowSave(f *excelize.File, qz string, offsetRow int, rows []interface{}) {
	for i, row := range rows {
		if mp, ok := row.(map[string]interface{}); ok {
			if i != 0 {
				offsetRow = offsetRow + 1
			}
			offsetRow = e.mapSave(f, offsetRow, qz, mp)
		}
	}
}

//mapSave
//f 文件
//rowIndex 第几行写入
//qz 前缀匹配
//mp 数据
//returns 返回改修改占用的行数
func (e Excel) mapSave(f *excelize.File, rowIndex int, qz string, mp map[string]interface{}) (result int) {
	result = rowIndex
	for k, val := range mp {
		var n = k
		if qz != "" {
			n = fmt.Sprintf("%s.%s", qz, k)
		}
		if cmp, ok := val.(map[string]interface{}); ok {
			e.mapSave(f, rowIndex, n, cmp)

		} else if cmp, ok := val.([]interface{}); ok {
			e.rowSave(f, n, rowIndex, cmp)
			if result < len(cmp) {
				result = len(cmp) //默认占用的行数将设置为最大数据占用行数
			}
		} else {
			var isjb = false //是否匹配
			for _, c := range e.CellOptions {
				if strings.ToLower(c.Name) == strings.ToLower(n) {
					if strings.ToLower(c.Type) == strings.ToLower("time.Time") { //if c.IsTime { //时间格式替换成特定的时间格式
						//2020-09-14T15:05:22.9111439+08:00
						tm, _ := mzjtime.ParseInlocation(fmt.Sprintf("%s %s", val.(string)[0:10], val.(string)[11:19]), 0) //这句将2020-09-14T15:05:22.9111439+08:00字符串切割成yyyy:MM:dd HH:mm:ss
						ts := mzjtime.Format(tm, c.TimeFormat)
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+1), ts)
					} else {
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+1), val)
					}
					isjb = true
					break
				}
			}
			if isjb {
				continue
			}
			for _, c := range e.CellOptions {
				if strings.ToLower(c.Name) == strings.ToLower(k) && val != nil {
					if strings.ToLower(c.Type) == strings.ToLower("time.Time") { //if c.IsTime { //时间格式替换成特定的时间格式
						tm, _ := mzjtime.ParseInlocation(fmt.Sprintf("%s %s", val.(string)[0:10], val.(string)[11:19]), 0)
						ts := mzjtime.Format(tm, c.TimeFormat)
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+1), ts)
					} else {
						f.SetCellValue(e.SheetName, fmt.Sprintf("%s%d", c.CellType.String(), rowIndex+1), val)
					}
					break
				}
			}
		}
	}
	return result
}

//ReadOne 读取一个
func (e Excel) ReadOne(resp interface{}) error {
	f, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return err
	}
	v := reflect.ValueOf(resp)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Ptr:
		t = t.Elem()
		v = v.Elem()
		for i := 0; i < t.NumField(); i++ {
			var rowIndex = 1 //默认读取第一行
			if e.IsHasHead { //有表头那么读取第二行为内容
				rowIndex = 2
			}
			for _, o := range e.CellOptions {
				if strings.ToLower(o.Name) == strings.ToLower(t.Field(i).Name) || strings.ToLower(fmt.Sprintf("%s.%s", t.Name(), t.Field(i).Name)) == strings.ToLower(o.Name) {
					cellV := f.GetCellValue(e.SheetName, fmt.Sprintf("%s%d", o.CellType.String(), rowIndex)) //单一实体只读取A1行
					vk := v.Field(i).Kind()
					switch vk {
					case reflect.String:
						v.Field(i).Set(reflect.ValueOf(cellV))
						break
					case reflect.Int:
						va, err := strconv.Atoi(cellV)
						if err != nil {
							return fmt.Errorf("%s字段无法解析成int", o.Name)
						}
						v.Field(i).Set(reflect.ValueOf(va))
					case reflect.Int64:
						va, err := strconv.Atoi(cellV)
						if err != nil {
							return fmt.Errorf("%s字段无法解析成int64", o.Name)
						}
						v.Field(i).Set(reflect.ValueOf(int64(va)))
					case reflect.Int32:
						va, err := strconv.Atoi(cellV)
						if err != nil {
							return fmt.Errorf("%s字段无法解析成int32", o.Name)
						}
						v.Field(i).Set(reflect.ValueOf(int32(va)))
					case reflect.Bool:
						va, err := strconv.ParseBool(cellV)
						if err != nil {
							return fmt.Errorf("%s字段无法解析成bool", o.Name)
						}
						v.Field(i).Set(reflect.ValueOf(va))
						break
					case reflect.Struct:
						if t.Field(i).Type.String() == "time.Time" {
							times, err := mzjtime.ParseInlocation(cellV, o.TimeFormat) //time.Parse(o.TimeFormat.String(), cellV)原来的使用这个
							if err != nil {
								return err
							}
							v.Field(i).Set(reflect.ValueOf(times))
						} else {
							//fmt.Errorf("暂时不支持%s类型解析", t.Field(i).Type.String())
						}
						break
					default:
						return fmt.Errorf("暂时不支持%s类型解析", vk)
					}
					break
				}
			}

		}
		break

	default:
		//panic("暂时不支持其他格式的excel读取")
		break
	}

	for i, row := range f.GetRows(e.SheetName) {
		if i == 0 { //第一行为表头
			for _, cell := range row {
				for _, o := range e.CellOptions {
					if o.CName == cell || o.Name == cell {
						continue
					}
				}
			}
		}
	}
	return nil
}

//Read 读取
func (e Excel) Read(resp interface{}, respMd interface{}) error {
	f, err := excelize.OpenFile(e.FileName)
	if err != nil {
		return err
	}
	v := reflect.ValueOf(respMd)
	t := v.Type()
	switch t.Kind() {
	case reflect.Struct:
		newCellOptions := []CellOption{}
		for i := 0; i < t.NumField(); i++ {
			for _, o := range e.CellOptions {
				if strings.ToLower(o.Name) == strings.ToLower(t.Field(i).Name) || strings.ToLower(o.Name) == strings.ToLower(fmt.Sprintf("%s.%s", t.Name(), t.Field(i).Name)) {
					o.Name = t.Field(i).Name
					jsonTag := t.Field(i).Tag.Get("json")
					if jsonTag != "" {
						name := strings.Split(jsonTag, ";")[0]
						if name != "" {
							o.Name = name
						}
					}
					o.Type = t.Field(i).Type.String()
					newCellOptions = append(newCellOptions, o)
					break
				}
			}
		}
		e.CellOptions = newCellOptions
		break
	default:
		break
	}
	jsonMp := make([]map[string]interface{}, 0)
	for rowIndex := 0; rowIndex < len(f.GetRows(e.SheetName)); rowIndex++ {
		if rowIndex == 0 && e.IsHasHead {
			continue
		}
		smp := map[string]interface{}{}
		for _, o := range e.CellOptions {
			cellV := f.GetCellValue(e.SheetName, fmt.Sprintf("%s%d", o.CellType.String(), rowIndex+1))
			switch o.Type {
			case reflect.String.String():
				smp[o.Name] = cellV
			case reflect.Bool.String():
				sv, err := strconv.ParseBool(cellV)
				if err != nil {
					return err
				}
				smp[o.Name] = sv
			case reflect.Int.String():
				sv, err := strconv.Atoi(cellV)
				if err != nil {
					return err
				}
				smp[o.Name] = sv
			case reflect.Int32.String():
				sv, err := strconv.Atoi(cellV)
				if err != nil {
					return err
				}
				smp[o.Name] = int32(sv)
			case reflect.Int64.String():
				sv, err := strconv.Atoi(cellV)
				if err != nil {
					return err
				}
				smp[o.Name] = int64(sv)
			case "time.Time":
				sv, err := mzjtime.ParseInlocation(cellV, o.TimeFormat)
				if err != nil {
					return err
				}
				smp[o.Name] = sv
			default:
				smp[o.Name] = cellV
			}
		}
		if smp != nil {
			jsonMp = append(jsonMp, smp)
		}
	}
	bt, _ := json.Marshal(jsonMp)
	json.Unmarshal(bt, resp)
	return nil
}
func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

//===============测试============================
func test1() {
	type Test1 struct {
		Name string
		Tag  string
	}
	type Test2 struct {
		Name  string
		Tag   string
		Test1 Test1
	}
	type Test struct {
		Name   string
		Arg    int
		Test2  Test2
		Test2s []Test2
	}
	ts := []Test{
		Test{
			Name:   "我的",
			Arg:    10,
			Test2:  Test2{Name: "ASD"},
			Test2s: []Test2{{Name: "A1"}, {Name: "B1"}, {Name: "B"}, {Name: "B"}, {Name: "B"}, {Name: "B"}, {Name: "B"}},
		},
		Test{
			Name:   "我的阿斯达",
			Arg:    102,
			Test2:  Test2{Name: "ASbD"},
			Test2s: []Test2{{Name: "A"}, {Name: "B"}, {Name: "B"}, {Name: "B"}, {Name: "B"}, {Name: "B"}, {Name: "B"}},
		},
	}
	e := Excel{
		FileName:  "test.xlsx",
		SheetName: "测试",
		IsHasHead: true,
		CellOptions: []CellOption{
			CellOption{Name: "name", CName: "姓名", CellType: 0},
			CellOption{Name: "arg", CName: "年龄", CellType: 1},
			CellOption{"Test2.Name", "年龄2", 2, "", 0},
			CellOption{"Test2s.Name", "姓名3", 3, "", 0},
		},
	}
	e.Write(ts)
}
func test2() {
	type Test1 struct {
		Arg int
	}
	type Test struct {
		Name  string
		Test1 Test1
		Tests []Test1
	}
	t := Test{
		Name:  "aaa",
		Test1: Test1{Arg: 44},
		Tests: []Test1{Test1{Arg: 1}, Test1{Arg: 22}, Test1{Arg: 13}},
	}
	e := Excel{
		FileName:  "test.xlsx",
		SheetName: "测试",
		IsHasHead: true,
		CellOptions: []CellOption{
			CellOption{"Test.name", "姓名", 0, "", 0},
			CellOption{"Test.Test1.Arg", "年龄", 1, "", 0},
			CellOption{"Test.Tests.Arg", "年龄1", 2, "", 0},
		},
	}
	e.Write(t)
}
func test3() {
	type Test struct {
		Name    string
		Arg     int64
		IsAdmin bool
		CreatAt time.Time
	}
	e := Excel{
		FileName:  "test.xlsx",
		SheetName: "测试",
		IsHasHead: true,
		CellOptions: []CellOption{
			CellOption{Name: "name", CName: "姓名", CellType: 0},
			CellOption{Name: "arg", CName: "年龄", CellType: 1},
			CellOption{Name: "IsAdmin", CName: "是否管理员", CellType: 2},
			CellOption{Name: "CreatAt", CName: "时间", CellType: 3, TimeFormat: 2},
		},
	}
	/*d := Test{
		Name:    "weixiao",
		Arg:     22,
		IsAdmin: true,
		CreatAt: time.Now(),
	}
	e.Write(d)*/
	resp := Test{}
	e.ReadOne(&resp)
	fmt.Println(resp)
}
func test4() {
	type Test struct {
		Name    string
		Arg     int64
		IsAdmin bool
		CreatAt time.Time
	}
	e := Excel{
		FileName:  "test.xlsx",
		SheetName: "测试",
		IsHasHead: true,
		CellOptions: []CellOption{
			CellOption{Name: "name", CName: "姓名", CellType: 0},
			CellOption{Name: "arg", CName: "年龄", CellType: 1},
			CellOption{Name: "IsAdmin", CName: "是否管理员", CellType: 2},
			CellOption{Name: "CreatAt", CName: "时间", CellType: 3, TimeFormat: 2, Type: "time.Time"},
		},
	}
	d := []Test{{
		Name:    "weixiao",
		Arg:     22,
		IsAdmin: true,
		CreatAt: time.Now(),
	}, {
		Name:    "weixiao11",
		Arg:     11,
		IsAdmin: true,
		CreatAt: time.Now(),
	}}
	e.Write(d)
	resp := []Test{}
	e.Read(&resp, Test{})
	fmt.Println(resp)
}
