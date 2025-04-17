package router

type RouterGroup struct {
	UserRouter
	AuthRouter
	BookRouter
	CartRouter
	OrderRouter
}

var NewRouterGroup = new(RouterGroup)
