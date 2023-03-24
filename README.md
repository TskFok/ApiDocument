# swagger接口文档

单独的swagger

``````
修改接口文档:
修改swagger/yml下的文件

新增接口文档:
在swagger/yml下新增文件 ****.yml
在config/conf.yaml新增对应的配置

启动:
go run main.go

访问http://127.0.0.1:8011/swagger

build见makefile

守护进程:
go run main.go bg

关闭守护进程:
kill -2 pid

进程日志:
swagger.log
``````
