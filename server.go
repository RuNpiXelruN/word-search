package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	wordsearch "github.com/RuNpiXelruN/word-search/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Testt func
func Testt() {
	fmt.Println("hihi")
}

// Start func
func (sc *SearchClient) Start() {
	fmt.Println("Hi from grpc package")

	sc.wg.Add(1)
	go func() {
		log.Fatal(sc.StartGRPC())
		sc.wg.Done()
	}()

	sc.wg.Add(1)
	go func() {
		log.Fatal(sc.StartRest())
		sc.wg.Done()
	}()

	sc.wg.Wait()
}

// StartGRPC func
func (sc *SearchClient) StartGRPC() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	wordsearch.RegisterWordSearchServer(grpcServer, sc)

	grpcServer.Serve(lis)
	return nil
}

// StartRest func
func (sc *SearchClient) StartRest() error {
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
