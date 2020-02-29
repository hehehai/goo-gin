go 项目目录的一些文件及文件及的介绍
api go版本的api文件
bin 存放标准的命令文件，也就是可执行文件
blog 官方的blog文档
doc go 标准库的介绍文档，我们可以启动一个 godoc 在本地生成一个标准库文档服务
lib 存放一些特殊的库文件
misc 存放一些辅助类的说明和工具
pkg 存放安装go标准库后的归档文件，以.a 结尾的文件
src 是存放go项目的，也就是存放源码文件的
test 是测试相关的文件


indirect 间接的，非直接的，在我们的依赖包被按照后，依赖包的依赖也会安装，那么对于我们当前的模块，依赖包就分为两种
一类是直接依赖，也就是程序中有使用的代码，另一类是简介依赖，程序中没有使用

我们的目录结构

```
├── README.md 
├── conf 用来存放配置文件 
├── go.mod
├── go.sum
├── main.go
├── middleware 应用的中间件
├── models 应用数据库模型
├── pkg 第三方使用的依赖包
├── routers 路由逻辑处理
└── runtime 应用运行时的数据
```

go mod 中有5种动词
module 用来定义当前模块的路径
go 用于设置预期的go的版本
require 依赖的模块版本信息
exclude 从使用中排除一个特定的模块依赖版本
replace 将一个模块替换为另一个模块版本