package router

type RouterGroup struct {
	UserRouter
	AuthRouter
}

var NewRouterGroup = new(RouterGroup)
