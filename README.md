## Запуск проекта
Перед запуском проекта проверьте, что порт `8080` на `localhost` свободен, а дальше можете запустить его с помощью команды `go run cmd/main.go`


## О проекте

Имеется только одна схема `task`, которая выглядит следующим образом:
```
{
    "id": 4,
    "name": "task21",
    "description": "desc",
    "progress": 41,
    "createdAt": "2025-06-18T11:24:08.036891821+03:00",
    "workedSeconds": 25,
    "status": "failed"
}
```
При создании задачи случайным образом выбирается время её выполнения (не более 5 минут). Каждые 5 секунд вычисляется прогресс выполнения (в процентах). Если задачу отменяют, её прогресс останавливается, а статус меняется на `failed`. По истечении времени статус задачи становится `success` а прогресс `100`

## Роуты

- Листинг задач - GET `localhost:8080/tasks`.    
Запрос: `curl --location 'localhost:8080/tasks'`

- Создание задач - POST `localhost:8080/tasks` запрос принимает боди следующего типа: `{
    name": "task",
    "description?: "desc"
}` .   
Пример запроса: `curl --location 'localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "name": "task",
    "description": "desc"
}'`

- Отмена задачи - PATCH `localhost:8080/tasks/{taskId}/cancel`.  
Запрос:  `curl --location --request PATCH 'localhost:8080/tasks/4/cancel'`
- Получение задачи - GET `localhost:8080/tasks/{taskId}`.  
Запрос: `curl --location 'localhost:8080/tasks/1'`

