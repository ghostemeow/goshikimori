## EN | [RU](https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md)

## At the moment, the GraphQL API is stated as experimental.

Next, let's look at an examples:
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

  // The first parameter is the values of the anime; values: "id", "name", "score", "episodes", "airedOn{year month day date}".
  // The second parameter is the name of the anime; name: "initial d".
  // Now let's move on to the interface:
  //    1)  page: 1;
  //    2)  limit: 5;
  //    3)  score: 8;
  //    4)  order: ""; skipped;
  //    5)  kind: ANIME_KIND_TV;
  //    6)  status: ANIME_STATUS_RELEASED;
  //    7)  season: ""; skipped;
  //    8)  duration: ""; skipped;
  //    9)  rating: ANIME_RATING_PG_13;
  //    10) mylist: ""; skipped;
  //    11) censored: false;
  //    12) genre: nil; skipped
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.AnimeSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(a.Errors)
  // Standard output of our search, nothing new.
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

  // The first parameter is the values of the manga; values: "id", "name", "score", "volumes", "chapters", "releasedOn{year}".
  // The second parameter is the name of the manga; name: "initial d".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 1;
  //    3) score: 8;
  //    4) order: ""; skipped;
  //    5) kind: MANGA_KIND_MANGA;
  //    6) status: MANGA_STATUS_RELEASED;
  //    7) season: ""; skipped;
  //    8) mylist: MY_LIST_COMPLETED;
  //    9) censored: false;
  //    10) genre: nil; skipped;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.MangaSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(m.Errors)
  // Standard output of our search, nothing new.
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

  // The first parameter is the values of the character; values: "id", "name", "russian", "url", "description".
  // The second parameter is the name of the character; name: "onizuka".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 2;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.CharacterSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(ch.Errors)
  // Standard output of our search, nothing new.
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

  // The first parameter is the values of the people; values: "id", "name", "russian", "url",
  // "website", "birthOn{year month day date}".
  // The second parameter is the name of the people; name: "satsuki".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) isSeyu: true;
  //    4) isMangaka: false;
  //    5) isProducer: false;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.PeopleSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(p.Errors)
  // Standard output of our search, nothing new.
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

  // The first parameter is the values of the userRates; values: "id",
  // "text", "score", "createdAt", "anime{id name}",
  // The second parameter is the user Id; userId: 181833.
  // In the third parameter, we introduce an auxiliary function that
  // will separate the two additional fields: "order: {field: id, order: desc}".
  // Now let's move on to the interface:
  //    1) page: 1;
  //    2) limit: 2;
  //    3) status: completed;
  //    4) targetType: Anime;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.UserRatesSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(ur.Errors)
  // Standard output of our search, nothing new.
  for _, v := range ur.Data.UserRates {
    fmt.Println(v.Id, v.Text, v.Score, v.CreatedAt, v.Anime.Id, v.Anime.Name)
  }
}
```
