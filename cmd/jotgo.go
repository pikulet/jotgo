package main

import (
    "github.com/pikulet/jotgo/game"
tm  "github.com/buger/goterm"
)

func main() {
    tm.Clear()

    tm.Println(tm.Color(tm.Bold("JOTGO"), tm.RED))

    tm.Flush()
}
