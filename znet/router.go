package znet

import "zinx/ziface"

//实现router时，先嵌入这个BaseRoute基类，然后根据需要去重写
type BaseRouter struct {}


//这里之所以BaseRouter的方法都为空
//在处理conn业务之前的钩子方法
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//处理conn业务的方法
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//处理conn业务之后的狗子方法
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}

