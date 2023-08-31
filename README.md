# Тестовое задание для стажёра Backend
## Сервис динамического сегментирования пользователей

### Проблема

В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие.
На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.

### Задача

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

### Запуск приложения

```
make build
make run
```
Если приложение запускается впервые, то необходимо применить миграции к базе данных:
```
make migrate-up
```

### Описание методов

#### Добавление сегмента

Пример запроса:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/9f88ea40-5d0d-40da-ac17-378f3bcbb957)

Пример ответа:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/4f926f8f-3b04-440c-b8d5-d352a6209675)

#### Удаление сегмента

Пример запроса:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/64683a4c-9514-4430-8dd9-2273b835c8db)

Пример ответа:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/5d77d6e1-21d2-4c82-8990-b1a7e693d861)

#### Добавление и удаление сегментов для пользователя

Пример запроса:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/6d7a1586-7833-4a95-ae28-6f63ee1eef59)

Пример ответа:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/16e3b2dc-0770-44e8-9fa4-c92aedb33b95)

#### Получение сегментов пользователя

Пример запроса:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/2325c917-5040-4fa1-a4df-46d2e5111ec6)

Пример ответа:

![image](https://github.com/DenChika/backend-trainee-assignment-2023/assets/79001610/58b0faf4-a2eb-455b-a168-4ea5fab503c3)

### Swagger

Swagger файл описан в директории [docs](docs)

При работающем приложении можно посмотреть [здесь](http://localhost:8080/swagger/index.html)

### Схема бд

Схема бд описана в [файле первой миграции](migrations/000001_init.up.sql)

### Неопределённости по тз

#### Нет требований по тому, какие поля запросов обязательные, а какие нет

Решение: 
* В методах [добавления](#добавление-сегмента) и [удаления](#удаление-сегмента) сегментов поля обязательные
* В остальных методах обязательным является только поле **user_id**

#### Что делать, если в методах [добавления и удаления сегментов для пользователя](#добавление-и-удаление-сегментов-для-пользователя) не все имена сегментов возможно добавить или удалить

Может быть несколько сценариев, почему сегмент не может быть добавлен пользователю или удалён у пользователя:
* Сегмента с такими именем не существует
* Пользователь уже находится в сегменте, который мы хотели добавить
* Пользователь не содержится в сегменте, который мы хотели удалить

Решение:
Все невалидные сегменты игнорируются, и запрос проходит. В ответе запроса содержатся названия только тех сегментов, которые были успешно добавлены/удалены
