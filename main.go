
package main

import (
  "fmt"
  "github.com/gochaos-app/go-chaos-plugin/addon"
)


func ChaosFunc(c addon.ChaosConfig) {
	fmt.Println("Hello from Chaos")
	fmt.Println(c.Region)
}

