package facebook

import (
	"context"
	"fmt"
)

// unexported key type prevents collisions
type key int

const (
	userKey    key = iota
	rawJSONKey key = iota
)

// WithUser returns a copy of ctx that stores the Facebook User.
func WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// WithRawJSON stores a copy of the response body
func WithRawJSON(ctx context.Context, json string) context.Context {
	return context.WithValue(ctx, rawJSONKey, json)
}

// UserFromContext returns the Facebook User from the ctx.
func UserFromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil, fmt.Errorf("facebook: Context missing Facebook User")
	}
	return user, nil
}

// JSONFromContext returns the Facebook User from the ctx.
func JSONFromContext(ctx context.Context) (string, error) {
	json, ok := ctx.Value(rawJSONKey).(string)
	if !ok {
		return "", fmt.Errorf("facebook: Context missing raw JSON")
	}
	return json, nil
}
