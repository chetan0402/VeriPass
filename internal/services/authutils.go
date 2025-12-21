package veripass

import "context"

type Claim string

const EmailKey Claim = "email"
const UsernameKey Claim = "name"

func GetCtxWithEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, EmailKey, email)
}

func GetCtxWithUsername(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, UsernameKey, name)
}

func GetEmailFromCtx(ctx context.Context) string {
	return ctx.Value(EmailKey).(string)
}

func GetUsernamefromCtx(ctx context.Context) string {
	return ctx.Value(UsernameKey).(string)
}
