package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func sleepAndTalk(ctx context.Context, duration time.Duration, message string) {
	select {
	case <-time.After(duration):
		fmt.Println(message)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 5*time.Second, "Hello context!")
}
