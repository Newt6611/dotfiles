package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	p_redis "github.com/redis/go-redis/v9"
	"github.com/strike-finance/strike-v2-backend/pkg/redis"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/config"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/service"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/service/price"
)

func main() {

	cfg := config.Init()

	// Init redis
	redisClient := initRedis(cfg)

	services := initAllServices(cfg, redisClient)

	activeAllServices(services)

	log.Info("🚀 Sequencer Started")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Info("🛑 Shutting down Sequencer...")

	shutdownAllServices(services)

	time.Sleep(time.Second * 3)

	log.Info("✅ Shutdown complete")
}

func initRedis(cfg *config.Config) *redis.Redis {
	redisOpt, err := p_redis.ParseURL(cfg.RedisURL)
	if err != nil {
		panic(err)
	}
	r := p_redis.NewClient(redisOpt)
	return redis.NewRedis(r)
}

func initAllServices(
	cfg *config.Config,
	redisClient *redis.Redis,
) []service.Service {

	services := []service.Service{
		price.NewService(cfg, redisClient),
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
