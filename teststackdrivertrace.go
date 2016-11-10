package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"cloud.google.com/go/trace"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/option"
)

var (
	projectID         = flag.String("projectID", "", "projcect id")
	serviceAccountKey = flag.String("serviceAccountKey", "", "path to serviceacount json")
)

func main() {
	flag.Parse()
	if err := run(*serviceAccountKey); err != nil {
		log.Fatalln(err)
	}
}

func run(keyPath string) error {
	mux := http.NewServeMux()
	return http.ListenAndServe(":9000", Trace(*projectID, *serviceAccountKey, mux))
}

func Trace(projectID string, keyPath string, h http.Handler) http.HandlerFunc {
	client, err := trace.NewClient(context.Background(), projectID, option.WithServiceAccountFile(keyPath))
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s := client.SpanFromRequest(r)
		defer func() { err := s.FinishWait(); spew.Dump(err) }()
		h.ServeHTTP(w, r)
	}
}
