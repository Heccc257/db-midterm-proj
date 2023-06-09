#### 后端程序使用方法
* 云服务器测试（推荐）
    * 已在阿里云上部署后端程序，启动前端程序可直接连接云服务器
* 本地测试
    * 在`./backend`目录下执行`go build && ./server`指令，会编译并运行后端程序。
    * 需要`go1.18.1`及以上版本
    * 服务端默认监听`:9999`，确保本机该端口未被占用。也可通过`./server --port XXX`的方式指定端口
    * 服务端会连接`mysql`，连接dsn为`mysql -h127.0.0.1 -uroot -p123456`。
        * 需要确保root账户的身份验证方式(authentication plugin)为`mysql_native_password`。
    * 打开后端后会默认清空数据库

#### 测试样例
* 文件存放在`test_bench`目录下，使用curl命令发送http请求到本地。
* 执行`sh first_execute.sh`命令会初始化数据库中的内容，包括注册用户、发出offer等。
    * **必须在数据库中内容为空时调用**，否则会有编号不匹配问题。
* 执行`sh restart_db.sh`命令会清空数据库。

