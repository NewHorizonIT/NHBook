package router

type RouterGroup struct {
	UserRouter
	AuthRouter
}

var NewRoutergroup = new(RouterGroup)
