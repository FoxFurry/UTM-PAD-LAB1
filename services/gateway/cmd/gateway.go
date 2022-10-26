package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pad/services/gateway/internal/config"
	"pad/services/gateway/internal/http/server"
	"pad/services/locator/go-client/locator"
	userClient "pad/services/user/client"
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
	conn, err := grpc.Dial(cfg.LocatorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		osError("failed to create locator dialer: %v", err)
	}

	locatorClient := locator.NewLocatorClient(conn)

	user, err := userClient.NewUserClient(cfg.UserAddress)
	if err != nil {
		osError("failed to connect to user server: %v", err)
	}

	srv := server.NewGatewayServer(user, locatorClient)

	log.Println("Starting the server")
	if err := srv.Start(fmt.Sprintf("127.0.0.1:%d", cfg.Port)); err != nil {
		osError("failed to run gateway server: %v", err)
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
