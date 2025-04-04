package testutils

import (
	"context"

	"github.com/inter-hubly/pilot/hctx"
)

func SetLoggedUser(ctx context.Context) context.Context {
	tenantId := "tenant"
	ctx = hctx.Tenant.New(tenantId)
	return hctx.LoggedUser.Set(ctx, hctx.Logged{
		UserId:   "userId",
		Username: "username",
		Tenant:   tenantId,
	})
}
