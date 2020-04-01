package wordsearch

import (
	context "context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
)

// SearchClient type
type SearchClient struct {
	wg sync.WaitGroup
}

// NewSearchClient func
func NewSearchClient() *SearchClient {
	return &SearchClient{}
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

	RegisterWordSearchServer(grpcServer, sc)

	grpcServer.Serve(lis)
	return nil
}

// StartRest func
func (sc *SearchClient) StartRest() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterWordSearchHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8090", mux)
}

// ListSearchTerms func
func (sc *SearchClient) ListSearchTerms(ctx context.Context, req *ListWordsRequest) (*ListWordsResponse, error) {
	return &ListWordsResponse{
		Message: fmt.Sprint("Hiii!!"),
	}, nil
}
