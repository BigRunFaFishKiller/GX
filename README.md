# GX 为仿造GIN搭建的HTTP服务器框架 <br>
参考博客：https://geektutu.com/post/gee.html <br>

## 请求执行流程
1.当接收到请求后，调用引擎的ServeHTTP方法<br>
2.在ServeHTTP方法中，创建上下文并封装ResponseWriter，http.Request<br>
3.调用Router的Handle方法，并将上下文作为参数传入<br>
4.Handle方法解析请求路径，并在路由引擎的所有路由中寻找匹配的处理器方法并执行<br>

## 动态路由
### 将动态路由插入前缀树
1.调用路由引擎的绑定路由方法时，会调用Router的addRouter方法，此时将路径解析为一个字符串切片<br>
2.以请求方法为根节点，调用节点的insert方法，将请求路径的每一层作为前缀树的一个节点，插入前缀树<br>
3.在前缀树的最后一层的节点中，将整个请求路径赋值给节点的pattern属性<br>
4.以请求方法和请求路径的字符串为key，将处理方法放入引擎路由的handlers属性的map中<br>

### 匹配动态路由
1.调用Router的getRouter方法，查找出对应的请求路径，并查找前缀树分支的最后一个pattern不为空的节点<br>
2.取出最后一个节点的值即匹配到的路由<br>
3.将请求方法和请求路径作为key，查找引擎路由的handlers属性的map中的处理器方法，并执行<br>

## 路由分组
1.定义分组的结构体，定义分组的前缀，中间件，父分组，以及路由引擎，使引擎继承分组即路由引擎作为最顶层的分组，因此可以通过分组调用引擎统一处理请求<br>
2.创建分组时，将父分组的前缀和分组时传入的前缀拼接作为新的前缀，并将调用Group方法的分组作为父分组，在引擎中注册该分组<br>
3.添加处理器方法时，将分组的前缀和处理器方法的路径拼接<br>