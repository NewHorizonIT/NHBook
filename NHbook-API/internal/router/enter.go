package router

type RouterGroup struct {
	UserRouter
	AuthRouter
	BookRouter
	CartRouter
	OrderRouter
	CategoryRouter
}

var NewRouterGroup = new(RouterGroup)
