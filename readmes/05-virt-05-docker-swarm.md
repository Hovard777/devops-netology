1.   
- В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?  
Ответ: replication - реплика на заданное количество нод, global - репликация на каждую ноду
- Какой алгоритм выбора лидера используется в Docker Swarm кластере?  
Ответ: Лидер нода выбирается из управляючих нод путем Raft согласованного алгоритма.
- Что такое Overlay Network?  
Ответ: Overlay-сеть создает подсеть, которую могут использовать контейнеры в разных хостах swarm-кластера. Контейнеры на разных физических хостах могут обмениваться данными по overlay-сети (если все они прикреплены к одной сети).

2. 

```commandline
[root@node01 ~]# docker node ls
ID                            HOSTNAME             STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
l3rde753tnln3dyto4yp4ee3o *   node01.netology.yc   Ready     Active         Leader           20.10.21
n9n0btrvhbnmag1tljd1be4c3     node02.netology.yc   Ready     Active         Reachable        20.10.21
gf6de435d415pkf2q4o8exrey     node03.netology.yc   Ready     Active         Reachable        20.10.21
ibdrgpa5eoj48lki8iy7ao36n     node04.netology.yc   Ready     Active                          20.10.21
991fhblldavechpacc8r1at2x     node05.netology.yc   Ready     Active                          20.10.21
oxjphjpe7v584miszhnzb08cs     node06.netology.yc   Ready     Active                          20.10.21

```
3.   
```commandline
[root@node01 ~]# docker service ls
ID             NAME                                MODE         REPLICAS   IMAGE                                          PORTS
fz9kvczj97cc   swarm_monitoring_alertmanager       replicated   1/1        stefanprodan/swarmprom-alertmanager:v0.14.0    
yelyxobo2p26   swarm_monitoring_caddy              replicated   0/1        stefanprodan/caddy:latest                      *:3000->3000/tcp, *:9090->9090/tcp, *:9093-9094->9093-9094/tcp
sa5gt1mvk4h9   swarm_monitoring_cadvisor           global       6/6        google/cadvisor:latest                         
4mgklwv1sq6f   swarm_monitoring_dockerd-exporter   global       6/6        stefanprodan/caddy:latest                      
dc0ouady4nhv   swarm_monitoring_grafana            replicated   1/1        stefanprodan/swarmprom-grafana:5.3.4           
ollwnmeia2jf   swarm_monitoring_node-exporter      global       6/6        stefanprodan/swarmprom-node-exporter:v0.16.0   
kuto5qm1flfn   swarm_monitoring_prometheus         replicated   1/1        stefanprodan/swarmprom-prometheus:v2.5.0       
8v8p17ngb5v2   swarm_monitoring_unsee              replicated   1/1        cloudflare/unsee:v0.8.0                        

```