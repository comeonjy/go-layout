package main

import (
	"context"
	"log"

	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-layout/cmd"
	"github.com/comeonjy/go-layout/pkg/consts"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(c *cobra.Command, args []string) {
		xenv.Init(consts.EnvMap)
		logger := xlog.New(xlog.WithTrace(consts.TraceName))
		ctx, cancel := context.WithCancel(context.Background())
		app := cmd.InitApp(ctx, logger)
		if err := app.Run(cancel); err != nil {
			log.Println("Server Exit:", err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
