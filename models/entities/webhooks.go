package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type WebhookAlmanax struct {
	WebhookId    string
	WebhookToken string
	GuildId      string        `gorm:"primaryKey"`
	ChannelId    string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookFeed struct {
	WebhookId    string
	WebhookToken string
	GuildId      string        `gorm:"primaryKey"`
	ChannelId    string        `gorm:"primaryKey"`
	FeedTypeId   string        `gorm:"primaryKey"`
	Locale       amqp.Language `gorm:"primaryKey"`
	FeedSource   FeedSource    `gorm:"foreignKey:FeedTypeId,Locale"`
	RetryNumber  int64         `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitch struct {
	WebhookId    string
	WebhookToken string
	GuildId      string  `gorm:"primaryKey"`
	ChannelId    string  `gorm:"primaryKey"`
	StreamerId    string  `gorm:"primaryKey"`
	Streamer      Streamer `gorm:"foreignKey:StreamerId"`
	RetryNumber  int64   `gorm:"default:0"`
	UpdatedAt    time.Time
}

type WebhookTwitter struct {
	WebhookId      string
	WebhookToken   string
	GuildId        string         `gorm:"primaryKey"`
	ChannelId      string         `gorm:"primaryKey"`
	Locale         amqp.Language  `gorm:"primaryKey"`
	TwitterAccount TwitterAccount `gorm:"foreignKey:Locale"`
	RetryNumber    int64          `gorm:"default:0"`
	UpdatedAt      time.Time
}

type WebhookYoutube struct {
	WebhookId    string
	WebhookToken string
	GuildId      string  `gorm:"primaryKey"`
	ChannelId    string  `gorm:"primaryKey"`
	VideastId    string  `gorm:"primaryKey"`
	Videast      Videast `gorm:"foreignKey:VideastId"`
	RetryNumber  int64   `gorm:"default:0"`
	UpdatedAt    time.Time
}
