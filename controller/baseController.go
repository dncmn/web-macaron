package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
)

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
