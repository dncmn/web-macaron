package main

import (
	"code.dncmn.io/web-macaron/dao"
	"fmt"
	"github.com/dncmn/com"
	"os"
	"testing"
)

func TestOsRead(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(dir)
}

func TestConvert(t *testing.T) {
	str := "1234"
	v, err := com.StrTo(str).Int64()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v)
}

func TestURLEncode(t *testing.T) {
	str := "http://www.oschina.net/search?scope=bbs&q=C语言"
	rs := com.UrlEncode(str)
	fmt.Println(rs)
}

func TestGetGetDB(t *testing.T) {
	//data, isHave, err := dao.GetDeptInfoByID(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(isHave)
	//fmt.Println(data.Name)
	//
	//// 获取所有部门信息--slice
	//list, err := dao.GetAllDeptInfoSlice()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(list)
	//
	//// 获取所有的部门信息--map
	//l2, err := dao.GetAllDeptInfoMap()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(l2)
	//
	//// 测试获取所有的部门的名字
	//names, err := dao.GetAllDeptNames()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(names)
	//
	//// 获取某一个部门的信息
	//resp, err := dao.GetDeptInfoByName("体育部")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(len(resp), resp)
	//for _, v := range resp {
	//	for k, val := range v {
	//		fmt.Printf("key=%v,val=%v\n", k, val)
	//	}
	//}

	// update by map
	//cnt, err := dao.UpdateDeptInfoByDepNoByMap(1, map[string]interface{}{"addr": "this is for first", "tag": 11})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(cnt)
	//dept := model.ConfigDept{
	//	ID:   1,
	//	Tag:  1,
	//	Addr: "this is for one",
	//}
	//n, err := dao.UpdageDeptInfoByStruct(dept)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(n)

	// update aimed column
	//dept := model.ConfigDept{}
	////dept.ID = 2
	////dept.Tag = 2
	////dept.Addr = "this is for two"
	////n, err := dao.UpdateDeptInfoByFuncStruct(dept)
	////if err != nil {
	////	t.Error(err)
	////	return
	////}
	////fmt.Println(n)
	//
	//dept.Name = "计算机科学与技术"
	//dept.Addr = "http://www.baidu.com"
	//dept.Tag = 3
	//n, err := dao.InsertNewDeptToDB(dept)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(n)

	// delete info from dept
	depno := int64(5)
	n, err := dao.DeptDeleteByDepNo(depno)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(n)

}
