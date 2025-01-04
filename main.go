package main

import (
	"context"
	"github.com/rodziievskyi-maksym/go-essentials/v2_2025/util"
	"os/signal"
	"syscall"
)

func main() {
	util.PrintMainSeparationMessage()
	/* alt + command + / -> creates comment block section*/

	// Part one
	//basics.RunFirstPart()

	// Part two
	//basics.RunSecondPart()

	// Part three
	//basics.WorkWithFiles()

	// Part Four
	//basics.MapsAndGenericsAndStructs()

	// Part Five Struct Encapsulation Getters/Setters And Package Practice
	//mypackage.SetterGetterEncapsulation()

	// Interfaces
	//basics.StringCaller()
	//basics.AreaCaller()

	// Concurrency -> GoRoutine
	//concurrency.CountToCaller()

	// Concurrency -> GoRoutine -> Channels
	//concurrency.IntReceiverCaller()

	// Concurrency -> GoRoutine -> Mutex
	//concurrency.MutexCaller()

	//Closures
	//basics.Closures()

	//Recursion
	//basics.Recursion()

	//shutdownCtx, cancel := context.WithCancel(context.Background())
	//osSighCh := make(chan os.Signal)
	//signal.Notify(osSighCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//go func() {
	//	<-osSighCh
	//	cancel()
	//}()

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	_ = cancel

	<-shutdownCtx.Done()
	//graceful_shutdown.GracefulShutdown()
}
