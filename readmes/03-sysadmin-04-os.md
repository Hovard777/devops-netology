1. 
```
cat /etc/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter

[Service]
User=vagrant
Group=vagrant
Type=simple
ExecStart=/usr/local/bin/node_exporter $MYVAR
EnvironmentFile=/etc/default/node_exporter_cfg
Restart=on-failure


[Install]
WantedBy=multi-user.target
```
```
vagrant@vagrant:~$ sudo systemctl daemon-reload
vagrant@vagrant:~$ sudo systemctl enable node_exporter.service
vagrant@vagrant:~$ sudo systemctl restart node_exporter 
vagrant@vagrant:~$ ps -ef|grep node
vagrant     1701       1  0 21:00 ?        00:00:00 /usr/local/bin/node_exporter
vagrant     1719    1383  0 21:01 pts/0    00:00:00 grep --color=auto node
vagrant@vagrant:~$ cat /proc/1701/environ 
LANG=en_US.UTF-8PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/binHOME=/home/vagrantLOGNAME=vagrantUSER=vagrantSHELL=/bin/bashINVOCATION_ID=2fa690aafdea47e297c4f0fa4ed5c57bJOURNAL_STREAM=9:32539MYVAR=some_val
```
Пробросил порты на хост и перезагрузился.
![img_1.png](img_1.png)

2.  
CPU:   
    node_cpu_seconds_total  
    node_cpu_seconds_total  
    node_cpu_seconds_total  
    process_cpu_seconds_total   

Memory:  
    node_memory_MemAvailable_bytes   
    node_memory_MemFree_bytes  
    
Disk  
    node_disk_io_time_seconds_total  
    node_disk_read_bytes_total  
    node_disk_read_time_seconds_total  
    node_disk_write_time_seconds_total  
    
Network:  
    node_network_receive_errs_total   
    node_network_receive_bytes_total   
    node_network_transmit_bytes_total  
    node_network_transmit_errs_total    

3.

![img_2.png](img_2.png)  

4. Да можно понять:
```
vagrant@vagrant:~$ dmesg |grep virt
[    0.015375] CPU MTRRs all blank - virtualized system.
[    1.487605] Booting paravirtualized kernel on KVM
[    9.315084] systemd[1]: Detected virtualization oracle.

```

5. 
```
vagrant@vagrant:~$ sysctl -n fs.nr_open
1048576
```
Это максимальное число открытых дескрипторов для ядра (системы), для пользователя задать больше этого числа нельзя (если не менять). 
Число задается кратное 1024, в данном случае =1024*1024.  
```commandline
vagrant@vagrant:~$ ulimit -Sn 
1024
vagrant@vagrant:~$ ulimit -Hn
1048576
```
Жесткий лимит `ulimit -Hn` не может быть изменен пользователем после его установки. Жесткие ограничения могут быть изменены только пользователем `root`. Мягкий лимит `ulimit -Sn`, однако, может быть изменен пользователем, но не может превышать жесткий лимит, т. е. он может иметь минимальное значение 0 и максимальное значение, равное «жесткому лимиту».  

6. 
```
root@vagrant:/# unshare  -f --pid --mount-proc sleep 2h
```
```commandline
root@vagrant:~# ps -e|grep sleep
   1758 pts/0    00:00:00 sleep
root@vagrant:~# nsenter --target 1758 --pid --mount
root@vagrant:/# ps
    PID TTY          TIME CMD
      2 pts/1    00:00:00 bash
     13 pts/1    00:00:00 ps
root@vagrant:/# ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   7228   580 pts/0    S+   21:56   0:00 sleep 2h
root           2  0.0  0.4   8960  4120 pts/1    S    21:58   0:00 -bash
root          14  0.0  0.3  10612  3324 pts/1    R+   21:58   0:00 ps aux
```  
7. `:(){ :|:& };:`

для понятности заменим : именем f и отформатируем код.

f() {  
  f | f &  
}  
f  

Таким образом это функция, которая параллельно пускает два своих экземпляра. Каждый пускает ещё по два и т.д. 
При отсутствии лимита на число процессов машина быстро исчерпывает физическую память и уходит в своп. Заканчиваются пиды (PID).
  Затем `cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-7.scope` помогает решить проблему.  
Установка придельного числа `ulimit -u` спасёт ситуацию.  
>https://www.cyberciti.biz/faq/understanding-bash-fork-bomb/