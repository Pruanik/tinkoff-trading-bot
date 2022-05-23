# Tinkoff Trading Bot (TTB)
### Торговый робот для автоматической торговкли

Торговый робот разработан в рамках конкурса [Tinkoff Invest Robot Contest](https://github.com/Tinkoff/invest-robot-contest) и предназначен для ведения автоматической торговли на бирже Tinkoff Invest, используя APIv2 на основе gRPC протокола.

[![Screenshot 1](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_1.png)]()

Основные преимущества данного проекта:
- ✨ Возможность пробовать различные настраиваемые стратегии. Любой алгоритм и его правила могут быть добавлены и выбраны в интерфейсе для использования.
- ✨ Интерфейс для отслеивания активности бота с визуализацией графиков и просмотров логов системы
- ✨ Все написано на быстром языке GoLang и упакованов контейнер для удобства использования

## Запуск приложения
Склонируется к себе репозиторий для локального использования:
```sh
git clone git@github.com:Pruanik/tinkoff-trading-bot.git && cd ./tinkoff-trading-bot
```

Скопируйте файл .env и заполните необходимые поля:
```sh
cp .env.example .env
```

Запустите сборку контейнеров необходимых для работы приложения:
```sh
make build-all
```

Запускаем миграции для создания структуры базы приложения:
```sh
make migration-up
```

После сборки, поднимем приложение командой:
```sh
make up
```

Веб интерфейс откроется на порту, который указан в .env файле в параметре WEB_APPLICATION_HOST_PORT

## UI приложения
Приложение состоит из нескольких страниц: 
- Dashboard с возможностью отслеживания и настройки процессов работы
- Profile с отображением актуальной информации об аккаунтах пользователя
- Logs с информацией из таблицы logs в которой фиксируются ключевые факторы работы системы
[![Screenshot 2](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_2.png)]()
[![Screenshot 3](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_3.png)]()

## Структура проекта
[![Screenshot 4](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_4.png)]()
Проект поделен на несколько контейнеров:
- WebApplication - все что связано с UI частью: веб сервер, frontend логика, api.
- TradingStrategy - запуск анализа накопленных данных для создания условий к покупке или продаже позиций. Может переодически запускать разные стратегии с разными параметрами на разные иснтрументы для работы с накопленными данными.
- TinkoffInvestConnection - все что связано с взаимодействием с Tinkoff Invest API: подгрузка данных в локальную базу, выставление рдеров и их отмена, отслеживание выполнения условий созданных модулем TradingStrategy

## Команды для работы с приложением
Команда для сборки всех контейнеров приложения
```sh
make build-all
```
Команда для сборки контейнера webapplication
```sh
make build-webapp
```
Команда для сборки контейнера migrator
```sh
make build-migrator
```
Команда для сборки контейнера frontend
```sh
make build-frontend
```
Команда для запуска всех базовых контейнеров приложения
```sh
make up
```
Команда для создания миграции "create_candles_table"
```sh
make migration-create create_candles_table
```
Команда для накатывания всех недостающих миграций
```sh
make migration-up
```
Команда для отката миграций
```sh
make migration-down
```
Команда для инициализации Frontend часть проекта
```sh
make frontend-install
```
Команда для сборки Frontend часть проекта
```sh
make frontend-build
```
Команда для отслеживания в реальном времени всех изменений в папке /web/src для пересборки Frontend часть проекта
```sh
make frontend-watch
```
Команда для запуска GoLang приложения на локальной машине
```sh
make run-webapp
```
Команда для запуска GoLang приложения на локальной машине
```sh
make run-tinkoffconnection
```