Wrangler CLI
========

Wrangler CLI is a trivial wrapper around `github.com/spf13/cobra` that allow one to define
a simple struct with tags as a way to create `cobra.Command` objects.  For example

```go
package main

import (
    "fmt"

    "github.com/rancher/wrangler-cli"
    "github.com/spf13/cobra"
)

type App struct {
    OptionOne string `name:"custom-name" usage:"Takes an option""`
}

func (a *App) Run(cmd *cobra.Command, args []string) error {
    if len(args) == 0 {
        return cmd.Help()
    }
    fmt.Printf("Your option one %s and args: %v\n", a.OptionOne, args)
    return nil
}

func main() {
    root := cli.Command(&App{}, cobra.Command{
        Short: "Some short description",
        Long:"Some long description",
    })
    cli.Main(root)
}

```

## License
Copyright (c) 2020 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
