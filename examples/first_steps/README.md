## Сссылки
* [OAuth2](https://shikimori.one/oauth)

## Базовая структура

```golang
package main

import (
  g "github.com/ghostemeow/goshikimori"
)

/*

ГАЙД ДЛЯ ЧАЙНИКОВ.

ЕСЛИ ТЫ ЗНАЕШЬ С ЧЕМ РАБОТАЕШЬ, ТО НЕ ТРАТЬ СВОЕ ВРЕМЯ И
ПЕРЕХОДИ СРАЗУ К ДОКУМЕНТАЦИИ.

--------------------------

Для получения полей APPLICATION_NAME и PRIVATE_KEY
перейди по ссылки вверху и создай приложение,
либо присоединись к "Test Api" и используй его, дело твое.

--------------------------

Для большинства запросов, например поиск
аниме/манги/пользователей и тому подобного,
поле APPLICATION_NAME будет достаточным, а
PRIVATE_KEY можно оставить пустым.

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "",
  )
}

--------------------------

Если нужно изменить какую-то информацию, а именно
добавить друга, добавить в игнор или написать
сообщение, то тогда обязательно поле PRIVATE_KEY.

ВАЖНО: у твоего приложения, при регистрации, должны
быть права доступа(что оно может делать от твоего имени)
для тех или иных операций.

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

*/

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}
```
