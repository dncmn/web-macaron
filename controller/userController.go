package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
)

type UserRegisterJsonReq struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserRegisterXMLReq struct {
	Name     string `xml:"name"`
	Age      int    `xml:"age"`
	Password string `xml:"password"`
}

// UserRegisterHancler:用户注册
func UserRegisterHancler(c *macaron.Context) {
	var (
		err  error
		req  UserRegisterJsonReq
		resp interface{}
	)

	//defer SendJson(c, resp)
	if err = BindXml(c, &req); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("username=%v,age=%v,all=%v\n", req.Name, req.Age, req)
	resp = map[string]interface{}{
		"age": 2019,
	}
	SendJson(c, resp)
	return
}

func GetUserInfoHandler(c *macaron.Context) {
	var (
		err error
		sex int64
	)

	sex = c.QueryInt64("sex")

	fmt.Printf("request addr=%v,sex=%v\n", c.Req.RemoteAddr, sex)
	_, err = c.Resp.Write([]byte(c.Req.RemoteAddr))
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// GetUserByUIDHandler:g根据uid获取用户信息
func GetUserByUIDHandler(c *macaron.Context) {
	var (
		uid string
		err error
	)
	uid = c.Params("uid")
	if len(strings.TrimSpace(uid)) == 0 {
		err = errors.New("param error")
		fmt.Println(err)
		return
	}
	fmt.Printf("uid=%v\n", uid)
	resp := map[string]interface{}{
		"uid": uid,
	}

	cnt, _ := json.Marshal(resp)
	c.Resp.WriteHeader(http.StatusOK)
	c.Resp.Write([]byte(cnt))
	return
}
