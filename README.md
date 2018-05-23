# Teleport server

[![Build Status](https://travis-ci.com/iMega/teleport-server.svg?token=DhjjgmgtJp2pAr6izscn&branch=master)](https://travis-ci.com/iMega/teleport-server)

### Запуск тестов

make test

### Обновление зависимостей

dep ensure

dep status

### Фиксировать версии пакетов зависимостей

teleport-server/Gopkg.toml


### Изменения в схеме

teleport-server/schema/schema.graphql

make generate


### Изменения в API

teleport-server/api/service.proto

make proto

