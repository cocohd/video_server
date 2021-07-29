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


*****
> **stream**
> 
> limit: 防止请求过多，造成server链接数消耗完,ram消耗完就会crash。
> 流控制算法实现：如使用一个bucket：有20个token，request进来使用一个token，repsonse后返回占用的token；以此实现流控。
>

****
> **scheduler**
> 实现异步操作(诸如延迟删除等)
> 
> 经由一个channel实现生产者/消费者的通信