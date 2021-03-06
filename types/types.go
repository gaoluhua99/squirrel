package types

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	ID        snowflake.ID `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"password,omitempty"`
	Locked    bool         `json:"locked"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  time.Time    `json:"update_at"`
}

type UserNoPass struct {
	ID        snowflake.ID `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"-"`
	Locked    bool         `json:"locked"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  time.Time    `json:"update_at"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIN   int    `json:"expires_in"`
}
