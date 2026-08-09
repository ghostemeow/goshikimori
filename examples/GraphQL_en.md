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
  // The third parameter is the AnimeOptions struct; the zero value of a field means the option is skipped:
  //    Page: 1;
  //    Limit: 5;
  //    Score: 8;
  //    Kind: ANIME_KIND_TV;
  //    Status: ANIME_STATUS_RELEASED;
  //    Rating: ANIME_RATING_PG_13;
  //    Censored: false;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.AnimeSchema();
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
  // The second parameter is the name of the manga; name: "liar game".
  // The third parameter is the MangaOptions struct; the zero value of a field means the option is skipped:
  //    Page: 1;
  //    Limit: 1;
  //    Score: 8;
  //    Kind: MANGA_KIND_MANGA;
  //    Status: MANGA_STATUS_RELEASED;
  //    Mylist: MY_LIST_COMPLETED;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.MangaSchema();
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
  // The third parameter is the CharacterOptions struct; the zero value of a field means the option is skipped:
  //    Page: 1;
  //    Limit: 2;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.CharacterSchema();
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
  // The third parameter is the PeopleOptions struct:
  //    Page: 1;
  //    Limit: 1;
  //    IsSeyu: true;
  //    IsMangaka: false;
  //    IsProducer: false;
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.PeopleSchema();
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
  // The third parameter is the UserRatesOptions struct; the zero value of a field means the option is skipped:
  //    Page: 1;
  //    Limit: 10;
  //    Status: MY_LIST_COMPLETED;
  //    TargetType: TARGET_TYPE_ANIME;
  //    Order: UserRatesOrder() - an auxiliary function that adds "order: {field: id, order: desc}".
  //
  // The available values can be found in the function description: shiki.ValuesSchema();
  // The available interface parameters can be found in the function description: shiki.UserRatesSchema();
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

  // Here you can track errors received during server response.
  fmt.Println(ur.Errors)
  // Standard output of our search, nothing new.
  for _, v := range ur.Data.UserRates {
    fmt.Println(v.Id, v.Text, v.Score, v.CreatedAt, v.Anime.Id, v.Anime.Name)
  }
}
```
