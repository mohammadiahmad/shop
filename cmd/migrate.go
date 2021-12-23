/*
Copyright © 2021 NAME HERE <cen.ahmadm@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/mohammadiahmad/shop/internal/config"
	"github.com/mohammadiahmad/shop/internal/storage"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate db schema and insert test data",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate() {
	cfg := config.Load()
	l, _ := zap.NewProduction()
	s, err := storage.NewDB(cfg.Storage, l)
	if err != nil {
		panic(err)
	}

	err = s.Migrate()
	if err != nil {
		fmt.Printf("Error in db migrate %+v\n", err)
		panic(err)
	}

	s.InitDB()

	fmt.Println("Db migrate successfull")

}
