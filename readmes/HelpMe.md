для задания 5-5 docker swarm

1.Авторизуемся в Yandex.Cloud.  
```commandline
ivan@ivan-virtual-machine:~/packer$ yc init 
Welcome! This command will take you through the configuration process.
Pick desired action:
 [1] Re-initialize this profile 'netology-terraform-profile' with new settings 
 [2] Create a new profile
 [3] Switch to and re-initialize existing profile: 'default'
 [4] Switch to and re-initialize existing profile: 'netology-ifebres'
 [5] Switch to and re-initialize existing profile: 'sa-profile-terraform'
Please enter your numeric choice: 4
Please go to https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990a*********a7bec2fb in order to obtain OAuth token.

Please enter OAuth token: [y0_AgAAAAAAnCt******************************M9_vjY876EfTJU] 
You have one cloud available: 'netology-cloud-febres' (id = b1grp5c*****b323pqkm). It is going to be used by default.
Please choose folder to use:
 [1] default (id = b1gr1****e10ct1g0t8q)
 [2] Create a new folder
Please enter your numeric choice: 2
Please enter a folder name: netology-swarm
Your current folder has been set to 'netology-swarm' (id = b1g6****utu4gsqs).
Do you want to configure a default Compute zone? [Y/n] y
Which zone do you want to use as a profile default?
 [1] ru-central1-a
 [2] ru-central1-b
 [3] ru-central1-c
 [4] Don't set default zone
Please enter your numeric choice: 1
Your profile default Compute zone has been set to 'ru-central1-a'.
There is a new yc version '0.98.0' available. Current version: '0.97.0'.
See release notes at https://cloud.yandex.ru/docs/cli/release-notes
You can install it by running the following command in your shell:
	$ yc components update

```
2.Создаём сеть и подсеть, чтобы собрать образ ОС с помощью Packer и запускаем сборку образа.  
```commandline
ivan@ivan-virtual-machine:~/packer$ yc vpc network create --name net --labels my-label=netology --description "my network first"
id: enp98mp***icnbp3
folder_id: b1g6****9utu4gsqs
created_at: "2022-11-14T19:37:17Z"
name: net
description: my network first
labels:
  my-label: netology

ivan@ivan-virtual-machine:~/packer$ yc vpc subnet create --name my-subnet-a --zone ru-central1-a --range 10.1.2.0/24 --network-name net --description "my subnet first"
id: e9bv2k***g3f296
folder_id: b1g66m***utu4gsqs
created_at: "2022-11-14T19:38:27Z"
name: my-subnet-a
description: my subnet first
network_id: enp98m****6icnbp3
zone_id: ru-central1-a
v4_cidr_blocks:
  - 10.1.2.0/24
ivan@ivan-virtual-machine:~/packer$ yc vpc subnet list
+----------------------+-------------+----------------------+----------------+---------------+---------------+
|          ID          |    NAME     |      NETWORK ID      | ROUTE TABLE ID |     ZONE      |     RANGE     |
+----------------------+-------------+----------------------+----------------+---------------+---------------+
| e9bv2*******ugg3f296 | my-subnet-a | enp98************bp3 |                | ru-central1-a | [10.1.2.0/24] |
+----------------------+-------------+----------------------+----------------+---------------+---------------+

```
Правим файл образа centos.json. Вставляем правильный folder_id и subnet_id   
folder_id: b1g66****9utu4gsqs  
subnet_id: e9bv2k******cugg3f296

