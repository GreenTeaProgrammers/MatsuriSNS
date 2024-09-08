package main

import (
    "log"

    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
)

func main() {
    err := entc.Generate("./schema", &gen.Config{
        Target: "./ent",
    })
    if err != nil {
        log.Fatalf("failed to generate ent code: %v", err)
    }
}

