package main

import (
	"context"
	"fmt"
	http1 "net/http"

	pb "github.com/go-kratos/kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func globalFilter(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("global filter")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func globalFilter2(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("global filter 2")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func routeFilter(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("route filter")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func routeFilter2(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("route filter 2")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func pathFilter(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("path filter")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func pathFilter2(next http1.Handler) http1.Handler {
	return http1.HandlerFunc(func(w http1.ResponseWriter, r *http1.Request) {
		// Do stuff here
		fmt.Println("path filter 2")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func serviceMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		fmt.Println("service middleware")
		return handler(ctx, req)
	}
}

func serviceMiddleware2(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		fmt.Println("service middleware 2")
		return handler(ctx, req)
	}
}

// this example shows how to add middlewares,
// execution order is globalFilter(http) --> routeFilter(http) --> pathFilter(http) --> serviceFilter(service)
func main() {
	s := &server{}
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			// add service filter
			serviceMiddleware,
			serviceMiddleware2,
		),
		// add global filter
		http.Filter(globalFilter, globalFilter2),
	)
	// register http hanlder to http server
	pb.RegisterGreeterHTTPServer(httpSrv, s)

	// add route filter
	r := httpSrv.Route("/", routeFilter, routeFilter2)
	// add path filter to custom route
	r.GET("/test/{name}", testHandler, pathFilter, pathFilter2)

	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
