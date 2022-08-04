1. > Разрежённый файл (англ. sparse file) — файл, в котором последовательности нулевых байтов[1] заменены на информацию об этих последовательностях (список дыр).  
  Дыра (англ. hole) — последовательность нулевых байт внутри файла, не записанная на диск. Информация о дырах (смещение от начала файла в байтах и количество байт) хранится в метаданных ФС.


2. > В Linux каждый файл имеет уникальный идентификатор - индексный дескриптор (inode). Это число, которое однозначно идентифицирует файл в файловой системе. Жесткая ссылка и файл, для которой она создавалась имеют одинаковые inode. Поэтому жесткая ссылка имеет те же права доступа, владельца и время последней модификации, что и целевой файл. Различаются только имена файлов. Фактически жесткая ссылка это еще одно имя для файла.  
3. Done
4. Используя fdisk, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство. 
```
Device     Boot   Start     End Sectors  Size Id Type
/dev/sdb1          2048 4196351 4194304    2G 83 Linux
/dev/sdb2       4196352 5242879 1046528  511M 83 Linux
```

5.
```
   root@vagrant:~# sfdisk -d /dev/sdb | sfdisk /dev/sdc
   Checking that no-one is using this disk right now ... OK 
   
   Disk /dev/sdc: 2.51 GiB, 2684354560 bytes, 5242880 sectors
   Disk model: VBOX HARDDISK
   Units: sectors of 1 * 512 = 512 bytes
   Sector size (logical/physical): 512 bytes / 512 bytes
   I/O size (minimum/optimal): 512 bytes / 512 bytes 
   
   Created a new DOS disklabel with disk identifier 0x0e702756.
   /dev/sdc1: Created a new partition 1 of type 'Linux' and of size 2 GiB.
   /dev/sdc2: Created a new partition 2 of type 'Extended' and of size 511 MiB.
   /dev/sdc3: Done.
   
   New situation:
   Disklabel type: dos
   Disk identifier: 0x0e702756
   Device     Boot   Start     End Sectors  Size Id Type
   /dev/sdc1          2048 4196351 4194304    2G 83 Linux
   /dev/sdc2       4196352 5242879 1046528  511M 83 Linux
   ```
6. 
```
mdadm --create --verbose  /dev/md0 --level=mirror --raid-devices=2 /dev/sdb1 /dev/sdc1
mdadm --detail /dev/md0
/dev/md0:
           Version : 1.2
     Creation Time : Wed Aug  3 21:17:01 2022
        Raid Level : raid1
        Array Size : 2094080 (2045.00 MiB 2144.34 MB)
     Used Dev Size : 2094080 (2045.00 MiB 2144.34 MB)
      Raid Devices : 2
     Total Devices : 2
       Persistence : Superblock is persistent

       Update Time : Wed Aug  3 21:17:13 2022
             State : clean 
    Active Devices : 2
   Working Devices : 2
    Failed Devices : 0
     Spare Devices : 0

Consistency Policy : resync

              Name : vagrant:0  (local to host vagrant)
              UUID : 24f137da:902a6799:81675c9f:026b6136
            Events : 17

    Number   Major   Minor   RaidDevice State
       0       8       17        0      active sync   /dev/sdb1
       1       8       33        1      active sync   /dev/sdc1

```
7. 
```
mdadm --create --verbose  /dev/md1 --level=0 --raid-devices=2 /dev/sdb2 /dev/sdc2
lsblk
NAME                      MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
loop0                       7:0    0 61.9M  1 loop  /snap/core20/1328
loop1                       7:1    0 67.2M  1 loop  /snap/lxd/21835
loop2                       7:2    0 43.6M  1 loop  /snap/snapd/14978
loop3                       7:3    0 67.8M  1 loop  /snap/lxd/22753
loop4                       7:4    0   47M  1 loop  /snap/snapd/16292
loop5                       7:5    0   62M  1 loop  /snap/core20/1587
sda                         8:0    0   64G  0 disk  
├─sda1                      8:1    0    1M  0 part  
├─sda2                      8:2    0  1.5G  0 part  /boot
└─sda3                      8:3    0 62.5G  0 part  
  └─ubuntu--vg-ubuntu--lv 253:0    0 31.3G  0 lvm   /
sdb                         8:16   0  2.5G  0 disk  
├─sdb1                      8:17   0    2G  0 part  
│ └─md127                   9:127  0    2G  0 raid1 
└─sdb2                      8:18   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
sdc                         8:32   0  2.5G  0 disk  
├─sdc1                      8:33   0    2G  0 part  
│ └─md127                   9:127  0    2G  0 raid1 
└─sdc2                      8:34   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
```
8.  
```
    root@vagrant:~# pvcreate /dev/md0
    root@vagrant:~# pvcreate /dev/md1
    root@vagrant:~# pvs
    PV         VG        Fmt  Attr PSize    PFree
    /dev/md0             lvm2 ---    <2.00g   <2.00g
    /dev/md1             lvm2 ---  1018.00m 1018.00m
    /dev/sda3  ubuntu-vg lvm2 a--   <62.50g   31.25g
```

