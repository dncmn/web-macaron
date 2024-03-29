package main

import (
	"code.dncmn.io/web-macaron/controller"
	_ "code.dncmn.io/web-macaron/dao"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func() string {
		return "helloWorld"
	})

	m.Group("/user", func() {
		m.Get("/username", controller.GetUserInfoHandler)
		m.Get("/userByUID/:uid", controller.GetUserByUIDHandler)
		m.Post("", controller.UserRegisterHancler)
	})

	m.Group("/dept", func() {
		m.Get("/list", controller.GetDeptListHancler)
		m.Post("/addNewDept", controller.InsertNewDeptHandler)
	})

	m.Run("localhost", 9002)
}
