package controller

import (
	"code.dncmn.io/web-macaron/utils/logging"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/macaron.v1"
	"net/http"
)

var (
	logger = logging.GetLogger()
)

// SendJson:response json body
func SendJson(c *macaron.Context, v interface{}) {
	cnt, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Infof("response body.url[%s],body=[%v]\n", c.Req.RequestURI, string(cnt))
	c.JSON(http.StatusOK, v)
	return
}

// BindXml: parse requestBody into resp by xml
func BindXml(c *macaron.Context, resp interface{}) (err error) {
	var (
		cntBytes []byte
	)
	logger.Infof("start parse param url[%s] into [%v]\n", c.Req.URL.String(), resp)
	cntBytes, err = c.Req.Body().Bytes()
	if err != nil {
		logger.Infof("read content from url[%v],err[]%v\n", c.Req.RequestURI, err)
		return
	}
	err = xml.Unmarshal(cntBytes, &resp)
	if err != nil {
		logger.Infof("[Bind Xml error: url=%s,errInfo=%s]\n", c.Req.RequestURI, err.Error())
		return
	}
	return
}

// BindJson: parse requestBody into resp by json
func BindJson(c *macaron.Context, resp interface{}) (err error) {
	var (
		cntBytes []byte
	)
	logger.Infof("start parse param url[%s] into [%v]\n", c.Req.RequestURI, resp)
	cntBytes, err = c.Req.Body().Bytes()
	if err != nil {
		logger.Infof("read content from url[%v],err[]%v\n", c.Req.RequestURI, err)
		return
	}
	err = json.Unmarshal(cntBytes, &resp)
	if err != nil {
		logger.Infof("[Bind Json error: url=[%s] errorInfo[%s]]\n", c.Req.RequestURI, err.Error())
	}
	return
}
