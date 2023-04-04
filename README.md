## Проект на Golang Аналог Zabbix
реализовать простой аналог Microsoft SCOM или Zabbix:
- Получение списка запущенных процессов в системе, получение дополнительной
информации от них (pid, владелец, сколько времени работают и т.д.)
- Управление этими процессами: запуск, остановка и т.д.
- Выполнение произвольной команды (запуск простого bash скрипта)
- Сбор, экспорт ряда характеристик системы (всевозможные данные для инвентаризации)

### доп зависимости go
``` bash
go get github.com/labstack/echo/v4  # Веб сервер
go get github.com/shirou/gopsutil/v3  # Для работы с процессами и прочим
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/joho/godotenv
```

## Запуск
```
$ go run main.go
frontend $ npm run serve
```