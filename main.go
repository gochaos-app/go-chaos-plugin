
package main

import (
  "fmt"
)


func ChaosFunc(region string, service string, project string, namespace string, tag string, chaos string, number int) {
	fmt.Println("Hello from Chaos")
	fmt.Println(region)
}

