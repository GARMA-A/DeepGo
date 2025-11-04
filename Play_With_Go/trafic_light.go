package main

import (
	"fmt"
	"time"
)

func loop() {
	for {
		t := 5 * time.Second
		for i := 0; i <= 5; i += 2 {
			fmt.Print("open", i, "\n")
			time.Sleep(t)
			fmt.Print("close", i, "\n")

			if t == 5*time.Second {
				t = 2 * time.Second
			} else {
				t = 5 * time.Second
			}
		}
	}
}
