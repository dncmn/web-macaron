package service

import (
	"code.dncmn.io/web-macaron/dao"
	"code.dncmn.io/web-macaron/model"
	"fmt"
)

type AddNewDeptReq struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Tag  int    `json:"tag"`
}

type DeptListResp struct {
	No   uint64 `json:"deptno"` // 部门编号
	Name string `json:"name"`   // 部门名称
	Addr string `json:"addr"`   // 部门地址
}

func GetDeptList() (resp []DeptListResp, err error) {
	var (
		list []model.ConfigDept
	)
	list, err = dao.GetAllDeptInfoSlice()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		r := DeptListResp{
			No:   v.ID,
			Name: v.Name,
			Addr: v.Addr,
		}
		resp = append(resp, r)
	}
	return
}

//
func AddNewDept(body AddNewDeptReq) (err error) {
	var (
		affectLineNum int64
	)
	dbDept := model.ConfigDept{
		Name: body.Name,
		Addr: body.Addr,
		Tag:  body.Tag,
	}
	affectLineNum, err = dao.InsertNewDeptToDB(dbDept)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("new add dept:no=%v,name=%s,afftectedLineNo=%d", dbDept.ID, dbDept.Name, affectLineNum)
	return
}
