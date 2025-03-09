## Overview

The repository of warehouse service

## Endpoints

Method | Path                             | Description                                   |                                                                         
---    |----------------------------------|------------------------------------------------
GET    | `/health`                        | Health page                                   |
GET    | `/metrics`                       | Страница с метриками                          |
GET    | `/v1/warehouses/list`            | Получение всех складов                        |
GET    | `/v1/warehouses/get/{userId}`    | Получение всех товаров id склада              |
POST   | `/v1/warehouses/book-item`       | Забронировать товар на складе                 |
POST   | `/v1/warehouses/rollback-book`   | Снять бронь с товара на складе                |