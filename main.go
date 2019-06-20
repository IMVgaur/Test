package main

import (
	"fmt"
)

func main() {
	http.ListenAndServe(":3000", controllers.Handlers())
}
