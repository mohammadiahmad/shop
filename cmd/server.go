/*
Copyright © 2021 NAME HERE <cen.ahmadm@gmail.com>

*/
package cmd

import (
	"github.com/mohammadiahmad/shop/internal/app"
	"github.com/mohammadiahmad/shop/internal/cart_storage"
	"github.com/mohammadiahmad/shop/internal/config"
	"github.com/mohammadiahmad/shop/internal/readisearch"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run docs server",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func server() {
	cfg := config.Load()
	l, _ := zap.NewProduction()
	s, err := readisearch.NewReadisearch(cfg.Storage, l)
	if err != nil {
		panic(err)
	}

	cs:=cart_storage.NewCartStorage(cfg.Redis)
	a := app.NewApp(&cfg.Server, s,cs)
	a.Run()
}