```commandline
ivan@ivan-virtual-machine:~/packer$ ./packer validate centos-7-base.json 
The configuration is valid.

ivan@ivan-virtual-machine:~/packer$ ./packer build centos-7-base.json 
yandex: output will be in this color.

==> yandex: Creating temporary RSA SSH key for instance...
==> yandex: Using as source image: fd89dg08jjghmn88ut7p (name: "centos-7-v20221114", family: "centos-7")
==> yandex: Use provided subnet id e9bv2k91jb7cugg3f296
==> yandex: Creating disk...
==> yandex: Creating instance...
==> yandex: Waiting for instance with id fhm57j8vd655175r0mb9 to become active...
    yandex: Detected instance IP: 84.252.131.10
==> yandex: Using SSH communicator to connect: 84.252.131.10
==> yandex: Waiting for SSH to become available...
==> yandex: Connected to SSH!
==> yandex: Provisioning with shell script: /tmp/packer-shell3867838152
    yandex: Loaded plugins: fastestmirror
    yandex: Loading mirror speeds from cached hostfile
    yandex:  * base: centos-mirror.rbc.ru
    yandex:  * extras: centos-mirror.rbc.ru
    yandex:  * updates: mirror.yandex.ru
    yandex: No packages marked for update
    yandex: Loaded plugins: fastestmirror
    yandex: Loading mirror speeds from cached hostfile
    yandex:  * base: centos-mirror.rbc.ru
    yandex:  * extras: centos-mirror.rbc.ru
    yandex:  * updates: mirror.yandex.ru
    yandex: Package iptables-1.4.21-35.el7.x86_64 already installed and latest version
    yandex: Package curl-7.29.0-59.el7_9.1.x86_64 already installed and latest version
    yandex: Package net-tools-2.0-0.25.20131004git.el7.x86_64 already installed and latest version
    yandex: Package rsync-3.1.2-11.el7_9.x86_64 already installed and latest version
    yandex: Package openssh-server-7.4p1-22.el7_9.x86_64 already installed and latest version
    yandex: Resolving Dependencies
    yandex: --> Running transaction check
    yandex: ---> Package bind-utils.x86_64 32:9.11.4-26.P2.el7_9.10 will be installed
    yandex: --> Processing Dependency: bind-libs-lite(x86-64) = 32:9.11.4-26.P2.el7_9.10 for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: bind-libs(x86-64) = 32:9.11.4-26.P2.el7_9.10 for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: liblwres.so.160()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libisccfg.so.160()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libisc.so.169()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libirs.so.160()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libdns.so.1102()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libbind9.so.160()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: --> Processing Dependency: libGeoIP.so.1()(64bit) for package: 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64
    yandex: ---> Package bridge-utils.x86_64 0:1.5-9.el7 will be installed
    yandex: ---> Package tcpdump.x86_64 14:4.9.2-4.el7_7.1 will be installed
    yandex: --> Processing Dependency: libpcap >= 14:1.5.3-10 for package: 14:tcpdump-4.9.2-4.el7_7.1.x86_64
    yandex: --> Processing Dependency: libpcap.so.1()(64bit) for package: 14:tcpdump-4.9.2-4.el7_7.1.x86_64
    yandex: ---> Package telnet.x86_64 1:0.17-66.el7 will be installed
    yandex: --> Running transaction check
    yandex: ---> Package GeoIP.x86_64 0:1.5.0-14.el7 will be installed
    yandex: --> Processing Dependency: geoipupdate for package: GeoIP-1.5.0-14.el7.x86_64
    yandex: ---> Package bind-libs.x86_64 32:9.11.4-26.P2.el7_9.10 will be installed
    yandex: --> Processing Dependency: bind-license = 32:9.11.4-26.P2.el7_9.10 for package: 32:bind-libs-9.11.4-26.P2.el7_9.10.x86_64
    yandex: ---> Package bind-libs-lite.x86_64 32:9.11.4-26.P2.el7_9.10 will be installed
    yandex: ---> Package libpcap.x86_64 14:1.5.3-13.el7_9 will be installed
    yandex: --> Running transaction check
    yandex: ---> Package bind-license.noarch 32:9.11.4-26.P2.el7_9.10 will be installed
    yandex: ---> Package geoipupdate.x86_64 0:2.5.0-1.el7 will be installed
    yandex: --> Finished Dependency Resolution
    yandex:
    yandex: Dependencies Resolved
    yandex:
    yandex: ================================================================================
    yandex:  Package            Arch       Version                        Repository   Size
    yandex: ================================================================================
    yandex: Installing:
    yandex:  bind-utils         x86_64     32:9.11.4-26.P2.el7_9.10       updates     261 k
    yandex:  bridge-utils       x86_64     1.5-9.el7                      base         32 k
    yandex:  tcpdump            x86_64     14:4.9.2-4.el7_7.1             base        422 k
    yandex:  telnet             x86_64     1:0.17-66.el7                  updates      64 k
    yandex: Installing for dependencies:
    yandex:  GeoIP              x86_64     1.5.0-14.el7                   base        1.5 M
    yandex:  bind-libs          x86_64     32:9.11.4-26.P2.el7_9.10       updates     158 k
    yandex:  bind-libs-lite     x86_64     32:9.11.4-26.P2.el7_9.10       updates     1.1 M
    yandex:  bind-license       noarch     32:9.11.4-26.P2.el7_9.10       updates      91 k
    yandex:  geoipupdate        x86_64     2.5.0-1.el7                    base         35 k
    yandex:  libpcap            x86_64     14:1.5.3-13.el7_9              updates     139 k
    yandex:
    yandex: Transaction Summary
    yandex: ================================================================================
    yandex: Install  4 Packages (+6 Dependent packages)
    yandex:
    yandex: Total download size: 3.8 M
    yandex: Installed size: 9.0 M
    yandex: Downloading packages:
    yandex: --------------------------------------------------------------------------------
    yandex: Total                                              8.1 MB/s | 3.8 MB  00:00
    yandex: Running transaction check
    yandex: Running transaction test
    yandex: Transaction test succeeded
    yandex: Running transaction
    yandex:   Installing : 32:bind-license-9.11.4-26.P2.el7_9.10.noarch                1/10
    yandex:   Installing : geoipupdate-2.5.0-1.el7.x86_64                              2/10
    yandex:   Installing : GeoIP-1.5.0-14.el7.x86_64                                   3/10
    yandex:   Installing : 32:bind-libs-lite-9.11.4-26.P2.el7_9.10.x86_64              4/10
    yandex:   Installing : 32:bind-libs-9.11.4-26.P2.el7_9.10.x86_64                   5/10
    yandex:   Installing : 14:libpcap-1.5.3-13.el7_9.x86_64                            6/10
    yandex: pam_tally2: Error opening /var/log/tallylog for update: Permission denied
    yandex: pam_tally2: Authentication error
    yandex: useradd: failed to reset the tallylog entry of user "tcpdump"
    yandex:   Installing : 14:tcpdump-4.9.2-4.el7_7.1.x86_64                           7/10
    yandex:   Installing : 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64                  8/10
    yandex:   Installing : bridge-utils-1.5-9.el7.x86_64                               9/10
    yandex:   Installing : 1:telnet-0.17-66.el7.x86_64                                10/10
    yandex:   Verifying  : GeoIP-1.5.0-14.el7.x86_64                                   1/10
    yandex:   Verifying  : 14:libpcap-1.5.3-13.el7_9.x86_64                            2/10
    yandex:   Verifying  : 1:telnet-0.17-66.el7.x86_64                                 3/10
    yandex:   Verifying  : geoipupdate-2.5.0-1.el7.x86_64                              4/10
    yandex:   Verifying  : 32:bind-license-9.11.4-26.P2.el7_9.10.noarch                5/10
    yandex:   Verifying  : 32:bind-libs-9.11.4-26.P2.el7_9.10.x86_64                   6/10
    yandex:   Verifying  : 14:tcpdump-4.9.2-4.el7_7.1.x86_64                           7/10
    yandex:   Verifying  : bridge-utils-1.5-9.el7.x86_64                               8/10
    yandex:   Verifying  : 32:bind-libs-lite-9.11.4-26.P2.el7_9.10.x86_64              9/10
    yandex:   Verifying  : 32:bind-utils-9.11.4-26.P2.el7_9.10.x86_64                 10/10
    yandex:
    yandex: Installed:
    yandex:   bind-utils.x86_64 32:9.11.4-26.P2.el7_9.10   bridge-utils.x86_64 0:1.5-9.el7
    yandex:   tcpdump.x86_64 14:4.9.2-4.el7_7.1            telnet.x86_64 1:0.17-66.el7
    yandex:
    yandex: Dependency Installed:
    yandex:   GeoIP.x86_64 0:1.5.0-14.el7
    yandex:   bind-libs.x86_64 32:9.11.4-26.P2.el7_9.10
    yandex:   bind-libs-lite.x86_64 32:9.11.4-26.P2.el7_9.10
    yandex:   bind-license.noarch 32:9.11.4-26.P2.el7_9.10
    yandex:   geoipupdate.x86_64 0:2.5.0-1.el7
    yandex:   libpcap.x86_64 14:1.5.3-13.el7_9
    yandex:
    yandex: Complete!
==> yandex: Stopping instance...
==> yandex: Deleting instance...
    yandex: Instance has been deleted!
==> yandex: Creating image: centos-7-base
==> yandex: Waiting for image to complete...
==> yandex: Success image create...
==> yandex: Destroying boot disk...
    yandex: Disk has been deleted!
Build 'yandex' finished after 2 minutes 719 milliseconds.

==> Wait completed after 2 minutes 719 milliseconds

==> Builds finished. The artifacts of successful builds are:
--> yandex: A disk image was created: centos-7-base (id: fd848lduell9luffrd9g) with family name centos

```

