package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent, parentCancel := context.WithTimeout(context.Background(), 10*time.Second)
	child, childCancel := context.WithTimeout(parent, 5*time.Second)
	defer parentCancel()
	defer childCancel()
	waitParent, waitChild := parent.Done(), child.Done()
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
