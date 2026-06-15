## [EN](https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL_en.md) | RU

### На данный момент API GraphQL заявлен как экспериментальный.

Далее рассмотрим на примерах:
```golang
package main

import (
  "fmt"

  shiki "github.com/ghostemeow/goshikimori"
)

func conf() *shiki.Configuration {
  return shiki.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "score", "episodes", "airedOn{year month day date}".
  // Вторым параметром идет название аниме; name: "initial d".
  // Теперь переходим к интерфейсу:
  //    1)  page: 1;
  //    2)  limit: 5;
  //    3)  score: 8;
  //    4)  order: ""; пропустил;
  //    5)  kind: ANIME_KIND_TV;
  //    6)  status: ANIME_STATUS_RELEASED;
  //    7)  season: ""; пропустил;
  //    8)  duration: ""; пропустил;
  //    9)  rating: ANIME_RATING_PG_13;
  //    10) mylist: ""; пропустил;
  //    11) censored: false;
  //    12) genre: nil; пропустил;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.AnimeSchema();
  schema, err := shiki.AnimeSchema(
    shiki.ValuesSchema("id", "name", "score", "episodes", "airedOn{year month day date}"),
    "initial d", 1, 5, 8, "", shiki.ANIME_KIND_TV,
    shiki.ANIME_STATUS_RELEASED, "", "", shiki.ANIME_RATING_PG_13, "", false, nil,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  a, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(a.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range a.Data.Animes {
    fmt.Println(
      v.Id, v.Name, v.Score, v.Episodes, v.AiredOn.Year,
      v.AiredOn.Month, v.AiredOn.Day, v.AiredOn.Date,
    )
  }
}
```
```golang
package main

import (
  "fmt"

  shiki "github.com/ghostemeow/goshikimori"
)

func conf() *shiki.Configuration {
  return shiki.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "score", "volumes", "chapters", "releasedOn{year}".
  // Вторым параметром идет название манги; name: "initial d".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 1;
  //    3) score: 8;
  //    4) order: ""; пропустил;
  //    5) kind: MANGA_KIND_MANGA;
  //    6) status: MANGA_STATUS_RELEASED;
  //    7) season: ""; пропустил;
  //    8) mylist: MY_LIST_COMPLETED;
  //    9) censored: false;
  //    10) genre: nil; пропустил;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.MangaSchema();
  schema, err := shiki.MangaSchema(
    shiki.ValuesSchema("id", "name", "score", "volumes", "chapters", "releasedOn{year}"),
    "liar game", 1, 1, 8, "", shiki.MANGA_KIND_MANGA, shiki.MANGA_STATUS_RELEASED,
    "", shiki.MY_LIST_COMPLETED, false, nil,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  m, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(m.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range m.Data.Mangas {
    fmt.Println(v.Id, v.Name, v.Score, v.Volumes, v.Chapters, v.ReleasedOn.Year)
  }
}
```
```golang
package main

import (
  "fmt"

  shiki "github.com/ghostemeow/goshikimori"
)

func conf() *shiki.Configuration {
  return shiki.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "russian", "url", "description"".
  // Вторым параметром идет название персонажа; name: "onizuka".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 2;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.CharacterSchema();
  schema, err := shiki.CharacterSchema(
    shiki.ValuesSchema("id", "name", "russian", "url", "description"),
    "onizuka", 1, 2,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  ch, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(ch.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range ch.Data.Characters {
    fmt.Println(v.Id, v.Name, v.Russian, v.Url, v.Description)
  }
}
```
```golang
package main

import (
  "fmt"

  shiki "github.com/ghostemeow/goshikimori"
)

func conf() *shiki.Configuration {
  return shiki.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := conf()

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "name", "russian", "url", "website", "birthOn{year month day date}".
  // Вторым параметром идет имя человека; name: "satsuki".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) isSeyu: true;
  //    4) isMangaka: false;
  //    5) isProducer: false;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.PeopleSchema();
  schema, err := shiki.PeopleSchema(
    shiki.ValuesSchema("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
    "satsuki", 1, 1, true, false, false,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  p, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(p.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range p.Data.People {
    fmt.Println(
      v.Id, v.Name, v.Russian, v.Url, v.Website,
      v.BirthOn.Year, v.BirthOn.Month, v.BirthOn.Day, v.BirthOn.Date,
    )
  }
}
```
```golang
package main

import (
  "fmt"

  shiki "github.com/ghostemeow/goshikimori"
)

func config() *shiki.Configuration {
  return shiki.SetConfiguration(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func main() {
  c := config()

  // Первым параметром идет перечисление значений которые мы хотим получить
  // от сервера; values: "id", "text", "score", "createdAt", "anime{id name}".
  // Вторым параметром идет id пользователя; userId: 181833.
  // Третьим параметром заводим вспомогательную функцию, которая разобьет два
  // дополнительных поля: "order: {field: id, order: desc}".
  // Теперь переходим к интерфейсу:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) status: completed;
  //    4) targetType: Anime;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.UserRatesSchema();
  schema, err := shiki.UserRatesSchema(
    shiki.ValuesSchema("id", "text", "score", "createdAt", "anime{id name}"),
    181833, graph.UserRatesOrder(shiki.GRAPHQL_ORDER_FIELD_ID, shiki.GRAPHQL_ORDER_ORDER_DESC),
    1, 10, shiki.MY_LIST_COMPLETED, shiki.TARGET_TYPE_ANIME,
  )
  if err != nil {
    fmt.Println(err)
    return
  }

  ur, status, err := c.SearchGraphql(schema)
  if status != 200 || err != nil {
    fmt.Println(err)
    return
  }

  // Тут можно отслеживать ошибки полученные при ответе сервера.
  fmt.Println(ur.Errors)
  // Стандартный вывод нашего поиска, ничего нового.
  for _, v := range ur.Data.UserRates {
    fmt.Println(v.Id, v.Text, v.Score, v.CreatedAt, v.Anime.Id, v.Anime.Name)
  }
}
```
