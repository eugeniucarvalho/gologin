package facebook

import (
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

const facebookAPI = "https://graph.facebook.com/v2.11/"

// User is a Facebook user.
//
// Note that user ids are unique to each app.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// client is a Facebook client for obtaining the current User.
type client struct {
	c     *http.Client
	sling *sling.Sling
}

func newClient(httpClient *http.Client) *client {
	base := sling.New().Client(httpClient).Base(facebookAPI)
	return &client{
		c:     httpClient,
		sling: base,
	}
}

func (c *client) Me(additionalFields []string) (*User, *http.Response, error) {
	user := new(User)
	var fields string
	if len(additionalFields) > 0 {
		fields = "," + strings.Join(additionalFields[:], ",")
	}
	// Facebook returns JSON as Content-Type text/javascript :(
	// Set Accept header to receive proper Content-Type application/json
	// so Sling will decode into the struct
	resp, err := c.sling.New().SaveResponseBody().Set("Accept", "application/json").Get("me?fields=name,email" + fields).ReceiveSuccess(user)
	return user, resp, err
}
