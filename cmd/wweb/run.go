package main

import (
	"context"
	"go-learn/internal/app/wweb/http"
	"go-learn/internal/app/wweb/persistence/orm"
	"go-learn/internal/app/wweb/persistence/redis"
	"go-learn/pkg/logger"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "start web server",
	Long:  "start web server",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	logger.InitLogger(viper.GetString("log.filename"), viper.GetString("log.level"))
	defer logger.Close()

	orm.Init()
	defer orm.Close()
	redis.Init()
	defer redis.Close()

	server := http.Init()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Http server listen err : %s\n", err)
		}
	}()

	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-sc
	atomic.StoreInt32(&state, 0)
	log.Println("收到退出信号[%s]", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("服务退出")
	os.Exit(int(atomic.LoadInt32(&state)))
}
