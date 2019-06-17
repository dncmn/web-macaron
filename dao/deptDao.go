package dao

import "code.dncmn.io/web-macaron/model"

// GetDeptInfoByID:根据部门id获取部门信息
func GetDeptInfoByID(id uint64) (data model.ConfigDept, isExist bool, err error) {
	isExist, err = engine.Where("depno=?", id).Get(&data)
	return
}

// GetDeptInfoByName:根据部门名称获取部门信息
func GetDeptInfoByName(depName string) (data []map[string]string, err error) {
	data = make([]map[string]string, 0)
	data, err = engine.SQL("select * from dept where depname=?", depName).QueryString()
	return
}

// GetAllDeptInfoSlice:获取所有部门信息--slice
func GetAllDeptInfoSlice() (data []model.ConfigDept, err error) {
	err = engine.Asc("depno").Find(&data)
	return
}

// GetAllDeptInfoMap:获取所有的部门信息-map
func GetAllDeptInfoMap() (data map[int64]model.ConfigDept, err error) {
	data = make(map[int64]model.ConfigDept)
	err = engine.Find(&data)
	return
}

// GetAllDeptNames():获取所有的部门的名字
func GetAllDeptNames() (names []string, err error) {
	err = engine.Table(model.ConfigDept{}.TableName()).Asc("depno").Cols("depname").Find(&names)
	return
}

// UpdateDeptInfoByDepNoByMap:更改部门信息--map
func UpdateDeptInfoByDepNoByMap(depNo uint64, updateMap map[string]interface{}) (affectedNum int64, err error) {
	affectedNum, err = engine.Table(&model.ConfigDept{}).Where("depno=?", depNo).Update(&updateMap)
	return
}

// UpdageDeptInfoByStruct:更改部门信息-struct  默认只更新非空和非0字段
func UpdageDeptInfoByStruct(dept model.ConfigDept) (num int64, err error) {
	num, err = engine.Where("depno=?", dept.ID).Update(&dept)
	return
}

// UpdateDeptInfoByFuncStruct:更改部门信息--更改指定列信息(指定更改函数)
func UpdateDeptInfoByFuncStruct(dept model.ConfigDept) (cnt int64, err error) {
	cnt, err = engine.Id(dept.ID).Cols("tag,addr").Update(&dept)
	return
}

// InsertNewDeptToDB:添加一条或者多条记录到部门表
func InsertNewDeptToDB(dept model.ConfigDept) (n int64, err error) {
	n, err = engine.Insert(&dept)
	return
}

// DeptDeleteByDepNo:删除部门信息
func DeptDeleteByDepNo(depNo int64) (n int64, err error) {
	n, err = engine.Where("depno=?", depNo).Delete(&model.ConfigDept{})
	return
}
