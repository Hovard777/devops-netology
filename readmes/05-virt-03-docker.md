## Задача 1
Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность: запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.
### Ответ
https://hub.docker.com/r/hovard777/netology
```commandline
vagrant@server1:~/nginx/docker$ docker run -d -p 81:80 hovard777/netology:latest
8082e5d142a1e0c80722ac10cb937e6c8fa4ac7e0f220a6f30b59e7f13815d94
vagrant@server1:~/nginx/docker$ docker ps
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS         PORTS                               NAMES
8082e5d142a1   hovard777/netology:latest   "/docker-entrypoint.…"   9 seconds ago   Up 6 seconds   0.0.0.0:81->80/tcp, :::81->80/tcp   laughing_antonelli
vagrant@server1:~/nginx/docker$ curl http://localhost:81
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>

```
## Задача 2
Посмотрите на сценарий ниже и ответьте на вопрос: "Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Сценарий:

- Высоконагруженное монолитное java веб-приложение;
- Nodejs веб-приложение;
- Мобильное приложение c версиями для Android и iOS;
- Шина данных на базе Apache Kafka;
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
- Мониторинг-стек на базе Prometheus и Grafana;
- MongoDB, как основное хранилище данных для java-приложения;
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.
### Ответ
>Высоконагруженное монолитное java веб-приложение - 
физический сервер, т.к. монолитное, и высоконагруженное - то необходим физический доступ к ресурсам, без использования гипервизора виртуалки.

>Nodejs веб-приложение - это веб приложение, для таких приложений достаточно докера, так гораздо удобнее выкатывать обновления

>Мобильное приложение c версиями для Android и iOS - Виртаулка - для приложения нужен GUI.

>Шина данных на базе Apache Kafka - Виртуалка, т.к. необходима сохранность данных и высокая производительность

> Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana - data nodes на физических серверах или виртуалках с SSD дисками, logstash на виртуалке, т.к. нужна производительность и удобнее в настройке. master nodes and kibana - можно в контейнерах

> Мониторинг-стек на базе Prometheus и Grafana - Prometheus на ВМ, т.к. требуется хранилище метрик, Grafana - app часть в контейнере, БД на виртуалке

> MongoDB, как основное хранилище данных для java-приложения - Виртуалка, т.к. требуется хранилище. физический сервер будет слишком накладно
 
>Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry - Виртуалка. т.к.требуется хранилище информации

## Задача 3
- Запустите первый контейнер из образа centos c любым тэгом в фоновом режиме, подключив папку /data из текущей рабочей директории на хостовой машине в /data контейнера;
- Запустите второй контейнер из образа debian в фоновом режиме, подключив папку /data из текущей рабочей директории на хостовой машине в /data контейнера;
- Подключитесь к первому контейнеру с помощью docker exec и создайте текстовый файл любого содержания в /data;
- Добавьте еще один файл в папку /data на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в /data контейнера.
### Ответ
```commandline
vagrant@server1:~$ docker ps
CONTAINER ID   IMAGE                     COMMAND            CREATED              STATUS              PORTS     NAMES
753ecdb3faa3   hovard777/debian:latest   "bash"             About a minute ago   Up About a minute             some-debian
a7098b3ce712   hovard777/centos:latest   "/usr/sbin/init"   25 minutes ago       Up 25 minutes                 some-centos
vagrant@server1:~$ docker exec -it some-centos /bin/bash
[root@a7098b3ce712 /]# ls
anaconda-post.log  bin  data  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
[root@a7098b3ce712 /]# touch /data/centos-file
touch: cannot touch '/data/centos-file': Read-only file system
[root@a7098b3ce712 /]# exit
exit
vagrant@server1:~$ docker exec -it some-centos /bin/bash
[root@3fad97ff8bf6 /]# ls
anaconda-post.log  bin  data  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
[root@3fad97ff8bf6 /]# touch /data/myfile
[root@3fad97ff8bf6 /]# echo netology>/data/myfile 
[root@3fad97ff8bf6 /]# exit
exit
vagrant@server1:~$ docker exec -it some-debian /bin/bash
root@057a599889fd:/# ls /data/
my-file-host  myfile
root@057a599889fd:/# cat /data/myfile 
netology
root@057a599889fd:/# exit
exit
vagrant@server1:~$ pwd
/home/vagrant
vagrant@server1:~$ ls
get-docker.sh  nginx
vagrant@server1:~$ cd nginx/
vagrant@server1:~/nginx$ ll data/
total 12
drwxrwxr-x 2 vagrant vagrant 4096 Oct 19 17:57 ./
drwxrwxr-x 6 vagrant vagrant 4096 Oct 19 17:18 ../
-rw-rw-r-- 1 vagrant vagrant    0 Oct 19 17:57 my-file-host
-rw-r--r-- 1 root    root       9 Oct 19 17:56 myfile

```