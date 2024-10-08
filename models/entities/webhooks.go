package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookFeed struct {
	WebhookID    string
	WebhookToken string
	GuildID      string        `gorm:"primaryKey"`
	ChannelID    string        `gorm:"primaryKey"`
	FeedTypeID   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	FeedSource   FeedSource    `gorm:"foreignKey:FeedTypeID,Locale"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitch struct {
	WebhookID    string
	WebhookToken string
	GuildID      string   `gorm:"primaryKey"`
	ChannelID    string   `gorm:"primaryKey"`
	StreamerID   string   `gorm:"primaryKey"`
	Streamer     Streamer `gorm:"foreignKey:StreamerID"`
	RetryNumber  int64    `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitter struct {
	WebhookID      string
	WebhookToken   string
	GuildID        string         `gorm:"primaryKey"`
	ChannelID      string         `gorm:"primaryKey"`
	Locale         amqp.Language  `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:Locale"`
	RetryNumber    int64          `gorm:"default:0"`
	UpdatedAt      time.Time
}

type WebhookYoutube struct {
	WebhookID    string
	WebhookToken string
	GuildID      string  `gorm:"primaryKey"`
	ChannelID    string  `gorm:"primaryKey"`
	VideastID    string  `gorm:"primaryKey"`
	Videast      Videast `gorm:"foreignKey:VideastID"`
	RetryNumber  int64   `gorm:"default:0"`
	UpdatedAt    time.Time
}
