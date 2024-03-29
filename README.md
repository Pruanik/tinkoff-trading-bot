# Tinkoff Trading Bot (TTB)
### Торговый робот для автоматической торговли

Торговый робот разработан в рамках конкурса [Tinkoff Invest Robot Contest](https://github.com/Tinkoff/invest-robot-contest) и предназначен для ведения автоматической торговли на бирже Tinkoff Invest, используя APIv2 на основе gRPC протокола.

**ВНИМАНИЕ!** Разработчик и автор проекта не несет никакой ответственности за любые финансовые потери или другой ущерб, которые могут произойти при использовании данного торгового робота. Использование данного продукта осуществляется на страх и риск самого пользователя (лица осуществляющего запуск программы).

[![Screenshot 1](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_1.png)]()

Основные преимущества данного проекта:
- ✨ Возможность пробовать различные настраиваемые стратегии. Любой алгоритм и его правила могут быть добавлены и выбраны в интерфейсе для использования.
- ✨ Интерфейс для отслеживания активности бота с визуализацией графиков и просмотров логов системы
- ✨ Все написано на быстром языке GoLang и упаковано в контейнер для удобства использования

## Запуск приложения
Склонируйте к себе репозиторий для локального использования:
```sh
git clone git@github.com:Pruanik/tinkoff-trading-bot.git && cd ./tinkoff-trading-bot
```

Скопируйте файл .env и заполните необходимые поля:
```sh
cp .env.example .env
```

Укажите в .env файле режим торговли и Ваш токен авторизации в полях:
```sh
TRADING_MOD
PRODUCTION_TOKEN
SANDBOX_TOKEN
```

Запустите сборку контейнеров необходимых для работы приложения:
```sh
make build-all
```

Поднимаем пока что только контейнер с базой, для того чтобы накатить миграции до того как начнет работать приложение:
```sh
make up-postgresql
```

Запускаем миграции для создания структуры базы приложения:
```sh
make migration-up
```

После сборки, поднимем приложение командой (в дальнейшем достаточно запускать только последнюю команду для использования готового приложения):
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
Расшифровка всех ошибок находится на [сайте документации](https://tinkoff.github.io/investAPI/errors/)

## Структура проекта
[![Screenshot 4](https://github.com/Pruanik/tinkoff-trading-bot/raw/master/screenshots/screenshot_4.png)]()
Проект поделен на несколько контейнеров:
- WebApplication - все что связано с UI частью: веб сервер, frontend логика, api.
- TradingStrategy - запуск анализа накопленных данных для создания "Условий" к покупке или продаже позиций. Может периодически запускать разные стратегии с разными параметрами на разные инструменты для работы с накопленными данными.
- TinkoffInvestConnection - все что связано с взаимодействием с Tinkoff Invest API: подгрузка данных в локальную базу, выставление ордеров и их отмена, отслеживание выполнения условий созданных модулем TradingStrategy.

Основное преимущество данного проекта, возможность добавлять и управлять алгоритмами в любом их виде. Это может быть дополнительный контейнер с программным кодом написанным на любом языке или добавление новой логики в текущий контейнер TradingStrategy. Условиями для корректной работы является:
- Добавить новый алгоритм в таблицу trading_strategies
- Описать все необходимые условия для работы этого алгоритма, которое в следствии будут заполняться пользователем в UI интерфейсе и сохранятся в json.
- Сохранять результат работы алгоритма в качестве "Условия" в таблицу conditions для последующей обработки этих значений в контейнере TinkoffInvestConnection при получении значений через Bidirectional-stream.

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
> К сожалению, на данный момент, не весь функционал реализован в полной мере.
> Пока написан этот комментарий, не полностью реализована система Bidirectional-stream, есть проблема с графиками в интерфейсе
> система формирования и отслеживания "Условий". Также планируется добавить вывод информации
> об состоянии аккаунтов в каждом из режимов торговли, возможность управлением аккаунтами в sandbox режиме и отображения списка сделок по покупкам с подсчетом прибыльности запущенных алгоритмов.