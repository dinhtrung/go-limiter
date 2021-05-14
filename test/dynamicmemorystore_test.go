package test

import (
	"context"
	"log"
	"time"

	memorystore "github.com/sethvargo/go-limiter/dynamicmemorystore"
)

func ExampleNew() {
	ctx := context.Background()

	store, err := memorystore.New(&memorystore.Config{
		Tokens:   []uint64{1, 2, 3, 4},
		Interval: time.Minute,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close(ctx)

	limit, remaining, reset, ok, err := store.Take(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, _ = limit, remaining, reset, ok
}
