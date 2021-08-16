package main

import (
	"bazel-example/packages/shared/handlers/router"
	"fmt"
)

func main() {
	fmt.Println("Run main second router")
	_ = router.GetEngine("Second App").Run(":8080")
}
