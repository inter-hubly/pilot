package base

import (
	"context"
	"time"

	"github.com/inter-hubly/pilot/hctx"
)

type Entity struct {
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt,omitempty"`
	RemovedAt time.Time `bson:"removedAt" json:"-"`
	CreatedBy string    `bson:"createdBy" json:"-"`
	RemovedBy string    `bson:"removedBy" json:"-"`
	TenantId  string    `bson:"tenantId" json:"-"`
	Removed   bool      `bson:"removed" json:"-"`
}

func NewBaseEntity(ctx context.Context, logged *hctx.Logged) Entity {
	return Entity{
		CreatedAt: time.Now(),
		TenantId:  logged.Tenant,
		CreatedBy: logged.UserId,
	}
}
