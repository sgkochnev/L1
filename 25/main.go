package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func Sleep1(d time.Duration) {
	<-time.After(d)
}

func Sleep2(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	<-ctx.Done()
}

func main() {
	s := time.Now()
	Sleep1(1 * time.Second)
	fmt.Println(time.Since(s))

	s = time.Now()
	Sleep2(1 * time.Second)
	fmt.Println(time.Since(s))
}
