package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pad/services/cache/internal/api/server"
	"pad/services/cache/services/cache"
)

var grpcPort string

// serveGrpcCmd represents the serveGrpc command
var serveGrpcCmd = &cobra.Command{
	Use:   "serve-grpc",
	Short: "A brief description of your command",
	Args:  cobra.NoArgs,
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if grpcPort == "" {
			fmt.Println("Cannot use empty port")
			return
		}

		srv := server.NewCacheServer()

		grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", grpcPort))
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
	},
}

func init() {
	rootCmd.AddCommand(serveGrpcCmd)

	serveGrpcCmd.PersistentFlags().StringVarP(&grpcPort, "grpc-port", "g", "23000", "Defines GRPC port for memstore server. Default value is 8080")
}

func osError(format string, opts ...any) {
	fmt.Fprintf(os.Stderr, format, opts...)
	os.Exit(1)
}
