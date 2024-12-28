# zaperr

A simple Go package for logging errors using Uber's zap logger.

`go get -u github.com/MiracleCanCode/zaperr`

## Example usage

### Simple usage

```go
package main

import (
	"net/http"
	"go.uber.org/zap"
	"github.com/MiracleCanCode/zaperr"
)

func main() {
	// Create a new Zap logger instance
	zap, _ := zap.NewDevelopment()

	// Initialize the zaperr package with the zap logger
	handleError := zaperr.NewZaperr(zap)

	// Log an error if the HTTP server fails
	handleError.LogError(http.ListenAndServe(":8080", nil), "Failed to run server")
}
```

### With additional custom logic

```go
package main

import (
	"net/http"
	"go.uber.org/zap"
	"fmt"
	"github.com/MiracleCanCode/zaperr"
)

func main() {
	// Create a new Zap logger instance
	zap, _ := zap.NewDevelopment()

	// Initialize the zaperr package with the zap logger
	handleError := zaperr.NewZaperr(zap)

	// Log an error and execute custom logic
	handleError.LogError(http.ListenAndServe(":8080", nil), "Failed to run server", func() {
		// Custom logic (e.g., send a notification or perform cleanup)
		fmt.Println("Executing additional logic due to error!")
	})
}
```