3.Удаляем подсеть и сеть, которую использовали для сборки образа ОС.  
```commandline
ivan@ivan-virtual-machine:~/packer$ yc vpc subnet delete --name my-subnet-a
done (2s)
ivan@ivan-virtual-machine:~/packer$ yc vpc network delete net
ivan@ivan-virtual-machine:~/packer$ yc vpc subnet list
+----+------+------------+----------------+------+-------+
| ID | NAME | NETWORK ID | ROUTE TABLE ID | ZONE | RANGE |
+----+------+------------+----------------+------+-------+
+----+------+------------+----------------+------+-------+
```
4.Создаём 6 виртуальных машин с помощью Terraform.  
Включаем впн и проверяем terraform
```commandline
ivan@ivan-virtual-machine:~/terrraform$ ./terraform init

Initializing the backend...

Initializing provider plugins...
- Reusing previous version of yandex-cloud/yandex from the dependency lock file
- Using previously-installed yandex-cloud/yandex v0.81.0

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.

```
Создадим сервисный аккаунт 
```commandline
ivan@ivan-virtual-machine:~/terrraform$ yc config list 
token: y0_*****************************_vjY876EfTJU
cloud-id: b1grp****sfkb323pqkm
folder-id: b1g6****69utu4gsqs
compute-default-zone: ru-central1-a
ivan@ivan-virtual-machine:~/terrraform$ yc iam service-account --folder-id ^Cist
ivan@ivan-virtual-machine:~/terrraform$ yc iam service-account --folder-id b1g66****b69utu4gsqs list
+----+------+
| ID | NAME |
+----+------+
+----+------+

ivan@ivan-virtual-machine:~/terrraform$ yc iam service-account create --name sa-netology-terraform
id: ajeqab1snruj88ragr4s
folder_id: b1g66m6nub69utu4gsqs
created_at: "2022-11-14T20:02:36.060619011Z"
name: sa-netology-terraform


```
Затем в интерфейсе облака назначим ему роль editor

