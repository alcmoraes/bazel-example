package main

import (
	"fmt"
	"go-webview-wasm/packages/shared/handlers/router"
)

func main() {
	fmt.Println("Run main second router")
	_ = router.GetEngine("Main App").Run(":8080")
}
