## 一. interface server

解析：想要创建`http`服务，就需要实现`http.Handler`接口
以及是实现启动服务方法`Start`和停止服务方法`Stop`，
并且给开发者提供自定义接口`addRouter`添加路由。

```go
/* ========= 接口定义 ========= */
type server interface {
/*
	## Means
		make goland's http package intake server's interface_defined
	## Server's function?
		* Start
		* Stop
	## Why we need to abstract the http server?
		* We can use the http server to implement the server interface_defined
		* We need compatible with the http server
*/
http.Handler
// Start : 开启服务
Start(addr string) error
// Stop : 关闭服务
Stop() error
// addRouter : 注册路由 addRouter(请求方法，请求路径，视图函数)
/*
	视图函数是用来处理请求的函数，它的参数是http.ResponseWriter和*http.Request
*/
addRouter(method string, pattern string, handleFunc HandleFunc)
}
```

[interface_server.go](../server/interface_server.go)

### 1. `http.Handler`

    硬性要求

### 2. `Start`

    启动服务

&emsp;对于`Start`方法而言，
我们需要返回`http.ListenAndServer(addr, h)`，
其中`addr`是服务地址，`h`是`http.Handler`接口。

### 3. `Stop`

    停止服务

&emsp;对于`Stop`方法而言，
就是思考如何停止，当然可以直接停止，
但是也可以优雅停止。
对于优雅关闭而言，因为对于不同开发者而言，
其可能需要不同的优雅关闭方式，所以提供了`HTTPOption`模式，
主要体现在“选择性”上,如果开发者需要优雅关闭，
那么只需要开发者传入`HTTPOption`即可，反之，
就使用默认的关闭方式。

### 4. `addRouter`

    添加路由

&emsp;在服务启动之后，现在可以停止服务，也可以启动服务，
接下来就是处理请求，想要处理请求，就是存储路由`addRouter`了，

```go
    addRouter(method string, pattern string, handleFunc HandleFunc)
```

`addRouter`方法的参数：

* method : 请求方法
* pattern : 请求路径
* handleFunc : 视图函数

> 注册路由和启动路由的顺序:
> * 先注册路由
> * 之后启路由

> 路由和视图的对应关系
> * key-value
> * 路由的组成：
> * key:请求方法+请求路径 eg:GET-/hello
> * value:视图函数 eg:HandleFunc()

> 路由的注册(也可以称为存储)
> * 在`HTTPServer`结构体之中注册路由
> * 需要保证`key`的唯一性
> * 存储方式：map[string]HandleFunc

## 二. struct server

```go
/* ========= 结构体对象定义 ========= */
// HTTPServer : This is a http server.
type HTTPServer struct {
srv *http.Server
// 优雅关闭服务
stop func () error
/* routers 临时存放路由信息
   {
   	"GET /": func(w http.ResponseWriter, r *http.Request) { // 处理请求 },
   	"POST /": func(w http.ResponseWriter, r *http.Request) { // 处理请求 },
   }
*/
routers map[string]HandleFunc
}
```

[struct_httpserver.go](../server/struct_httpserver.go)

### 1. `srv`

    服务对象

> `server`层需要存在一个对外的接口，
> 此处`srv`就是对外的接口，
> 开发者可以通过`srv`来访问`server`层的方法，
> 对于`HTTPServer`结构体而言，我们创建`ServerHTTP`方法，
> 该方法的作用就是将`HTTPServer`结构体转换成`http.Handler`接口，
> 这样就可以使用`http.ListenAndServe`方法来启动服务了。

&emsp;对于`srv`而言，
就是`http.Server`对象，
主要是为了方便开发者使用`http.Server`对象的方法，
例如：`ListenAndServe`、`Shutdown`等。

### 2. `stop`

    优雅关闭服务

&emsp;对于`stop`而言，
就是优雅关闭服务的函数，
主要是为了方便开发者使用`http.Server`对象的方法，
例如：`Shutdown`等。

### 3. `routers`

    路由存储

&emsp;对于`routers`而言，
就是路由存储，
主要是为了方便开发者使用`http.Server`对象的方法，
例如：`ServeHTTP`等。

> 匹配路由：
> 如果存在匹配的路由，那么返回其视图函数
> 反之，返回nil

## 三. 视图函数

[struct_context.go](../context/struct_context.go)

> 请求上下文

### 1. 为什么重写视图函数？

> 扩展性

&emsp;对于传统视图函数而言，
传入的参数有限，
分别是:`http.ResponseWriter`和`*http.Request`，
现在重写视图函数，
在其中加上请求路径和请求方法，
也是为了优雅起见。

### 2. 为了迎合上下文

&emsp;上下文贯穿整个请求和响应过程，
因此，我们需要在视图函数中加入上下文，
比如说，对于事件A而言，
它在对`Context`中的某些数据进行修改之后，
对于事件`B`而言，它需要使用`A`修改之后的，
也就是`Context`中的某些数据，
就是通过上下文来实现的。

## 四. 总结

* 先匹配路由

* 然后对于每一个路由

* 存在一个上下文管理器

* 上下文管理器中存在一个请求和响应的上下文处理事件

&emsp;不可能存在多个视图函数，共用一个上下文处理事件的情况，
上下文处理事件和视图函数是一对一的关系。