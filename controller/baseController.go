package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/macaron.v1"
)

func BindXml(c *macaron.Context, resp interface{}) (err error) {
	var (
		cntBytes []byte
	)
	fmt.Printf("start parse param url[%s] into [%v]\n", c.Req.RequestURI, resp)
	cntBytes, err = c.Req.Body().Bytes()
	if err != nil {
		fmt.Printf("read content from url[%v],err[]%v\n", c.Req.RequestURI, err)
		return
	}
	err = xml.Unmarshal(cntBytes, &resp)
	if err != nil {
		fmt.Printf("[Bind Xml error: url=%s,errInfo=%s]\n", c.Req.RequestURI, err.Error())
		return
	}
	return
}

func BindJson(c *macaron.Context, resp interface{}) (err error) {
	var (
		cntBytes []byte
	)
	fmt.Printf("start parse param url[%s] into [%v]\n", c.Req.RequestURI, resp)
	cntBytes, err = c.Req.Body().Bytes()
	if err != nil {
		fmt.Printf("read content from url[%v],err[]%v\n", c.Req.RequestURI, err)
		return
	}
	err = json.Unmarshal(cntBytes, &resp)
	if err != nil {
		fmt.Printf("[Bind Json error: url=[%s] errorInfo[%s]]\n", c.Req.RequestURI, err.Error())
	}
	return
}
