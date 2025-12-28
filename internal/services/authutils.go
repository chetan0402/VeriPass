// Package veripass
package veripass

import "context"

// Claim type is to be used for defining keys for context.
type Claim string

// EmailKey is the typed key used to get/set value in context.
const EmailKey Claim = "email"

// UsernameKey is the typed key used to get/set value in context.
const UsernameKey Claim = "name"

// GetCtxWithEmail Returns a new context with email set using EmailKey
func GetCtxWithEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, EmailKey, email)
}

// GetCtxWithUsername Returns a new context with name set using UsernameKey
func GetCtxWithUsername(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, UsernameKey, name)
}

// GetEmailFromCtx Returns value of email from context using EmailKey
func GetEmailFromCtx(ctx context.Context) string {
	return ctx.Value(EmailKey).(string)
}

// GetUsernamefromCtx Returns value of username from context using UsernameKey
func GetUsernamefromCtx(ctx context.Context) string {
	return ctx.Value(UsernameKey).(string)
}
