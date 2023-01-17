1. Build docker
Dockerfile
```commandline
#6.5. Elasticsearch
FROM centos:7
LABEL ElasticSearch Lab
MAINTAINER Ivan Febres <hovard777@gmail.com>
ENV PATH=/usr/lib:/usr/lib/jvm/jre-11/bin:$PATH

RUN yum install java-11-openjdk -y 
RUN yum install wget -y 

RUN wget https://fossies.org/linux/www/elasticsearch-8.6.0-linux-x86_64.tar.gz 
RUN yum install perl-Digest-SHA -y 
RUN tar -xzf elasticsearch-8.6.0-linux-x86_64.tar.gz \
    && yum upgrade -y
    
ADD elasticsearch.yml /elasticsearch-8.6.0/config/
ENV JAVA_HOME=/elasticsearch-8.6.0/jdk/
ENV ES_HOME=/elasticsearch-8.6.0
RUN groupadd elasticsearch \
    && useradd -g elasticsearch elasticsearch
    
RUN mkdir /var/lib/logs \
    && chown elasticsearch:elasticsearch /var/lib/logs \
    && mkdir /var/lib/data \
    && chown elasticsearch:elasticsearch /var/lib/data \
    && chown -R elasticsearch:elasticsearch /elasticsearch-8.6.0/
RUN mkdir /elasticsearch-8.6.0/snapshots &&\
    chown elasticsearch:elasticsearch /elasticsearch-8.6.0/snapshots
    
USER elasticsearch
CMD ["/usr/sbin/init"]
CMD ["/elasticsearch-8.6.0/bin/elasticsearch"]
```
[elasticsearch.yml] [elasticsearch.yml](..%2Felastic%2Felasticsearch.yml)  

Ссылка на репозиторий
https://hub.docker.com/repository/docker/hovard777/elk/general  

 Ответ от localhost:9200
```commandline
name	"a43b94626a1c"
cluster_name	"netology_test"
cluster_uuid	"fP2FgMWPRGKAQJvbXo9_Gg"
version	
number	"8.6.0"
build_flavor	"default"
build_type	"tar"
build_hash	"f67ef2df40237445caa70e2fef79471cc608d70d"
build_date	"2023-01-04T09:35:21.782467981Z"
build_snapshot	false
lucene_version	"9.4.2"
minimum_wire_compatibility_version	"7.17.0"
minimum_index_compatibility_version	"7.0.0"
tagline	"You Know, for Search"
```
2. Index
Create
```commandline
[ifebres@febres-i elastic]$ curl -k -X PUT -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/ind-1 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
[ifebres@febres-i elastic]$ curl -k -X PUT -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/ind-2 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 2,  "number_of_replicas": 1 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"ind-2"}

[ifebres@febres-i elastic]$ curl -k -X PUT -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/ind-3 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 4,  "number_of_replicas": 2 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"ind-3"}
```
List
```commandline
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cat/indices?v'
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   ind-1 T-AGHS0sRta3WqxxOjc8rg   1   0          0            0       225b           225b
yellow open   ind-3 6viuTPhLTcGJyO7J4OMhWg   4   2          0            0       900b           900b
yellow open   ind-2 MIUiytLXTO-AAYprXRLIzQ   2   1          0            0       450b           450b
```
Status index
```commandline
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cluster/health?pretty&level=indices'
{
  "cluster_name" : "netology_test",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 9,
  "active_shards" : 9,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 47.368421052631575,
  "indices" : {
    ".geoip_databases" : {
      "status" : "green",
      "number_of_shards" : 1,
      "number_of_replicas" : 0,
      "active_primary_shards" : 1,
      "active_shards" : 1,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 0
    },
    ".security-7" : {
      "status" : "green",
      "number_of_shards" : 1,
      "number_of_replicas" : 0,
      "active_primary_shards" : 1,
      "active_shards" : 1,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 0
    },
    "ind-1" : {
      "status" : "green",
      "number_of_shards" : 1,
      "number_of_replicas" : 0,
      "active_primary_shards" : 1,
      "active_shards" : 1,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 0
    },
    "ind-3" : {
      "status" : "yellow",
      "number_of_shards" : 4,
      "number_of_replicas" : 2,
      "active_primary_shards" : 4,
      "active_shards" : 4,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 8
    },
    "ind-2" : {
      "status" : "yellow",
      "number_of_shards" : 2,
      "number_of_replicas" : 1,
      "active_primary_shards" : 2,
      "active_shards" : 2,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 2
    }
  }
}
```
Status Cluster
```commandline
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cluster/health/?pretty=true'
{
  "cluster_name" : "netology_test",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 9,
  "active_shards" : 9,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 47.368421052631575
}
```
Remove index
```commandline
[ifebres@febres-i elastic]$ curl -k -X DELETE -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/ind-1?pretty'
{
  "acknowledged" : true
}
[ifebres@febres-i elastic]$ curl -k -X DELETE -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/ind-2?pretty'
{
  "acknowledged" : true
}
[ifebres@febres-i elastic]$ curl -k -X DELETE -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/ind-3?pretty'
{
  "acknowledged" : true
}
```
Индексы в статусе Yellow потому что у них указано число реплик, а по факту нет других серверов, соответственно реплицировать некуда.  

