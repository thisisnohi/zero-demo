
## 脚本
```sql
drop table mike_user;
create table MIKE_USER(
id int primary key ,
mobile varchar(40) not null,
nickname varchar(100) not null
);
create unique index MIKE_USER_MOBILE_UK on MIKE_USER (mobile);

drop table mike_user_data;
create table mike_user_data(
id int primary key ,
user_id int not null,
data varchar(100),
create_time datetime not null default current_timestamp,
update_time datetime not null default current_timestamp
);
```

### sql2pb
* cd pb
* sql2pb -go_package ./pb -host localhost -port 3306  -package pb -user root -password 数据库用户密码 -schema go_zero -service_name ordercenter > ordercenter.prot
