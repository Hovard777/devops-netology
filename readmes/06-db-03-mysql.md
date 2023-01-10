1. Init Mysql docker
```commandline
[ifebres@febres-i 06-db-03-mysql]$ sudo docker create volume vol1
[ifebres@febres-i 06-db-03-mysql]$ sudo docker pull mysql:8.0
[ifebres@febres-i 06-db-03-mysql]$ sudo docker run --rm --name mysql-docker -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -ti -p 3306:3306 -v vol1:/etc/mysql/  mysql:8.0
[ifebres@febres-i 06-db-03-mysql]$ sudo docker exec -it mysql-docker bash
bash-4.4# ls /etc/mysql/
conf.d	test_dump.sql
bash-4.4# mysql -e 'create database test_db;'
bash-4.4# mysql test-db < /etc/mysql/test_dump.sql 
ERROR 1049 (42000): Unknown database 'test-db'
bash-4.4# mysql test_db < /etc/mysql/test_dump.sql 
bash-4.4# \s
bash: s: command not found
bash-4.4# mysql
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 11
Server version: 8.0.31 MySQL Community Server - GPL

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> \s
--------------
mysql  Ver 8.0.31 for Linux on x86_64 (MySQL Community Server - GPL)

Connection id:		11
Current database:	
Current user:		root@localhost
SSL:			Not in use
Current pager:		stdout
Using outfile:		''
Using delimiter:	;
Server version:		8.0.31 MySQL Community Server - GPL
Protocol version:	10
Connection:		Localhost via UNIX socket
Server characterset:	utf8mb4
Db     characterset:	utf8mb4
Client characterset:	latin1
Conn.  characterset:	latin1
UNIX socket:		/var/run/mysqld/mysqld.sock
Binary data as:		Hexadecimal
Uptime:			1 min 48 sec

Threads: 2  Questions: 38  Slow queries: 0  Opens: 459  Flush tables: 3  Open tables: 47  Queries per second avg: 0.351
--------------
```
Database version: 8.0.31 MySQL Community Server - GPL
Dabtabase Data
```commandline
mysql> use test_db;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+-------------------+
| Tables_in_test_db |
+-------------------+
| orders            |
+-------------------+
1 row in set (0.00 sec)

mysql> select count(*) from orders where price >300;
+----------+
| count(*) |
+----------+
|        1 |
+----------+
1 row in set (0.00 sec)
```
2. Ex2
```commandline
mysql> CREATE USER 'test' IDENTIFIED WITH mysql_native_password BY 'testpass'
    -> WITH MAX_QUERIES_PER_HOUR 100 PASSWORD EXPIRE INTERVAL 180 DAY FAILED_LOGIN_ATTEMPTS 3
    -> ATTRIBUTE '{"surname": "Pretty", "name": "James"}';
Query OK, 0 rows affected (0.01 sec)

mysql> GRANT SELECT on test_db.* TO test;
Query OK, 0 rows affected (0.02 sec)

mysql> select * from INFORMATION_SCHEMA.USER_ATTRIBUTES where user="test";
+------+------+----------------------------------------+
| USER | HOST | ATTRIBUTE                              |
+------+------+----------------------------------------+
| test | %    | {"name": "James", "surname": "Pretty"} |
+------+------+----------------------------------------+
1 row in set (0.00 sec)
```
3. Ex3
```commandline
mysql> show table status\G
*************************** 1. row ***************************
           Name: orders
         Engine: InnoDB
        Version: 10
     Row_format: Dynamic
           Rows: 5
 Avg_row_length: 3276
    Data_length: 16384
Max_data_length: 0
   Index_length: 0
      Data_free: 0
 Auto_increment: 6
    Create_time: 2023-01-10 19:25:51
    Update_time: 2023-01-10 19:25:51
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options: 
        Comment: 
1 row in set (0.02 sec)


mysql> ALTER TABLE orders ENGINE = MyISAM;
Query OK, 5 rows affected (0.12 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql> ALTER TABLE orders ENGINE = InnoDB;
Query OK, 5 rows affected (0.19 sec)
Records: 5  Duplicates: 0  Warnings: 0

mysql> show profiles;
+----------+------------+------------------------------------+
| Query_ID | Duration   | Query                              |
+----------+------------+------------------------------------+
|        1 | 0.01555000 | show table status                  |
|        2 | 0.12003300 | ALTER TABLE orders ENGINE = MyISAM |
|        3 | 0.18180750 | ALTER TABLE orders ENGINE = InnoDB |
+----------+------------+------------------------------------+
3 rows in set, 1 warning (0.00 sec)
```
Продолжительность переключения на MyISAM: 0,120  
ПроПродолжительность переключения на InnoDB: 0,181

4. Ex4
```commandline
[mysqld]
skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql

pid-file=/var/run/mysqld/mysqld.pid
[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/

#Set IO Speed
# 0 - скорость
# 1 - сохранность
# 2 - универсальный параметр
innodb_flush_log_at_trx_commit = 0 

#Set compression
# Barracuda - формат файла с сжатием
innodb_file_format=Barracuda

#Set buffer
innodb_log_buffer_size	= 1M

#Set Cache size
key_buffer_size = 7.011G

#Set log size
max_binlog_size	= 100M
```
