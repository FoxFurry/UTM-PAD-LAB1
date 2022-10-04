package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pad/services/catalogue/internal/config"
	"pad/services/catalogue/internal/server"
	"pad/services/catalogue/internal/store"
	catalogue "pad/services/catalogue/proto"
	"pad/services/common/database"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	cfgFlag  = flags.String("config", "", "File containing configuration for. Find more about it on official github page")
	helpFlag = flags.Bool("help", false, "Test")
)

func main() {
	flags.Parse(os.Args[1:])

	if *helpFlag {
		usage()
		os.Exit(0)
	}

	if *cfgFlag == "" {
		osError("cfgFlag is mandatory for execution\n")
	}

	viper.SetConfigFile(*cfgFlag)
	if err := viper.ReadInConfig(); err != nil {
		osError("failed to read cfgFlag: %v\n", err)
	}

	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		osError("failed to read config: %v\n", err)
	}

	ctx := context.Background()

	db, err := database.InitDB(ctx, mysql.Config{
		DBName: cfg.DBName,
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Addr:   cfg.DBAddr,
	})
	if err != nil {
		osError("failed to initialize database connection: %v\n", err)
	}

	srv := server.NewCatalogueServer(store.NewCatalogueStore(db))

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
	if err != nil {
		osError("failed to listen to tcp server: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	catalogue.RegisterCatalogueServer(grpcServer, srv)
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
