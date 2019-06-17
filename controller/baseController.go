package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/macaron.v1"
	"net/http"
)

// SendJson:response json body
func SendJson(c *macaron.Context, v interface{}) {
	cnt, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("response body.url[%s],body=[%v]", c.Req.RequestURI, string(cnt))
	c.JSON(http.StatusOK, v)
	return
}

// BindXml: parse requestBody into resp by xml
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

// BindJson: parse requestBody into resp by json
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
