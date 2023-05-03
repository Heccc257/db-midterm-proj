create table Users
(
   user_id              int not null,
   nickname             varchar(20) not null,
   full_name            varchar(20) not null,
   phone_number         varchar(20) not null,
   password_hash        char(64) not null,
   order_count          int not null default 0,
   order_complaint_count int not null default 0,
   task_count           int not null default 0,
   task_complaint_count int not null default 0,
   primary key (user_id),
   key AK_u_phone (phone_number)
);

create table complaint
(
   complaint_id         int not null,
   complainant_id       int not null,
   defendant_id         int not null,
   order_id             int not null,
   complaint_reason     text not null,
   time                 timestamp not null,
   primary key (complaint_id)
);

create table location
(
   location_id          int not null,
   building        VARCHAR(30) not null,
   floor VARCHAR(30),
   primary key (location_id)
);

/*
这里不需要id,name就可以当主键了
*/
create table order_category
(
   category_name        varchar(50) not null,
   adviced_reward decimal(10,2) DEFAULT 100,
   adviced_timelimit_min int DEFAULT 100,
    decimal(10,2) DEFAULT 100,
   primary key (category_name)
);


create table `order`
(
   order_id             int not null,
   reward_amount        decimal(10,2) not null,
   customer_id          int not null,
   category          varchar(50) not null,
   pickup_location_id   int not null,
   delivery_location_id int not null,
   time_limit           Datetime,
   primary key (order_id)
);


/*
同一个用户如果多次接相同的单则修改原来的单
*/
create table accept_order
(
   order_id             int not null,
   user_id              int not null,
   created_at timestamp NULL DEFAULT NULL,
   status_code          ENUM('canceled', 'in_progress', 'completed') not null default 'in_progress',
   primary key (order_id, user_id)
)

create table user_rating
(
   rating_id            int not null,
   rater_id             int not null,
   rated_user_id        int not null,
   order_id             int not null,
   rating               int not null,
   comment              text,
   primary key (rating_id)
);
