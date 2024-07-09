package health

import (
	"fmt"
	"net/http"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The system is healthy")
}
