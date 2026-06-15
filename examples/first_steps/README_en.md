## Links
* [OAuth2](https://shikimori.one/oauth)

## Basic structure

```golang
package main

import (
  g "github.com/ghostemeow/goshikimori"
)

/*

DUMMIES GUIDE.

IF YOU KNOW WHAT YOU'RE WORKING WITH, DON'T WASTE YOUR TIME AND GO STRAIGHT TO THE DOCUMENTATION.

--------------------------

To get the APPLICATION_NAME and PRIVATE_KEY fields
go to the link above and create an application,
or join "Test Api" and use it, it's up to you.

--------------------------

For most queries, like searching
anime/manga/users and the like,
the APPLICATION_NAME field will suffice, and the
PRIVATE_KEY can be left blank.

func config() *g.Configuration {
  return g.SetConfiguration(
    "APPLICATION_NAME",
    "",
  )
}

--------------------------

If you need to change some information, such as
add a friend, add to ignore or write
message, then the PRIVATE_KEY field is mandatory..

IMPORTANT: your application, when registering, must have
have access rights (what it can do on your behalf)
for certain operations.

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
