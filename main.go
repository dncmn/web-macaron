package main

import (
	"code.dncmn.io/web-macaron/controller"
	"code.dncmn.io/web-macaron/router"
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
	m.Group("/c", router.UserRouter)
	m.Run("localhost", 9002)
}
