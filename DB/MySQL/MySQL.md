# MySQL

​	执行流程

<img src="https://static001.geekbang.org/resource/image/0d/d9/0d2070e8f84c4801adbfa03bda1f98d9.png" style="zoom: 25%;" />

## 存储引擎

常用：InnoDB，MylSAM

## Server层

Server 层包括连接器、查询缓存、分析器、优化器、执行器等，涵盖 MySQL 的大多数核心服务功能，以及所有的内置函数（如日期、时间、数学和加密函数等），所有跨存储引擎的功能都在这一层实现，比如存储过程、触发器、视图

## 查询语句

1. 连接器：客户端如何连接到数据库（mysql -uroot -proot）
2. 查询缓存：在缓存中查看是否之前执行过这条语句，缓存艺K-V形式存在，key查询语句，value查询结果
3. 分析器：识别输入的关键字（select）
4. 优化器：一条 SQL 语句可能有不同的执行逻辑（或者顺执行顺序），而优化器就是选择最优的执行顺序。
5. 执行器：真正执行sql语句



## 更新语句

### 日志模块

 	1. Redo Log（重做日志）
      	1. InnoDB引擎先把记录写到redo log 中，redo log 在哪，他也是在磁盘上，这也是一个写磁盘的过程，但是与更新过程不一样的是，更新过程是在磁盘上随机IO，费时。 而写redo log 是在磁盘上顺序IO。效率要高。*InnoDB的redo log 是固定大小，结构类似于循环队列*
      	2. WAL（Write-Ahead Logging）：先写日志在写磁盘
 	2. BinLog（归档日志）
      	1. Server层的日志
      	2. binlog 是逻辑日志，记录的是这个语句的原始逻辑，比如“给 ID=2 这一行的 c 字段加 1 ”
 	3. 差异
      	1. redo log 是 InnoDB 引擎特有的；binlog 是 MySQL 的 Server 层实现的，所有引擎都可以使用。
      	2. redo log 是物理日志，记录的是“在某个数据页上做了什么修改”；binlog 是逻辑日志，记录的是这个语句的原始逻辑，比如“给 ID=2 这一行的 c 字段加 1 ”。
      	3. redo log 是循环写的，空间固定会用完；binlog 是可以追加写入的。“追加写”是指 binlog 文件写到一定大小后会切换到下一个，并不会覆盖以前的日志。

<img src="https://static001.geekbang.org/resource/image/2e/be/2e5bff4910ec189fe1ee6e2ecc7b4bbe.png" style="zoom:25%;" />**两段提交**







