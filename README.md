# Виртуалка

ip: TBA

password: TBA

login: ваш логин на гитлабе


# Хакатон по курсу golang

Пишем биржу :)

Биржа состоит из 3-х компонентов:
- Клиент
- Брокер
- Биржа

## Биржа

Биржа - центральная точка всей системы. Она сводит между собой различных продавцов и покупателей из разных брокеров, результатом сделок которых является изменение цены.

В нашем хакатоне будут использоваться исторические данные реальных торгов по фьючерсным контрактам на бирже РТС за 2019-05-17 - https://cloud.mail.ru/public/2qHq/pd99CWTLh

Формат данных `<TICKER>;<PER>;<DATE>;<TIME>;<LAST>;<VOL>`:
* TICKER - название торгуемого инструмента
* PER - период, у нас тики ( отдельные сделки ), игнорируйте это поле
* DATE - дата
* TIME - время
* LAST - цена прошедшей сделки
* VOL - объём продшей сделки

Если на биржу ставится заявка на покупку или продажу, то она ставится в очередь и когда цена доходит до неё и хватает объёма - заявка исполняется, брокеру уходит соответствующее уведомление. Если не хватает объёма, то заявка исполняется частичсно, брокеру так же уходит уведомление.
Если несколько участников поставилос заявку на одинаковый уровеньт цены, то исполняются в порядке добавления.

Внутри биржи цена образует книгу заявок, то что называется стаканом - https://s.mail.ru/ESbV/M5iGMvkQR - это на каком уровне стоят заявки на покупку или продажу

В связи с тем что наши данные исторические - предполагем, что мы не можем выступить инициатором сделки (т.е. сдвинуть цену, купив по рынку), можем приобрести только если другая сторона выступила инциатором. Это значит что мы встаём в стакан и когда цена доходит до нас - происходит сделка.

Заявки могут быть 2 видов:
* на прокупку - "я хочу купить по цене Х" - когда цена сверху вниз доходит до нашей заявки - она исполняется
* на продажу - "я хочу продать по цене Х" - когда цена снизу вверх доходит до нашей заявки - она исполняется
* визуально это выглядит так https://s.mail.ru/4ikh/Bjmu8zPiv 

Помимо исполнения сделок брокер транслирует цену инструментов всем подключенным брокерам. Список транслируемых инструментов берётся из конфига, сами цены - из файла.
В связи с тем что цены исторические - мы не смотрим на дату, а просто начинаем таранслировать ту цену что есть. Инфомрация о изменении цены отправляется каждую секунду. Если в секунду ( под конфигом) произошло больше чем 1 сделка - они аггрегируются. Брокеру отправляется OHLCV ( open, high, low, close, volume ), где:
- open - цена открытия интервала (первая сделка)
- high - максимальная цена в интервале
- low - минимальная цена в интервале
- close - цена закрытия интервала ( последняя сделка )
- volume - количество проторгованных контрактов

Формат обмена данными с брокером - protobuf через GRPC

```protobuf
syntax = "proto3";

message OHLCV {
  int64 ID = 1; // внутренний идентификатор, просто авто-инкремент
  int32 Time = 2;
  int32 Interval = 3; // в данном случае - 1 секунда
  float Open = 4;
  float High = 5;
  float Low = 6;
  float Close = 7;
  string Ticker = 8;
}

message Deal {
    int64 ID = 1; // DealID который вернулся вам при простановке заявки
    int32 BrokerID = 2;
    int32 ClientID = 3;
    string Ticker = 4;
    int32 Amount = 5; // сколько купили-продали
    bool Partial = 6; // флаг что сделка клиента исполнилсь частично
    int32 Time = 7;
    float Price = 8;
}

message DealID {
    int64 ID = 1;
    int64 BrokerID = 2;
}

message BrokerID {
    int64 ID = 1;
}

message CancelResult {
    bool success = 1;
}

service Exchange {
    // поток ценовых данных от биржи к брокеру
    // мы каждую секнуду будем получать отсюда событие с ценами, которые броке аггрегирует у себя в минуты и показывает клиентам
    // устанавливается 1 раз брокером
    rpc Statistic (BrokerID) returns (stream OHLCV) {}

    // отправка на биржу заявки от брокера
    rpc Create (Deal) returns (DealID) {}

    // отмена заявки
    rpc Cancel (DealID) returns (CancelResult) {}

    // исполнение заявок от биржи к брокеру
    // устанавливается 1 раз брокером и при исполнении какой-то заявки 
    rpc Results (BrokerID) returns (stream Deal) {}
}


```

## Брокер

Брокер - это организация, которая предсотавляет своим клиентам доступ на биржу.
У неё есть список клиентов, которые могут взаимодействовать посредством неё с биржей, так же она хранит количество их позиицй и историю сделок.

Брокер аггрегирует внутри себя информацию от биржи по ценовым данным, позволяя клиенту посмотреть историю. По-умолчанию, хранится история за последнеи 5 минут (300 секунд).

! смотрите сначала скрин у клиента

Брокер предоставляет клиентам JSON-апи (REST или JSON-RPC) или же grpc-апи ( отдельный proto-файл от того что у биржи ), через который им доступныы следующие возможности:
* посмотреть свои позиции и баланс - возвращает баланс + список заявок ( слайс структур ), может быть преобразовано в таблицу на хтмл
```json
-> GET /api/v1/status

<-
{
    "body": {
        "balance": 10000000,
        "positions": [
            {"ticker": "SPFB.RTS", ... }
        ],
        "open_orders": [
            {"id": 123, "ticker": "SPFB.RTS", ... }
        ]
    }
}
```

