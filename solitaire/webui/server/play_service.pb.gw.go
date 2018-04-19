// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: webui/server/play_service.proto

/*
Package server is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package server

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_PlayService_NewGame_0(ctx context.Context, marshaler runtime.Marshaler, client PlayServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq NewGameRequest
	var metadata runtime.ServerMetadata

	msg, err := client.NewGame(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterPlayServiceHandlerFromEndpoint is same as RegisterPlayServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPlayServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPlayServiceHandler(ctx, mux, conn)
}

// RegisterPlayServiceHandler registers the http handlers for service PlayService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPlayServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPlayServiceHandlerClient(ctx, mux, NewPlayServiceClient(conn))
}

// RegisterPlayServiceHandler registers the http handlers for service PlayService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "PlayServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PlayServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PlayServiceClient" to call the correct interceptors.
func RegisterPlayServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PlayServiceClient) error {

	mux.Handle("GET", pattern_PlayService_NewGame_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PlayService_NewGame_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PlayService_NewGame_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PlayService_NewGame_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v2", "new"}, ""))
)

var (
	forward_PlayService_NewGame_0 = runtime.ForwardResponseMessage
)