# Тестовое задание для организации Openprovider

# Task

Требуется реализовать RESTful-сервис, цель которого - вычисление [чисел трибоначчи](https://ru.wikipedia.org/wiki/Числа_трибоначчи).


* Сервис должен принимать во входном запросе натуральное число N (натуральное число – это любое целое число больше 0) и выдавать в ответе N-ное число трибоначчи. Например, для N = 10 ответ - 81.



* Вопросы конфигурирования веб-сервиса, дизайна API и валидации входного запроса предлагается решить тем способом, который будет наиболее подходящим для данной задачи.



* Сервис должен иметь возможность запуска из Docker-контейнера.



* К сервису необходимо приложить инструкцию по сборке и запуску.

* Обоснование выбранного алгоритма также должно присутствовать в документации.

* Форматирование и документирование кода должно быть выполнено в соответствии со стандартами языка Go.

* Приветствуется покрытие кода юнит- и/или функциональными тестами.


# Solution

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
4. ```./run_container.sh```
5. test RESTapi

## Tests
### Manual Test Endpoints
Тестировать endpoints можно программой [Postman](https://www.getpostman.com/apps).
Настройки endpoints для Postman находятся в файле *Openprovider.postman_collection.json*

### UnitTests

1. ```git clone https://github.com/iworksrc/openprovider-back```
2. ```cd openprovider-back/tests```
3. ```go test```

## Documentation

1. ```git clone https://github.com/iworksrc/openprovider-back```
2. ```cd openprovider-back```
3. ```./gen_docs.sh```
4. Open in browser ``docs/index.html``
 
## OpenAPI
API version: 1.0.0

Swagger used for generation skels

