package zaperr

import (
	"net/http"

	"go.uber.org/zap"
)

func main() {
	zap, _ := zap.NewDevelopment()

	handleError := NewZaperr(zap)

	handleError.LogError(http.ListenAndServe(":8080", nil), "Failed run server")

}
