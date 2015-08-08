// Simpler counter server.
import (
	"fmt"
	"net/http"
)

type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}
