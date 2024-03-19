package main

import (
	"context"
	"fmt"
	"github.com/crewlinker/fleetmon/maritime"
	"os"
)

func main() {
	apiKey := os.Getenv("FLEETMON_KEY")
	if len(apiKey) == 0 {
		panic("FLEETMON_KEY unset")
	}

	cli := maritime.NewClient(apiKey)
	search, err := cli.VesselSearch(
		context.Background(),
		maritime.VesselSearchParams{
			ShipName: "FLORIDA",
		},
	)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	for _, v := range search {
		fmt.Printf("%+v\n", v)
	}
}
