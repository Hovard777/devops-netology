### Задача 1

*Электронные чеки в json виде* - NoSQL (key-value)  
*Склады и автомобильные дороги для логистической компании* - Графовые или сетевые. Т.к. надо выстроить отношение объектов  
*Генеалогические деревья* - NoSQL(Иерархические),т.к. необходимо Иерархическое представлние.  
*Кэш идентификаторов клиентов с ограниченным временем жизни для движка аутенфикации*  - column-oriented. Т.к. требуется обрабатывать пары значений и возможно большие данные.  
*Отношения клиент-покупка для интернет-магазина* - Реляционные СУБД. Нужно хранить данные в нескольких связанных таблицах  


### Задача 2


*Данные записываются на все узлы с задержкой до часа (асинхронная запись)* - CA, EL-PC  
*При сетевых сбоях, система может разделиться на 2 раздельных кластера* - AP, PA-EL  
*Система может не прислать корректный ответ или сбросить соединение* - CP, PA-EC  


### Задача 3

Принцыпы BASE и ACID сочетаться не могут. По ACID - данные согласованные, а по BASE - могут быть неверные - они противоречат друг другу.
### Задача 4

Вам дали задачу написать системное решение, основой которого бы послужили:

    фиксация некоторых значений с временем жизни
    реакция на истечение таймаута

Вы слышали о key-value хранилище, которое имеет механизм Pub/Sub. Что это за система? Какие минусы выбора данной системы?

#### Ответ.
Redis  
    Минусы Redis:
1. Если Redis падает - все данные из памяти теряются. Redis - in memmory хранилище. Но есть RDB(снапшот) и AOF(лог всех операций).
Обычно RDB выполняется по раписанию - поэтому можно потерять небольшую часть данных. Требует создания форка(дочерний процесс). Нужно скопировать всю память и если база большая то потребуется время для порождения форка. Редис форкается и он однопоточный - все пользоваттели бд ждут завершения форка.
AOF - логфайл достаточно большой, нужно следить за местом на диске и настраивать ротацию/перезапись. Если большая нагрузка чтения-записи - пишется медленнее чем RDB.
2. Redis недоступен все время, пока он поднимает данные с диска (это логично, конечно), т.е. если вдруг у вас упал Redis-процесс, то до тех пор, пока вновь поднятый сервис не поднимет все свое состояние с диска, сервис отвечать не будет, даже если для вас отсутствие данных гораздо меньшая проблема чем недоступная БД
3. Если используете хранение данных - следите за местом. Надо иметь запас минимум x2 от текущего размера данных в памяти иначе, в момент форка (неважно rdb или aof)  закончится место на диске, потому как в этот момент создается temp файл, куда скидываются данные и только после этого удаляется старый файл.
