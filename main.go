
package main

import (
  "fmt"
  "./addon"
)


func ChaosFunc(c addon.ChaosConfig) {
	fmt.Println("Hello from Chaos")
	fmt.Println(c.Region)
}