* отправить на биржу заявку на покупку или продажу тикера ( то что вы видите на скрине у клиента )
```json
-> POST /api/v1/deal
{
    "deal": {
        "ticker": "SPFB.RTS",
        "type": "BUY",
        "amout": 100,
        "price": 11
    }
}

<-
{
    "body": {
        "id": "123"
    }
}
```

* отменить ранее отправленную заявку - принимает ИД заявки
```json
-> POST /api/v1/cancel
{
    "id": 123
}

<-
{
    "body": {
        "id": "123",
        "status": ...
    }
}
```

* посмотреть последнюю истории торгов - возвращает слайс структур, может быть преобразовано в таблицу на хтмл
```json
-> /api/v1/history?ticker=SPFB.RTS

<-
{
    "body": {
        "ticker": "SPFB.RTS",
        "prices": [
            {"open": ...},
            ...
        ]
    }
}
```

```sql
CREATE TABLE `clients` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `login_id` int NOT NULL,
    -- `password` varchar(300) NOT NULL,
    `balance` int NOT NULL
);

-- INSERT INTO `clients` (`id`, `login`,  `password`, `balance`) 
--     VALUES (1, 'Vasily', '123456', 200000),
--     VALUES (2, 'Ivan', 'qwerty', 200000),
--     VALUES (3, 'Olga', '1qaz2wsx', 200000);

CREATE TABLE `positions` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` int NOT NULL,
    `ticker` varchar(300) NOT NULL,
    `vol` int NOT NULL,
    KEY user_id(user_id)
);

-- INSERT INTO `clients` (`user_id`, `ticker`, `amount`) 
--     VALUES (1, 'SiM7', '123456', 200000),
--     VALUES (1, 'RIM7', '123456', 200000),
--     VALUES (2, 'RIM7', 'qwerty', 200000);
    
CREATE TABLE `orders_history` (
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `time` int NOT NULL,
    `user_id` int,
    `ticker` varchar(300) NOT NULL,
    `vol` int NOT NULL,
    `price` float not null,
    `is_buy` int not null,
    KEY user_id(user_id)
);

CREATE TABLE `request` ( -- запросы
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` int,
    `ticker` varchar(300) NOT NULL,
    `vol` int NOT NULL,
    `price` float NOT NULL,
    `is_buy` int not null, -- 1 - покупаем, 0 - продаем
    KEY user_id(user_id)
);


CREATE TABLE `stat` ( -- запросы
    `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `time` int,
    `interval` int,
    `open` float,
    `high` float,
    `low` float,
    `close` float,
    `volume` int,
    `ticker` varchar(300),
    KEY id(id)
);
```

## Клиент

Клиент - это любой пользователь АПИ брокера. Это может быть биржевой терминал, веб-сайт, торговый робот.

Мы делаем простой биржевой терминал на хтмл, который нам позволяет покупать-продавать и смотреть свою историю.

https://s.mail.ru/Bszh/cxrEKBPqD ( там 3 разные вкладки )

Не заморачивайтесь с интерфейсом!

* Таблицы уже выводились в примерах курса, скопируйте код оттуда и подставьте свои поля
* Формы тоже были
* По кнопке "отменить" на биржу уходит запрос Cancel черех брокера
* не обязательно повторять интерфейс 1 в 1, достаточно чтобы оно просто работало

## Организация работы

Обсуждаем проект, делимся на команды, пишем код. 3 компонента, 3 команды. Работам короткими этапами по 45 минут, стараясь в конце этапа иметь логически завершенный кусок.

Каждая команда мерджит свой компонент в мастер, если считает что всё будет рабоать.

Проект делать через го модули и со структурой как в реддите.

Биржа - самый сложный компонент. Рекомендую его сначала замокать, чтобы выдавал "ок", а потом добивать.

Прежде чем начать яростно кодить - вспомните предыдущий хакатон. Как там всё прошло, какие были затыки и что можно сделать лучше.

------------------------------------

порядок работы:

0. обсудить задачу в целом
1. выбрать координаторов, которые будут отвечать за взаимодействие между командами и общий успех
2. разбиться на команды
3. обсудить задачу команды
4. обсудить протоколы общения с взаимодействующими компонентами
5. начинать писать код

внимание! 

* если не будет координатора - будет хаотичная разработка и в итоге ничего не заработает
* если вы не договоритесь о форматах и прочем, а сразу же броситесь писать код - в итоге выяснится, что компоненты между собой не совместимы и ничего не работает

!!!
По опыту прошлого хакатона вы могли заметить что:
* если начать интегрироваться в самом конце - выяснится много нестыковок. лучше начинать интеграцию на заглушках как можно раньше
!!!

проверено :)

писать код лучше в режиме парного-тройного программирования, периодически меняясь у клавиатуры

-----

так же вам надо будет заполнить отчет по результатам хакатона и ответить на следущие вопросы

1. над каким компонентом вы работали, что именно вы писали
2. кто работал вместе с вами в команде
3. с какими проблемами/трудностями вы встрелились (любого характера)
4. какие знания вам пригодились, какие знания вы применяли
5. каких знаний вам не хватило и пришлось быстро гуглить, читатть доку, разбираться
6. что получилось, почему?
7. что не получилось, почему?
8. что лично вы могли бы сделать лучше в следующий раз
9. что вы, как команда, могли бы сделать лучше в следующий раз

отчет закоммитить в свой репозиторий (не в хакатоновский!) в формате md ( это обычный текстовый файл ) с именем `hakaton2.md`

Без отчета баллы не выставляется
