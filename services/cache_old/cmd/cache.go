package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pad/services/cache_old/internal/config"
	"pad/services/cache_old/internal/server"
	"pad/services/cache_old/services/cache"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	helpFlag = flags.Bool("help", false, "Test")
)

func main() {
	flags.Parse(os.Args[1:])

	if *helpFlag {
		usage()
		os.Exit(0)
	}

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		osError("failed to load env: %v\n", err)
	}

	srv := server.NewCacheServer()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.GRPCPort))
	if err != nil {
		osError("failed to listen to tcp server: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	cache.RegisterCacheServer(grpcServer, srv)
	reflection.Register(grpcServer)

	log.Println("Starting the server")
	if err := grpcServer.Serve(grpcListener); err != nil {
		osError("failed to run grpc server: %v", err)
	}
}

func usage() {
	fmt.Printf(`
			Usage: %s [flags]
			TBD: General description will go here	
			Flags available:
		`,
		os.Args[0],
	)

	flags.PrintDefaults()
}

func osError(format string, opts ...any) {
	fmt.Fprintf(os.Stderr, format, opts...)
	os.Exit(1)
}
