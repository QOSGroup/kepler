package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/server/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the kepler server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		wg := sync.WaitGroup{}
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)

		r := gin.Default()
		router.InitRouter(r)

		go func() {
			if err := r.Run(":8080"); err != nil {
				panic(err)
			}
		}()

		wg.Wait()
		TrapSignal(func() {})

		return nil
	},
}

func TrapSignal(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, exiting...\n", sig)
			if cb != nil {
				cb()
			}
			os.Exit(1)
		}
	}()
	select {}
}
