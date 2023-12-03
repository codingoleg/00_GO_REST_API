# REST API

+ http://localhost:8000/api/order/ [POST] \
Отправляет информацию о заказе: id пользователя и описание
```bash
curl --location 'http://localhost:8000/api/order/' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "name": "Описание"
}'
```

+ http://localhost:8000/api/order/ [GET] \
  Получает информацию о заказе, где id заказа = 1
```bash
curl --location 'http://localhost:8000/api/order/1' \
--data ''
```

+ http://localhost:8000/api/order/ [GET] \
  Получает информацию о всех заказах, где id пользователя = 1 
```bash
curl --location 'http://localhost:8000/api/user/1'
```