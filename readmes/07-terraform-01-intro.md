1. Legend  
**Ответ:**
>   Какой тип инфраструктуры будем использовать для этого проекта: изменяемый или не изменяемый?
 
Оба варианта. Систему нужно развернуть + частые релизы, доработки.
>   Будет ли центральный сервер для управления инфраструктурой?

Центральный сервер не требуется 
>   Будут ли агенты на серверах?
 
Агенты не требуются.
>   Будут ли использованы средства для управления конфигурацией или инициализации ресурсов?

Да. Terraform будем использовать как средство инициализации ресурсов. Ansible - управление конфигурацией.
>Какие инструменты из уже используемых вы хотели бы использовать для нового проекта?  

Terraform будем использовать как средство инициализации ресурсов. Ansible - управление конфигурацией. Packer - для создания образов ВМ,
Docker - образы контейнеров. Bash скрипты возможно придётся использовать.
Teamcity тоже используем для CI/CD + автоматических тестов.

> Хотите ли рассмотреть возможность внедрения новых инструментов для этого проекта?

Рассмотрел бы внедрение GIT для хранения конфигураций.

2. Terraform
```commandline
[ifebres@febres-i whoerconfigs]$ sudo dnf -y install terraform

[ifebres@febres-i whoerconfigs]$ terraform --version
Terraform v1.3.7
on linux_amd64
```

3. Old version terraform

```commandline
[ifebres@febres-i terraform_12]$ ./terraform --version
Terraform v1.2.0
on linux_386

Your version of Terraform is out of date! The latest version
is 1.3.7. You can update by downloading from https://www.terraform.io/downloads.html


[ifebres@febres-i terraform_12]$ terraform --version
Terraform v1.3.7
on linux_amd64
```