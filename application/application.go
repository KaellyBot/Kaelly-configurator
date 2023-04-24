package application

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	guildRepo "github.com/kaellybot/kaelly-configurator/repositories/guilds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/configurators"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/utils/databases"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func New() (*Application, error) {
	// misc
	db, err := databases.New()
	if err != nil {
		return nil, err
	}

	broker, err := amqp.New(constants.RabbitMQClientId, viper.GetString(constants.RabbitMqAddress),
		[]amqp.Binding{configurators.GetBinding()})
	if err != nil {
		return nil, err
	}

	// repositories
	guildRepo := guildRepo.New(db)
	chanServerRepo := servers.New(db)
	almanaxRepo := almanax.New(db)
	feedsRepo := feeds.New(db)
	twitterRepo := twitter.New(db)

	// services
	guildService, err := guilds.New(guildRepo)
	if err != nil {
		return nil, err
	}

	channelService, err := channels.New(chanServerRepo, almanaxRepo, feedsRepo, twitterRepo)
	if err != nil {
		return nil, err
	}

	configService, err := configurators.New(broker, guildService, channelService)
	if err != nil {
		return nil, err
	}

	return &Application{
		guildService:        guildService,
		channelService:      channelService,
		configuratorService: configService,
		broker:              broker,
	}, nil
}

func (app *Application) Run() error {
	return app.configuratorService.Consume()
}

func (app *Application) Shutdown() {
	app.broker.Shutdown()
	log.Info().Msgf("Application is no longer running")
}
