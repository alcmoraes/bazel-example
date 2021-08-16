package main

import (
	"bazel-example/packages/shared/handlers/router"
	"fmt"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine("Main App").Run(":8080")
}

func Check() {
	fmt.Println("Check")
}
