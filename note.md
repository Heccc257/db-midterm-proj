使用方法
* 可以根据根目录中`pku_mutualhelper.sql`中的DDL来构建数据库
* 更推荐使用调用`./backend/server`。该可执行文件可以打开一个http服务，连接数据库(`mysql -h127.0.0.1 -uroot -p123456`)并且初始化整张表。
    * 如果安装了`go`，建议先在backend文件夹下执行`go build`
    * 调用了`backend/server`后，可以新开后台连接数据库进行后续操作
    * 在`testbench`文件夹下写好了一些测试用例，这些测试用例发送http请求到本机，



### 数据库特性的使用

#### 主键

* 每个表中都有主键

#### 外键

xxx

#### 索引

* offer_category中设置了category_name为unique key，保证不会有两个category重名

#### 视图

* 添加了`user`表的视图`user_info`，隐去了`password_hash`字段，提高了查询和修改的安全性。

#### 触发器

* 上传offer时候，如果没有填入time_limit或者reward_amout，则插入时自动的根据offer_category中的推荐值填入

  * ```sql
    CREATE TRIGGER set_default_offer BEFORE INSERT ON `offer`
    	FOR EACH ROW
    	BEGIN
    
    		DECLARE reward FLOAT;
    		DECLARE time_limit_min int;
    		SELECT adviced_reward, adviced_timelimit_min INTO reward, time_limit_min FROM offer_category WHERE category_name = NEW.category_name;
    
    		IF NEW.reward_amount < 1 THEN
    			SET NEW.reward_amount = reward;
    		END IF;
    		IF NEW.time_limit < NOW() THEN
    			SET NEW.time_limit = DATE_ADD(NOW(), INTERVAL time_limit_min MINUTE);
      		END IF;
    	END;
    ```

#### 复杂sql

* 将用户按照所有rating的平均值排序，返回最高的10个用户的id，nick_name等信息。

  * ```sql
    SELECT u.user_id, u.nick_name, AVG(ur.rating) as average_rating
    FROM user u
    JOIN user_rating ur ON u.user_id = ur.rated_user_id
    WHERE u.task_count > 1
    GROUP BY u.user_id, u.nick_name
    HAVING COUNT(ur.rating_id) > 0
    ORDER BY average_rating DESC;
    ```

#### 存储过程

* 将根据uid查询user_info添加为存储过程

  * ```sql
    CREATE PROCEDURE GetUserByID(IN uid INT)
    BEGIN
    	SELECT *
    	FROM user_info
    	WHERE user_id = uid;
    END
    ```

#### 事务 

* 上传offer的时候同时修改`user`表中的`offer_count`
* 接受offer的时候同时修改`user`表中的`task_count`
* 完成订单的时候必须同时查询对应的接单记录，失败时回滚。





## front-end

* 注册，登陆界面
    * 登陆后返回uid以及鉴权token
* 左侧项目
    * 个人主页
        * 返回user_info
    * 订单
        * 显示未被接取的订单。用户可以在这个界面接取订单 `accept_offer`
    * 我的发单
        * 显示用户已经发送的订单
        * 有发单按钮，可以发送订单 `offer_post`
        * 对于已经完成的订单`offer_state == completed`,可以提交评论 `user_rating`
    * 我的接单
        * 显示我已经接取了的订单
        * 有`完成订单`的按钮
    * 用户排名





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

选出`task_cout` > 1 (也就是说至少完成过一个订单)的用户，将他们按照rating的平均值排序 (这个sql够复杂了)

```sql
SELECT u.user_id, u.nick_name, AVG(ur.rating) as average_rating
FROM user u
JOIN user_rating ur ON u.user_id = ur.rated_user_id
WHERE u.task_count > 1
GROUP BY u.user_id, u.nick_name
HAVING COUNT(ur.rating_id) > 0
ORDER BY average_rating DESC;

```