```commandline
ivan@ivan-virtual-machine:~/terrraform$ yc iam key create --service-account-id ajeqab1snruj88ragr4s --folder-name netology-swarm --output key_swarm.json 
id: ajef6k0****uung6p8g
service_account_id: aje***snruj88ragr4s
created_at: "2022-11-14T20:13:02.331500959Z"
key_algorithm: RSA_2048

ivan@ivan-virtual-machine:~/terrraform$ yc config profile create swarm-terraform
Profile 'swarm-terraform' created and activated
ivan@ivan-virtual-machine:~/terrraform$ yc config set service-account-key key_swarm.json 
ivan@ivan-virtual-machine:~/terrraform$ yc config set cloud-id b1grp5c***fkb323pqkm
ivan@ivan-virtual-machine:~/terrraform$ yc config set folder-id b1g66m6nub69utu4gsqs
ivan@ivan-virtual-machine:~/terrraform$ export YC_TOKEN=$(yc iam create-token)
ivan@ivan-virtual-machine:~/terrraform$ export YC_CLOUD_ID=$(yc config get cloud-id)
ivan@ivan-virtual-machine:~/terrraform$ export YC_FOLDER_ID=$(yc config get folder-id)

```
заполняем файлы конфигурации как в /home/ivan/terrraform/swarm
https://cloud.yandex.ru/docs/tutorials/infrastructure-management/terraform-quickstart#before-you-begin
если делаем по инструкции - то не надо указывать ключ сервисного акк в провайдере

```commandline
terraform plan
terraform apply
```
После использования можно применить terraform destroy и удалить все машины

5.Создаём Docker Swarm кластер из виртуальных машин, созданных на предыдущем шаге.  
ansible-playbook -i ./inventory ./swarm-deploy-cluster.yml 
6.Запускаем деплой стека приложений.  
sync playbook
7.Проводим стресс тест Docker Swarm кластера.  
stack playbook
8.Удаляем всё, чтобы не тратить деньги!
