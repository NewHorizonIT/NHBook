package router

type RouterGroup struct {
	UserRouter
	AuthRouter
	BookRouter
	CartRouter
}

var NewRouterGroup = new(RouterGroup)
