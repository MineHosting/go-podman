# How Start a new Project with Go-Podman

To start a new project in go-podman is create the files or directories that you need.

And follow the normal go - standard to develop a new app. 🤷

Use go get to receive the most recent version of this lib and create a new client with

```go
package main

import (
  "log"
  "github.com/MineHosting/go-podman/pkg/engine"
)

func main() {
  client, err := engine.Start(true)

  if err != nil {
    log.Fatalf(err)
  }

  //...
}
```

The parameter true in Engine.Start is a answer that we make.

> engine: Do you want run in rootless mode?
> user: Yes, I want. (true)

In case the parameter be false the engine create a client in rootfull mode
