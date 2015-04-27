import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)

log.Fatal(http.ListenAndServe(":8080", nil))
