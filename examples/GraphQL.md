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
  // Третьим параметром идет структура AnimeOptions; нулевое значение поля означает,
  // что параметр пропускается:
  //    Page: 1;
  //    Limit: 5;
  //    Score: 8;
  //    Kind: ANIME_KIND_TV;
  //    Status: ANIME_STATUS_RELEASED;
  //    Rating: ANIME_RATING_PG_13;
  //    Censored: false;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.AnimeSchema();
  schema := shiki.AnimeSchema(
    shiki.ValuesSchema("id", "name", "score", "episodes", "airedOn{year month day date}"),
    "initial d",
    shiki.AnimeOptions{
      Page: 1, Limit: 5, Score: 8, Kind: shiki.ANIME_KIND_TV,
      Status: shiki.ANIME_STATUS_RELEASED, Rating: shiki.ANIME_RATING_PG_13,
    },
  )

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
  // Вторым параметром идет название манги; name: "liar game".
  // Третьим параметром идет структура MangaOptions; нулевое значение поля означает,
  // что параметр пропускается:
  //    Page: 1;
  //    Limit: 1;
  //    Score: 8;
  //    Kind: MANGA_KIND_MANGA;
  //    Status: MANGA_STATUS_RELEASED;
  //    Mylist: MY_LIST_COMPLETED;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.MangaSchema();
  schema := shiki.MangaSchema(
    shiki.ValuesSchema("id", "name", "score", "volumes", "chapters", "releasedOn{year}"),
    "liar game",
    shiki.MangaOptions{
      Page: 1, Limit: 1, Score: 8, Kind: shiki.MANGA_KIND_MANGA,
      Status: shiki.MANGA_STATUS_RELEASED, Mylist: shiki.MY_LIST_COMPLETED,
    },
  )

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
  // Третьим параметром идет структура CharacterOptions; нулевое значение поля
  // означает, что параметр пропускается:
  //    Page: 1;
  //    Limit: 2;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.CharacterSchema();
  schema := shiki.CharacterSchema(
    shiki.ValuesSchema("id", "name", "russian", "url", "description"),
    "onizuka",
    shiki.CharacterOptions{Page: 1, Limit: 2},
  )

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
  //    2) limit: 1;
  //    3) isSeyu: true;
  //    4) isMangaka: false;
  //    5) isProducer: false;
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.PeopleSchema();
  schema := shiki.PeopleSchema(
    shiki.ValuesSchema("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
    "satsuki",
    shiki.PeopleOptions{Page: 1, Limit: 1, IsSeyu: true},
  )

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
  // Третьим параметром идет структура UserRatesOptions; нулевое значение поля
  // означает, что параметр пропускается:
  //    Page: 1;
  //    Limit: 10;
  //    Status: MY_LIST_COMPLETED;
  //    TargetType: TARGET_TYPE_ANIME;
  //    Order: UserRatesOrder() - вспомогательная функция, которая разобьет два
  //           дополнительных поля: "order: {field: id, order: desc}".
  //
  // Про доступные значения можно почитать в описании функции: shiki.ValuesSchema();
  // Про доступные параметры интерфейса можно почитать в описании функции: shiki.UserRatesSchema();
  schema := shiki.UserRatesSchema(
    shiki.ValuesSchema("id", "text", "score", "createdAt", "anime{id name}"),
    181833,
    shiki.UserRatesOptions{
      Page: 1, Limit: 10, Status: shiki.MY_LIST_COMPLETED,
      TargetType: shiki.TARGET_TYPE_ANIME,
      Order: shiki.UserRatesOrder(shiki.GRAPHQL_ORDER_FIELD_ID, shiki.GRAPHQL_ORDER_ORDER_DESC),
    },
  )

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
