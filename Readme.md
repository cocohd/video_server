> handler -> validation(1request 2user) -> business logic -> response
>> validation:
>> 数据结构：data；error 

> 数据库设计 
>> 见DBFile 

> dbops/api_test.go是测试
> comments只涉及删和查

> session 记录用户在服务器的状态
> session和cookie：session实在服务端，cookie是客户端的。有时会将session存放在cookie里面。

> 路由中间件
- main->middleware->defs(message, err)->handlers->dbops->response