9. 
```
   root@vagrant:~# vgcreate vg-netology /dev/md0 /dev/md1
   root@vagrant:~# vgs
   VG          #PV #LV #SN Attr   VSize   VFree
   ubuntu-vg     1   1   0 wz--n- <62.50g 31.25g
   vg-netology   2   0   0 wz--n-  <2.99g <2.99g
```
10.  
``` 
root@vagrant:~# lvcreate -L100M -n lv100 vg-netology /dev/md1
  Logical volume "lv100" created.
root@vagrant:~# lvs
  LV        VG          Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
  ubuntu-lv ubuntu-vg   -wi-ao---- <31.25g                                                    
  lv100     vg-netology -wi-a----- 100.00m                         
```   

11. 
``` 
root@vagrant:~# mkfs.ext4 /dev/vg-netology/lv100 
```
12. 
``` 
root@vagrant:~# mount /dev/vg-netology/lv100 /tmp/new/
root@vagrant:~# df -h
Filesystem                         Size  Used Avail Use% Mounted on
udev                               445M     0  445M   0% /dev
tmpfs                               98M 1012K   97M   2% /run
/dev/mapper/ubuntu--vg-ubuntu--lv   31G  3.5G   26G  12% /
tmpfs                              489M     0  489M   0% /dev/shm
tmpfs                              5.0M     0  5.0M   0% /run/lock
tmpfs                              489M     0  489M   0% /sys/fs/cgroup
/dev/loop0                          62M   62M     0 100% /snap/core20/1328
/dev/loop1                          68M   68M     0 100% /snap/lxd/21835
/dev/loop2                          44M   44M     0 100% /snap/snapd/14978
/dev/sda2                          1.5G   76M  1.3G   6% /boot
vagrant                             39G   32G  6.8G  83% /vagrant
tmpfs                               98M     0   98M   0% /run/user/1000
/dev/mapper/vg--netology-lv100      93M   72K   86M   1% /tmp/new

```

13. 
``` 
root@vagrant:/tmp/new# ll
total 21664
drwxr-xr-x  3 root root     4096 Aug  4 16:52 ./
drwxrwxrwt 12 root root     4096 Aug  4 16:50 ../
drwx------  2 root root    16384 Aug  4 16:49 lost+found/
-rw-r--r--  1 root root 22157939 Aug  4 15:37 test.gz

```

14. 
``` 
root@vagrant:/tmp/new# lsb
lsb_release  lsblk        
root@vagrant:/tmp/new# lsblk 
NAME                      MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
loop0                       7:0    0 61.9M  1 loop  /snap/core20/1328
loop1                       7:1    0 67.2M  1 loop  /snap/lxd/21835
loop2                       7:2    0 43.6M  1 loop  /snap/snapd/14978
sda                         8:0    0   64G  0 disk  
├─sda1                      8:1    0    1M  0 part  
├─sda2                      8:2    0  1.5G  0 part  /boot
└─sda3                      8:3    0 62.5G  0 part  
  └─ubuntu--vg-ubuntu--lv 253:0    0 31.3G  0 lvm   /
sdb                         8:16   0  2.5G  0 disk  
├─sdb1                      8:17   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
└─sdb2                      8:18   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
    └─vg--netology-lv100  253:1    0  100M  0 lvm   /tmp/new
sdc                         8:32   0  2.5G  0 disk  
├─sdc1                      8:33   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
└─sdc2                      8:34   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
    └─vg--netology-lv100  253:1    0  100M  0 lvm   /tmp/new

```
15. 
``` 
root@vagrant:/tmp/new# gzip -t /tmp/new/test.gz
root@vagrant:/tmp/new# echo $?
0
```

16. 
``` 
root@vagrant:/tmp/new# pvmove  /dev/md1 /dev/md0
  /dev/md1: Moved: 100.00%
root@vagrant:/tmp/new# lsblk 
NAME                      MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
loop0                       7:0    0 61.9M  1 loop  /snap/core20/1328
loop1                       7:1    0 67.2M  1 loop  /snap/lxd/21835
loop2                       7:2    0 43.6M  1 loop  /snap/snapd/14978
sda                         8:0    0   64G  0 disk  
├─sda1                      8:1    0    1M  0 part  
├─sda2                      8:2    0  1.5G  0 part  /boot
└─sda3                      8:3    0 62.5G  0 part  
  └─ubuntu--vg-ubuntu--lv 253:0    0 31.3G  0 lvm   /
sdb                         8:16   0  2.5G  0 disk  
├─sdb1                      8:17   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
│   └─vg--netology-lv100  253:1    0  100M  0 lvm   /tmp/new
└─sdb2                      8:18   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
sdc                         8:32   0  2.5G  0 disk  
├─sdc1                      8:33   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
│   └─vg--netology-lv100  253:1    0  100M  0 lvm   /tmp/new
└─sdc2                      8:34   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
```

17. 
``` 
root@vagrant:/tmp/new# mdadm --fail /dev/md0 /dev/sdb1
mdadm: set /dev/sdb1 faulty in /dev/md0
```

18. 
``` 
[ 4163.968369] md/raid1:md0: Disk failure on sdb1, disabling device.
               md/raid1:md0: Operation continuing on 1 devices.
```
19. 
``` 
root@vagrant:/tmp/new# gzip -t /tmp/new/test.gz
root@vagrant:/tmp/new# echo $?
0
```