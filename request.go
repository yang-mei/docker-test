package main

import (
        "net/http"
)

func main() {

        http.HandleFunc("/v1/test", func(writer http.ResponseWriter,
                request *http.Request) {
                writer.Write([]byte("OK"))
                return
        })
        http.HandleFunc("/v1/offline-acceptance/trans-type-and-status", func(writer http.ResponseWriter, request *http.Request) {
                writer.Write([]byte("YES"))
                return
        })
        http.ListenAndServe(":6116", nil)
}
iii
