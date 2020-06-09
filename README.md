# 基于 gin 搭建的简单的 apiserver
## 目录结构
- admin.sh 进程的 start|stop|status|restart 控制文件
- conf 配置文件统一存放目录
- config    专门处理配置和配置文件的 go package
- db.sql    项目的 sql 文件
- api   存放消息的结构文件
- docs  swagger 文档
- handler   MVC 中的 C 层，处理请求将请求转发给实际处理的函数
- main.go   Go 程序的入口
- model    数据库相关的操作
- pkg   引用的包
  - auth    认证包
  - constvar    放置常量的位置
  - errno   错误码存放的位置
  
- router    路由
  - middleware  中间件
- service   实际业务处理
- util      工具函数
