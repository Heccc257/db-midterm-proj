/*
暂时先不把外键写在里面
*/
create table user
(
   user_id              int unsigned AUTO_INCREMENT not null,
   nick_name             varchar(20) not null,
   full_name            varchar(20) not null,
   phone_number         varchar(20) not null,
   password_hash        char(64) not null,
   offer_count          int not null default 0,
   -- offer_complaint_count int not null default 0,
   task_count           int not null default 0,
   -- task_complaint_count int not null default 0,
   primary key (user_id)
);


create table complaint
(
   complaint_id         int unsigned AUTO_INCREMENT not null,
   complainant_id       int unsigned not null,
   defendant_id         int unsigned not null,
   offer_id             int unsigned not null,
   complaint_reason     text not null,
   time                 timestamp not null,
   primary key (complaint_id)
);

create table location
(
   location_id          int unsigned AUTO_INCREMENT not null,
   building        VARCHAR(30) not null,
   floor 			VARCHAR(30),
   primary key (location_id)
);

/*
这里不需要id,name就可以当主键了
*/
create table offer_category
(
   category_id    int unsigned AUTO_INCREMENT,
   category_name        varchar(50) not null,
   adviced_reward decimal(10,2) DEFAULT 100,
   adviced_timelimit_min int DEFAULT 100,
   unique key(category_name),
   primary key (category_id)
);


create table offer
(
   offer_id             int unsigned not null AUTO_INCREMENT,
   reward_amount        decimal(10,2) not null,
   customer_id          int unsigned not null,
   category_name        varchar(50) not null,
   pickup_location_id   int unsigned not null,
   delivery_location_id int unsigned not null,
   offer_state          ENUM('pending', 'in_progress', 'completed') not null default 'pending',
	created_at           timestamp NULL DEFAULT NULL,
   time_limit           Datetime,
   primary key (offer_id)
);
alter table offer ADD CONSTRAINT FK_1 foreign key(category_name) references offer_category(category_name) on delete  CASCADE;


/*
同一个用户如果多次接相同的单则修改原来的单
*/
create table accept_offer
(
   offer_id             int unsigned not null,
   user_id              int unsigned  not null,
   created_at timestamp NULL DEFAULT NULL,
   complete_time Datetime,
   task_state          ENUM('in_progress', 'completed') not null default 'in_progress',
   primary key (offer_id, user_id)
);


create table user_rating
(
   rating_id            int unsigned AUTO_INCREMENT not null,
   rater_id             int unsigned not null,
   rated_user_id        int unsigned not null,
   offer_id             int unsigned not null,
   rating               int not null,
   comment              text,
   primary key (rating_id)
);


alter table accept_offer add constraint FK_Reference_15 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;

alter table accept_offer add constraint FK_Reference_16 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;

alter table complaint add constraint FK_Reference_6 foreign key (complainant_id)
      references user (user_id) on delete restrict on update restrict;

alter table complaint add constraint FK_Reference_7 foreign key (defendant_id)
      references user (user_id) on delete restrict on update restrict;

alter table complaint add constraint FK_Reference_8 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;

alter table offer add constraint FK_Reference_1 foreign key (customer_id)
      references user (user_id) on delete restrict on update restrict;

alter table offer add constraint FK_Reference_19 foreign key (category_name)
      references offer_category (category_name) on delete restrict on update restrict;

alter table offer add constraint FK_Reference_4 foreign key (pickup_location_id)
      references location (location_id) on delete restrict on update restrict;

alter table offer add constraint FK_Reference_5 foreign key (delivery_location_id)
      references location (location_id) on delete restrict on update restrict;

alter table user_rating add constraint FK_Reference_10 foreign key (rater_id)
      references user (user_id) on delete restrict on update restrict;

alter table user_rating add constraint FK_Reference_11 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;

alter table user_rating add constraint FK_Reference_9 foreign key (rated_user_id)
      references user (user_id) on delete restrict on update restrict;

