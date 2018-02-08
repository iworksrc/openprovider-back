# Тестовое задание для организации Openprovider

## Overview

Содержит **endpoint** вычисляющий n-й член [последовательности трибоначчи](https://oeis.org/A000073/list)

```http://```**servername:serverport**```/api/v1/openprovider/tribonachi/10```

## QuickStart (demo)
1. ```git clone https://github.com/iworksrc/openprovider-back```
2. ```cd openprovider-back```
3. ```go run main.go```
4. test RESTapi


## Docker

1. ```git clone https://github.com/iworksrc/openprovider-back```
2. ```cd openprovider-back```
3. ```./build.sh```
4. ```sudo docker run -p 8080:8080 -t openprovider-back openprovider-back:latest```
5. test RESTapi

## Tests
### Manual Test Endpoints
Тестировать endpoints можно программой [Postman](https://www.getpostman.com/apps).
Настройки endpoints для Postman находятся в файле *Openprovider.postman_collection.json*

### UnitTests

## OpenAPI
API version: 1.0.0

Swagger used for generation skels

