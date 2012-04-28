package webserver

import (
  "net/http"
)

func ExampleRunCLI() {
  dispatcher := func (w http.ResponseWriter, r *http.Request) {
    // do something with the request, this could dispatch
    // the request to another component, for example.
  }

  srv := New(http.HandlerFunc(dispatcher))
  srv.RunCLI()
}
