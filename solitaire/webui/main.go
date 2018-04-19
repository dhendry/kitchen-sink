package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dhendry/kitchen-sink/solitaire/webui/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//<editor-fold desc="GRPC server stuff">
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	server.RegisterPlayServiceServer(grpcServer, &server.PlayServiceImpl{})

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	go func() {
		log.Println("Starting gRPC server")
		defer func() {
			log.Println("gRPC goroutine exiting")
		}()
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	//</editor-fold>

	//<editor-fold desc="GRPC gateway stuff">
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gmux := runtime.NewServeMux()
	err = server.RegisterPlayServiceHandlerFromEndpoint(ctx, gmux, ":9090", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}
	// Note that this is hardcoding the /api/v2/ dispatch to the grpc gateway
	http.HandleFunc("/api/v2/", func(writer http.ResponseWriter, request *http.Request) {
		gmux.ServeHTTP(writer, request)
	})
	//</editor-fold>

	// API v1
	//server.RegisterApiHandlers()
	//server.RegisterUiHandlers()

	//<editor-fold desc="Proxy to frontend server">
	frontendUrl, e := url.Parse("http://localhost:8080/")
	if e != nil {
		panic(e)
	}
	proxy := httputil.NewSingleHostReverseProxy(frontendUrl)
	http.Handle("/", proxy)
	//</editor-fold>

	// TODO: Configurable port
	log.Println("Starting server")
	defer func() {
		log.Println("Main goroutine exiting")
	}()
	log.Fatal(http.ListenAndServe(":8081", nil))


	log.Println("Exiting")
}
