# mysql basic

## mysql branch

1. mysql (oracle license)
2. mariadb (gpl free) 原创始人 https://mariadb.org/
3. Percona (free) http://www.percona.com/

4. 其他利用 mysql 原存储引擎包装的产品,aliyun, tencent.
5. TiDB (只引擎部分是 mysql)

## 差异

先熟悉 mariadb 上手.

## install

1. linux(ubuntu)
2. mac osx

   brew install mariadb

## start server

mysql.server start | stop | restart

## mysql config

server config
client config

启动参数, 常用设置

listening ip port, socket

使用管理命令工具查找服务器配置参数

```sh
mysqladmin variables  |grep -i max_conn
```

```
| max_connections   |   151
```

如何临时在线修改?
如何把修改的参数放到命令行或者配置启动文件启动初始化;

观察你修改后的效果;

## 常用 command

```sql
show processlist;
show variables;

show databases;
use $databaseName;

show tables;
desc $tableName;
select * from $tableName;
```

## privilege manager

1. user manager

   1. create a new user, modify user properties , drop a exist user

2. password set

   1. 修改用户密码
   2. flush privileges ( what meaning this? )

3. grant role on database

   1. create a new database, modify a database with special charset, drop a database.
   2. grant a user with role on database, e.t. (can [read|write|drop])

## backup | restore

1. create a backup user, only read and lock database;

```sql
# to add a backup user
GRANT LOCK TABLES, SELECT ON *.* TO 'backupuser'@'%' IDENTIFIED BY 'password';

flush privileges;
```

## database manager

1. create a new database
2. show databases list on a server
3. use a databases
4. show tables list on a database
5. copy a database
6. lock a database, unlock
7. dump database or table to a file
8. drop a database or table
9. recover the delete database or table with the dump file

### table manager

1. table
2. relation
3. primary key, uniq key,
4. index

## tips:

你每做一个命令操作的时候要思考下,这个命令是不是在修改数据,如果做错,要如何恢复.再整理下以后操作的流程,以后按流程步骤操作.

## 扩展

1. 如何设置一个主从服务器
2. 如何观察服务器运行状况

3. 如何监视 slow search operator
