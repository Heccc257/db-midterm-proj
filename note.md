使用方法
* 可以根据根目录中`pku_mutualhelper.sql`中的DDL来构建数据库
* 更推荐使用调用`./backend/server`。该可执行文件可以打开一个http服务，连接数据库(`mysql -h127.0.0.1 -uroot -p123456`)并且初始化整张表。
    * 如果安装了`go`，建议先在backend文件夹下执行`go build`
    * 调用了`backend/server`后，可以新开后台连接数据库进行后续操作
    * 在`testbench`文件夹下写好了一些测试用例，这些测试用例发送http请求到本机，


## TODO

#### 创建存储过程

根据用户的id找到对应的User的行，以及根据用户的phone_number找到对应行
```sql
// code here
DELIMITER //
CREATE PROCEDURE GetUserByID(IN UserID INT)
BEGIN
    SELECT * FROM user WHERE user_id = UserID;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE GetUserByPhoneNumber(IN PhoneNumber VARCHAR(20))
BEGIN
    SELECT * FROM user WHERE phone_number = PhoneNumber;
END //
DELIMITER ;



```

#### 根据user_rating排序
选出task_cout > 1 (也就是说至少完成过一个订单)的用户，将他们按照rating的平均值排序 (这个sql够复杂了)
```sql

SELECT u.user_id, u.nick_name, AVG(ur.rating) as average_rating
FROM user u
JOIN user_rating ur ON u.user_id = ur.rated_user_id
WHERE u.task_count > 1
GROUP BY u.user_id, u.nick_name
HAVING COUNT(ur.rating_id) > 0
ORDER BY average_rating DESC;

```