# GX 为仿造GIN搭建的HTTP服务器框架 <br>
参考博客：https://geektutu.com/post/gee.html <br>

## 请求执行流程
1.http.ListenAndServe(addr, e)，其中第二个参数为引擎，当接收到请求后，调用引擎的ServeHTTP方法<br>
2.在ServeHTTP方法中，创建上下文并封装ResponseWriter，http.Request<br>
3.调用Router的Handle方法，并将上下文作为参数传入<br>
4.Handle方法解析请求路径，并在路由引擎的所有路由中寻找匹配的处理器方法并执行<br>