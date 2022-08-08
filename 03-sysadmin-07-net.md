1. 
```
   vagrant@vagrant:~$ ip -c -br l
lo               UNKNOWN        00:00:00:00:00:00 <LOOPBACK,UP,LOWER_UP> 
eth0             UP             08:00:27:a2:6b:fd <BROADCAST,MULTICAST,UP,LOWER_UP> 
```
В Windows запустить приложение "Командная строка" и в открывшемся окне введите команду ipconfig /all

2. LLDP - протокол для обмена информацией между соседними устройствами. В Linux надо установить пакет lldpd.
```
ifebres@ifebres-virtual-machine:~/GIT$ lldpctl 
-------------------------------------------------------------------------------
LLDP neighbors:
-------------------------------------------------------------------------------
Interface:    ens33, via: LLDP, RID: 1, Time: 0 day, 00:47:35
  Chassis:     
    ChassisID:    unhandled 46 45 42 52 45 53 2d 49
  Port:        
    PortID:       mac 00:50:56:c0:00:08
    PortDescr:    {E1FE3508-8455-4378-8A0E-5728372C2EC0}
    TTL:          121
  LLDP-MED:    
    Device Type:  Generic Endpoint (Class I)
    Capability:   Capabilities, yes
    Capability:   Policy, yes
    Capability:   Location, yes
    Capability:   Inventory, yes
    Inventory:   
      Software Revision: Windows 10 Enterprise
      Model:        MSI on Windows
-------------------------------------------------------------------------------
```
3. VLAN - виртуальное разделение коммутатора. В Linux пакет vlan. 
``` 
vi /etc/network/interfaces
auto vlan1400
iface vlan1400 inet static
        address 192.168.1.1
        netmask 255.255.255.0
        vlan_raw_device eth0
auto eth0.1400
iface eth0.1400 inet static
        address 192.168.1.1        
        netmask 255.255.255.0        
        vlan_raw_device eth0
```
Вручную настройка VLAN выполняется с помощью программы vconfig 

4. Типы LAG:
- статический (на Cisco mode on);
- динамический – LACP протокол (на Cisco mode active).
Пример конфига:
``` 
$ sudo nano /etc/network/interfaces
# The primary network interface
auto bond0
iface bond0 inet static
    address 192.168.1.150
    netmask 255.255.255.0    
    gateway 192.168.1.1
    dns-nameservers 192.168.1.1 8.8.8.8
    dns-search domain.local
        slaves eth0 eth1
        bond_mode 0
        bond-miimon 100
        bond_downdelay 200
        bound_updelay 200
```
Режимы работы:
mode=0 (balance-rr)
При этом методе объединения трафик распределяется по принципу «карусели»: пакеты по очереди направляются на сетевые карты объединённого интерфейса. Например, если у нас есть физические интерфейсы eth0, eth1, and eth2, объединенные в bond0, первый пакет будет отправляться через eth0, второй — через eth1, третий — через eth2, а четвертый снова через eth0 и т.д.

mode=1 (active-backup)
Когда используется этот метод, активен только один физический интерфейс, а остальные работают как резервные на случай отказа основного.

mode=2 (balance-xor)
В данном случае объединенный интерфейс определяет, через какую физическую сетевую карту отправить пакеты, в зависимости от MAC-адресов источника и получателя.

mode=3 (broadcast) Широковещательный режим, все пакеты отправляются через каждый интерфейс. Имеет ограниченное применение, но обеспечивает значительную отказоустойчивость.

mode=4 (802.3ad)
Особый режим объединения. Для него требуется специально настраивать коммутатор, к которому подключен объединенный интерфейс. Реализует стандарты объединения каналов IEEE и обеспечивает как увеличение пропускной способности, так и отказоустойчивость.

mode=5 (balance-tlb)
Распределение нагрузки при передаче. Входящий трафик обрабатывается в обычном режиме, а при передаче интерфейс определяется на основе данных о загруженности.

mode=6 (balance-alb)
Адаптивное распределение нагрузки. Аналогично предыдущему режиму, но с возможностью балансировать также входящую нагрузку.

5. в сети с маской /29 - 8 адресов. 
> Сколько /29 подсетей можно получить из сети с маской /24 - 32 подсети.  

Приведите несколько примеров /29 подсетей внутри сети 10.10.10.0/24.

- 10.10.10.10/29
- 10.10.10.19/29
- 10.10.10.28/29

6. Возьмём из Carrier-Grade NAT - 100.64.5.0/26  

7. Windows
```
arp -a
Интерфейс: 10.153.0.220 --- 0xb
  адрес в Интернете      Физический адрес      Тип
  10.153.0.1            10-38-d2-2a-ad-06     динамический
  10.153.0.4            b0-82-dc-25-fb-ea     динамический
  10.153.0.5            00-90-4c-8b-86-86     динамический
  10.153.0.13           ac-a6-2d-7e-b5-ec     динамический
  10.153.0.14           ac-c6-2d-78-a1-ac     динамический
  10.153.0.15           44-9e-a1-d9-d5-58     динамический
  
  Чтобы очистить ARP-кэш, необходимо выполнить команду: netsh interface ip delete arpcache После чего ARP-таблица будет очищена. Либо arp -d
arp -d <ip-address> - Удаляет только нужный адрес
```  
Linux
``` 
ifebres@ifebres-virtual-machine:~/GIT/ELK/KUBERNETES$ arp -e
Адрес HW-тип HW-адрес Флаги Маска Интерфейс
192.168.46.254           ether   00:50:56:e2:a6:61   C                     ens33
_gateway                 ether   00:50:56:e0:ec:ac   C                     ens33

Clear ARP
root@vagrant:~# ip -s -s neigh flush all
10.0.2.2 dev eth0 lladdr 52:54:00:12:35:02 ref 1 used 0/0/0 probes 1 DELAY
10.0.2.3 dev eth0 lladdr 52:54:00:12:35:03 used 5335/5335/5319 probes 1 STALE

*** Round 1, deleting 2 entries ***
10.0.2.2 dev eth0 lladdr 52:54:00:12:35:02 ref 1 used 0/0/0 probes 4 REACHABLE

*** Round 2, deleting 1 entries ***
10.0.2.2 dev eth0  ref 1 used 0/0/0 probes 4 INCOMPLETE

*** Round 3, deleting 1 entries ***
*** Flush is complete after 3 rounds ***

arp -d <ip-address> - Удаляет только нужный адрес

```
