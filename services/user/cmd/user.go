package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pad/services/user/internal/config"
	"pad/services/user/internal/server"
	"pad/services/user/internal/store"
	"pad/services/user/services/user"

	"pad/services/common/database"
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

	//ctx := context.Background()

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

	srv := server.NewUserServer(store.NewUserStore(db))

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.GRPCPort))
	if err != nil {
		osError("failed to listen to tcp server: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	user.RegisterUserServer(grpcServer, srv)
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
