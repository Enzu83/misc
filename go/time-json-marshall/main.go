package main

import (
    "encoding/json"
    "fmt"
    "time"
)

func main() {
    tm := time.Unix(1756201396, 0)
    tm_marshall, _ := json.Marshal(tm)
	var tm_unmarshall string
	json.Unmarshal(tm_marshall, &tm_unmarshall)
	fmt.Println(tm_unmarshall)

	const layout = "2006-01-02T15:04:05.999999-07:00"
	tm_parsed, _ := time.Parse(layout, tm_unmarshall)

	fmt.Println(tm_parsed)
}