package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/wellWINeo/GoIpBot"
	"github.com/wellWINeo/GoIpBot/pkg/handler"
	"github.com/wellWINeo/GoIpBot/pkg/repository"
	"github.com/wellWINeo/GoIpBot/pkg/service"
)


func main() {
	GoIpBot.Log("main.go").Info("starting...")

	if err := godotenv.Load(); err != nil {
		GoIpBot.Log("main.go").Error(err)
	} else {
		GoIpBot.Log("main.go").Info(".env file read")
	}

	// read config file
	err := InitConfig()
	if err != nil {
		GoIpBot.Log("main.go").Fatal(err)
	} else {
		GoIpBot.Log("main.go").Info("config read properly")
	}

	// run tg bot
	bot, err := GoIpBot.NewBot(GoIpBot.BotConfig{
		Token: os.Getenv("BOT_TOKEN"),
		Debug: true,
		Timeout: 60,
	})
	if err != nil {
		GoIpBot.Log("main.go").Fatal(err)
	} else {
		GoIpBot.Log("main.go").Info("bot instance created")
	}


	db, err := repository.NewGORMDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.name"),
		Port:     viper.GetInt("db.port"),
	})

	if err != nil {
		GoIpBot.Log("main.go").Fatal(err)
	} else {
		GoIpBot.Log("main.go").Info("connected to database")
	}

	db.AutoMigrate(&GoIpBot.User{}, &GoIpBot.HistoryRecord{})

	repo := repository.NewRepository(os.Getenv("IPSTACK_TOKEN"), db)
	services := service.NewService(repo)
	tgHandler := handler.NewTeleHandler(services)
	webHandler := handler.NewWebHandler(services)

	errChan, inChan, outChan := bot.Poll()
	tgHandler.InitRoutes(inChan, outChan)

	GoIpBot.Log("main.go").Info("bot polling")

	go func () {
		errChan <- webHandler.InitRoutes(viper.GetInt("server.port"))
	}()

	GoIpBot.Log("main.go").Info("web server started")

	err = <- errChan
	GoIpBot.Log("main.go").Fatal(err)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
