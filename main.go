package main

import (
	"net/http"
	"github.com/docker/go-plugins-helpers/sdk"
	"fmt"
)

func main() {
	fmt.Println(">>>> Start plugin")
	h := sdk.NewHandler(`{"Implements": ["authz"]}`)
	handlers(&h)
	if err := h.ServeUnix("authz", 0); err != nil {
		panic(err)
	}
}

func handlers(h *sdk.Handler) {
	h.HandleFunc("/AuthZPlugin.AuthZReq", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AuthZReq: %v\n", r)

	})
	h.HandleFunc("/AuthZPlugin.AuthZRes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AuthZRes: %v\n", r)
		return
	})
}
