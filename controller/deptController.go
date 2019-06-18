package controller

import (
	"code.dncmn.io/web-macaron/service"
	"fmt"
	"gopkg.in/macaron.v1"
)

// GetDeptListHancler:获取部门列表
func GetDeptListHancler(c *macaron.Context) {
	var (
		err  error
		resp []service.DeptListResp
	)

	if resp, err = service.GetDeptList(); err != nil {
		SendJson(c, nil)
		fmt.Println(err)
		return
	}
	SendJson(c, resp)
	return
}

// 创建新部门
func InsertNewDeptHandler(c *macaron.Context) {
	var (
		body service.AddNewDeptReq
		err  error
	)

	if err = BindJson(c, &body); err != nil {
		fmt.Println(err)
		SendJson(c, nil)
		return
	}
	logger.Infof("name=%s,addr=%s,body=%v", body.Name, body.Addr, body)

	if err = service.AddNewDept(body); err != nil {
		fmt.Println(err)
		SendJson(c, nil)
		return
	}
	SendJson(c, map[string]interface{}{"status": "create success"})
	return
}
