1. Install PG
```commandline
[ifebres@febres-i ~]$ sudo docker pull postgres:13
[ifebres@febres-i ~]$ docker volume create vol_postgres
[ifebres@febres-i ~]$ sudo docker run --rm --name pg-docker -e POSTGRES_PASSWORD=postgres -ti -p 5432:5432 -v vol_postgres:/var/lib/postgresql/data postgres:13
[ifebres@febres-i ~]$ sudo docker exec -it pg-docker bash
[sudo] пароль для ifebres: 
root@447bfffbe878:/# psql -h localhost -p 5432 -U postgres -W
Password: 
psql (13.9 (Debian 13.9-1.pgdg110+1))
Type "help" for help.
```
Вывод списка БД:
```
postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)
```
Подключения к БД:
```
postgres=# \c postgres
Password: 
```
Вывод списка таблиц: \dt. Параметр S - для системных объектов.
```commandline
postgres=# \dtS
                    List of relations
   Schema   |          Name           | Type  |  Owner   
------------+-------------------------+-------+----------
 pg_catalog | pg_aggregate            | table | postgres
 pg_catalog | pg_am                   | table | postgres
 pg_catalog | pg_amop                 | table | postgres
 pg_catalog | pg_amproc               | table | postgres
 pg_catalog | pg_attrdef              | table | postgres
 pg_catalog | pg_attribute            | table | postgres
```
Вывод описания содержимого таблиц:
```commandline
postgres=# \dS+ pg_language
                                   Table "pg_catalog.pg_language"
    Column     |   Type    | Collation | Nullable | Default | Storage  | Stats target | Description 
---------------+-----------+-----------+----------+---------+----------+--------------+-------------
 oid           | oid       |           | not null |         | plain    |              | 
 lanname       | name      |           | not null |         | plain    |              | 
 lanowner      | oid       |           | not null |         | plain    |              | 
 lanispl       | boolean   |           | not null |         | plain    |              | 
 lanpltrusted  | boolean   |           | not null |         | plain    |              | 
 lanplcallfoid | oid       |           | not null |         | plain    |              | 
 laninline     | oid       |           | not null |         | plain    |              | 
 lanvalidator  | oid       |           | not null |         | plain    |              | 
 lanacl        | aclitem[] |           |          |         | extended |              | 
Indexes:
    "pg_language_name_index" UNIQUE, btree (lanname)
    "pg_language_oid_index" UNIQUE, btree (oid)
Access method: heap
```
Выход из psql:
```commandline
postgres=# \q
root@447bfffbe878:/# 
```
2. Create test_database
```commandline
postgres=# CREATE DATABASE test_database;
CREATE DATABASE
root@447bfffbe878:/var/lib/postgresql/data# psql -U postgres -f ./test_dump.sql test_database
postgres=# \c test_database 
Password: 
You are now connected to database "test_database" as user "postgres".
test_database=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner   
--------+--------+-------+----------
 public | orders | table | postgres
(1 row)

test_database=# ANALYZE VERBOSE public.orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE

test_database=# select tablename, attname,avg_width from pg_stats where tablename='orders' ORDER BY avg_width DESC LIMIT 1;
 tablename | attname | avg_width 
-----------+---------+-----------
 orders    | title   |        16
(1 row)

```
3. Sharding
Шардирование:
```commandline
test_database=# CREATE TABLE orders_1 (CHECK (price > 499)) INHERITS (orders);
CREATE TABLE
test_database=# CREATE TABLE orders_2 (CHECK (price <= 499)) INHERITS (orders);
CREATE TABLE
test_database=# INSERT INTO orders_1 SELECT * FROM orders WHERE price > 499;
INSERT 0 4
test_database=# DELETE FROM only orders WHERE price > 499;
DELETE 4
test_database=# INSERT INTO orders_2 SELECT * FROM orders WHERE price <= 499;
INSERT 0 4
test_database=# DELETE FROM only orders WHERE price <= 499;
DELETE 4
```
Изначально можно было создать:
```
test_database=# CREATE TABLE public.orders_new (
id integer NOT NULL,
title character varying(80) NOT NULL,
price integer DEFAULT 0
)
PARTITION BY RANGE (price);
CREATE TABLE

test_database=# CREATE TABLE orders_new1 PARTITION OF orders_new FOR VALUES FROM (500) TO (9999999);
CREATE TABLE
test_database=# CREATE TABLE orders_new2 PARTITION OF orders_new FOR VALUES FROM (0) TO (499);
CREATE TABLE
```
4. Backup
```commandline
root@447bfffbe878:/var/lib/postgresql/data# pg_dump -U postgres -d test_database >test_pg_database_dump.sql
root@447bfffbe878:/var/lib/postgresql/data# ls -la|grep test_pg
-rw-r--r--. 1 root     root      4808 Jan 16 21:06 test_pg_database_dump.sql
```
Чтобы добавить уникальность значения столбца ```title``` для таблиц ```test_database``` можно воспользоваться несколькими вариантами. 
создать уникальную связку 
```
ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_uniqe UNIQUE (id,title);
 ```
или составной первичный ключ 
```
CREATE TABLE public.orders (
    id integer NOT NULL,
    title character varying(80) NOT NULL,
    price integer DEFAULT 0
    PRIMARY KEY (id,title)
);
```
или задать для столбца только уникальные значения 
```
CREATE TABLE public.orders (
    id integer NOT NULL,
    title character varying(80) NOT NULL,
    price integer DEFAULT 0
    UNIQUE (title)
);
```
