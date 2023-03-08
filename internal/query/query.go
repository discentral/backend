package query

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/discentral/backend/pkg/config"
)

type Resolver struct{}

func (*Resolver) Authorize() (string, error) {
	u, err := url.Parse("https://discord.com/api/oauth2/authorize")
	if err != nil {
		return "", errors.New("an error occured while parsing authorize url")
	}

	q := u.Query()
	q.Set("client_id", config.Get("DISCORD_CLIENT_ID"))
	q.Set("redirect_uri", config.GetWithFallback("DISCORD_REDIRECT_URI", fmt.Sprintf("%s/auth", config.Get("APP_URL"))))
	q.Set("response_type", "code")
	q.Set("scope", config.GetWithFallback("SCOPE", "identify guilds"))
	u.RawQuery = q.Encode()

	return u.String(), nil
}
