package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/config"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/service"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/service/price"
)

func main() {

	cfg := config.Init()

	services := initAllServices(cfg)

	activeAllServices(services)

	log.Info("🚀 Sequencer Started")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Info("🛑 Shutting down Sequencer...")

	shutdownAllServices(services)

	log.Info("✅ Shutdown complete")
}

func initAllServices(cfg *config.Config) []service.Service {
	services := []service.Service{
		price.NewService(cfg),
	}
	return services
}

func activeAllServices(services []service.Service) {
	for _, service := range services {
		go service.Active()
	}
}

func shutdownAllServices(services []service.Service) {
	for _, service := range services {
		go service.Shutdown()
	}
}
