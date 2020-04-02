package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	wordsearch "github.com/RuNpiXelruN/word-search/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup

// Start func is the entry point
// to the word-search app. It accepts
// the WordSearchService interface.
func Start(wsc WordSearchService) {
	fmt.Println("GRPC Wordsearch Server is running...")

	wg.Add(1)
	go func() {
		log.Fatal(wsc.StartGRPC())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Fatal(wsc.StartRest())
		wg.Done()
	}()
	wg.Wait()
}

// StartGRPC func starts our
// GRPC server on port ':8080'
func (wsc *WordSearchClient) StartGRPC() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	wordsearch.RegisterWordSearchServer(grpcServer, wsc)

	grpcServer.Serve(lis)
	return nil
}

// StartRest func starts our Rest server on
// port ':8090', listens for incoming requests
// and forwards them onto our GRPC server
func (wsc *WordSearchClient) StartRest() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// the below is required to json marshal default
	// values, ie 0 for the starting search_count.
	// protocol buffers has 'emitempty' harcoded into
	// them and is a known behaviour they're looking
	// at fixing. If runtime.NewServeMux() only is used,
	// default values will not print out.
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := wordsearch.RegisterWordSearchHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8090", mux)
}
