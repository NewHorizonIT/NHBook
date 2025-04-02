package router

type RouterGroup struct {
	UserRouter
	AuthRouter
	BookRouter
}

var NewRouterGroup = new(RouterGroup)
