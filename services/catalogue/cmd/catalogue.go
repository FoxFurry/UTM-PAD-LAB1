package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"pad/services/cache/client"
	"pad/services/catalogue/internal/config"
	"pad/services/catalogue/internal/server"
	"pad/services/catalogue/internal/store"
	"pad/services/catalogue/services/catalogue"
	"pad/services/common/database"
	"pad/services/locator/go-client/locator"
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

	ctx := context.Background()

	db, err := database.InitDB(mysql.Config{
		DBName:               cfg.DBName,
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPass,
		Addr:                 cfg.DBAddr,
		ParseTime:            true,
		AllowNativePasswords: true,
	})
	if err != nil {
		osError("failed to initialize database connection: %v\n", err)
	}

	cacheClient, err := client.NewCacheClient(cfg.CacheAddress)
	if err != nil {
		osError("failed to connect to cache service: %v\n", err)
	}

	srv := server.NewCatalogueServer(store.NewCatalogueStore(db), cacheClient)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.GRPCPort))
	if err != nil {
		osError("failed to listen to tcp server: %v", err)
	}

	var opts = []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Second),               // REQ: Task timeout
		grpc.MaxConcurrentStreams(cfg.MaxConcurrentTasks), // REQ: Concurrent tasks limit
	}

	grpcServer := grpc.NewServer(opts...)

	catalogue.RegisterCatalogueServer(grpcServer, srv)
	reflection.Register(grpcServer)

	log.Println("Starting the server")
	go func() {
		if err := grpcServer.Serve(grpcListener); err != nil {
			osError("failed to run grpc server: %v", err)
		}
	}()

	time.Sleep(time.Millisecond)
	conn, err := grpc.Dial(cfg.LocatorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		osError("failed to create locator dialer: %v", err)
	}

	locatorClient := locator.NewLocatorClient(conn)

	if _, err := locatorClient.RegisterService(ctx, &locator.RegisterServiceRequest{
		Type:    1,
		Address: fmt.Sprintf("127.0.0.1:%d", cfg.GRPCPort),
	}); err != nil {
		osError("failed to register locator dialer: %v", err)
	}

	log.Println("Successfully registered service")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
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
