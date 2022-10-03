package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"pad/services/catalogue/internal/config"
	"pad/services/catalogue/internal/server"
	"pad/services/catalogue/internal/store"
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
		fmt.Fprintf(os.Stderr, "cfgFlag is mandatory for execution\n")
		os.Exit(1)
	}

	viper.SetConfigFile(*cfgFlag)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read cfgFlag: %v\n", err)
		os.Exit(1)
	}

	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	db, err := database.InitDB(ctx, mysql.Config{
		DBName: cfg.Name,
		User:   cfg.User,
		Passwd: cfg.Pass,
		Addr:   cfg.Addr,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize database connection: %v\n", err)
		os.Exit(1)
	}

	srv := server.NewCatalogueServer(store.NewCatalogueStore(db))

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
