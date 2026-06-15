## [EN](https://github.com/ghostemeow/goshikimori/blob/master/README_en.md) | RU

### О проекте
Небольшая библиотека для взаимодействия с шикимори, написанная на языке golang.
* Работа с API происходит через `OAuth2`. \
И начать нужно с ознакомления документации [первые шаги](https://github.com/ghostemeow/goshikimori/blob/master/examples/first_steps/README.md).
* Никаких зависимостей от других библиотек.
* Для тестов и сборки используется утилита [GNU make](https://www.gnu.org/software/make/manual/make.html).

### Установка
```bash
go get github.com/ghostemeow/goshikimori
```

### Готовые примеры
* [Нажать сюда](https://github.com/ghostemeow/goshikimori/tree/master/examples)

### Документация GODOC / GOPKG
Также доступна поддержка **godoc**.

Или вы можете использовать страницу с официального
[сайта](https://pkg.go.dev/github.com/ghostemeow/goshikimori) Go pkg.\
**P.S.** документация обновляется с опозданием.
```bash
# Способ #1: Используя докер.
git clone git@github.com:ghostemeow/goshikimori.git && cd goshikimori
make docker

# Открыть в браузере.
http://localhost:1337/pkg/github.com/ghostemeow/goshikimori/
```
```bash
# Способ #2(Linux): Установка godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Добавить 'экспорт' в файл /home/$USER/.profile и перезагружаемся.
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Проверяем работоспособность.
godoc --help

# После установки или если 'godoc' уже установлен.
git clone git@github.com:ghostemeow/goshikimori.git && cd goshikimori
make doc

# Открыть в браузере.
http://localhost:1337/pkg/github.com/ghostemeow/goshikimori
```

### Документация шикимори
* [GraphQL](https://shikimori.io/api/doc/graphql)
* [API v1](https://shikimori.io/api/doc/1.0)
* [API v2](https://shikimori.io/api/doc/2.0)

## Обратная связь
* Написать в личные сообщения на [сайте](https://shikimori.io/arctica).
* Открыть [проблему](https://github.com/ghostemeow/goshikimori/issues).
