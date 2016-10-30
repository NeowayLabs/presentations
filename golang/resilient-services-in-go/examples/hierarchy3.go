package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
	child, childCancel := context.WithTimeout(parent, 10*time.Second)
	defer childCancel()
	waitParent, waitChild := parent.Done(), child.Done()
	go func() { parentCancel() }()
	for {
		select {
		case <-waitParent:
			fmt.Println("Parent done")
			waitParent = nil
		case <-waitChild:
			fmt.Println("Child done")
			waitChild = nil
		}
	}
}
