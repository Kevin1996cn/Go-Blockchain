此项目是用Go语言实现的非常简单的区块链,用来熟悉Go语言,建立对区块链的粗浅认识.
----

项目结构:
--

core目录
-

- Block.go 用来定义区块及其属性,以及生成区块等操作。
- Blockchain.go 用来定义区块链，包括传送数据，链接区块等操作。

cmd目录
-
- main.go 用来在Terminal下测试整个系统。

rpc目录
-
- server.go 利用Go语言自带的http库，实现了一个简单的server。
通过浏览器访问接口返回Json格式的区块信息.

效果如下
-
 ![image]( picture/get.png)
 ![image]( picture/write-before.png)
 ![image]( picture/write-after.png)