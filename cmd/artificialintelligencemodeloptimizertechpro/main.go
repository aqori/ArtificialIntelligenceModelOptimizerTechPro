// cmd/artificialintelligencemodeloptimizertechpro/main.go
package main

import (
"flag"
"log"
"os"

"artificialintelligencemodeloptimizertechpro/internal/artificialintelligencemodeloptimizertechpro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialintelligencemodeloptimizertechpro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