3. Backup
```commandline
[ifebres@febres-i elastic]$ curl -k -X POST -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/_snapshot/netology_backup?pretty -H 'Content-Type: application/json' -d'{"type": "fs", "settings": { "location":"/elasticsearch-8.6.0/snapshots" }}'
{
  "acknowledged" : true
}
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_snapshot/netology_backup?pretty'
{
  "netology_backup" : {
    "type" : "fs",
    "settings" : {
      "location" : "/elasticsearch-8.6.0/snapshots"
    }
  }
}
```
create index test
```commandline
[ifebres@febres-i elastic]$ curl -k -X PUT -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/test -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
{"acknowledged":true,"shards_acknowledged":true,"index":"test"}[ifebres@febres-i elastic]$ 

[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cat/indices?v'
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test  loT5RipwSmarxBOB3AuQfw   1   0          0            0       225b           225b
```
create snapshot
```
[ifebres@febres-i elastic]$ curl -k -X PUT -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/_snapshot/netology_backup/elasticsearch?wait_for_completion=true
{"snapshot":
{"snapshot":"elasticsearch",
 "uuid":"u1lXaPuuTj-zq_po7O1W4A",
  "repository":"netology_backup",
  "version_id":8060099,
  "version":"8.6.0",
  "indices":["test",".geoip_databases",".security-7"],
  "data_streams":[],
  "include_global_state":true,
  "state":"SUCCESS",
  "start_time":"2023-01-17T16:33:12.398Z",
  "start_time_in_millis":1673973192398,
  "end_time":"2023-01-17T16:33:13.399Z",
  "end_time_in_millis":1673973193399,
  "duration_in_millis":1001,
  "failures":[],
  "shards":{"total":3,
  "failed":0,
  "successful":3},
  "feature_states":[{"feature_name":"geoip","indices":[".geoip_databases"]},{"feature_name":"security","indices":[".security-7"]}]}}
  [ifebres@febres-i elastic]$ 

[elasticsearch@a43b94626a1c /]$ ls -la /elasticsearch-8.6.0/snapshots/
total 36
drwxr-xr-x. 1 elasticsearch elasticsearch   176 Jan 17 16:33 .
drwxr-xr-x. 1 elasticsearch elasticsearch   156 Jan 17 09:11 ..
-rw-r--r--. 1 elasticsearch elasticsearch  1098 Jan 17 16:33 index-0
-rw-r--r--. 1 elasticsearch elasticsearch     8 Jan 17 16:33 index.latest
drwxr-xr-x. 1 elasticsearch elasticsearch   132 Jan 17 16:33 indices
-rw-r--r--. 1 elasticsearch elasticsearch 18714 Jan 17 16:33 meta-u1lXaPuuTj-zq_po7O1W4A.dat
-rw-r--r--. 1 elasticsearch elasticsearch   392 Jan 17 16:33 snap-u1lXaPuuTj-zq_po7O1W4A.dat
```
remove index test
```commandline
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cat/indices?v'
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 SiRCAxexQ4iQzjIHuwIYfw   1   0          0            0       225b           225b
[ifebres@febres-i elastic]$ 
```
restore
```commandline
[ifebres@febres-i elastic]$ curl -k -X POST -u elastic:SlNuNmu7tsVL8DkvCMRR https://localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty -H 'Content-Type: application/json' -d'{"include_global_state":true}'
{
  "accepted" : true
}
[ifebres@febres-i elastic]$ curl -k -X GET -u elastic:SlNuNmu7tsVL8DkvCMRR 'https://localhost:9200/_cat/indices?v'
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 SiRCAxexQ4iQzjIHuwIYfw   1   0          0            0       225b           225b
green  open   test   fSrWJ1HRTXGvgOJQuuXC7Q   1   0          0            0       225b           225b
```

