package base

import (
	"context"
	"time"

	"github.com/inter-hubly/pilot/hctx"
)

type Entity struct {
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt,omitempty"`
	RemovedAt time.Time `bson:"removedAt" json:"removedAt,omitempty"`
	CreatedBy string    `bson:"createdBy" json:"createdBy"`
	RemovedBy string    `bson:"removedBy" json:"removedBy,omitempty"`
	TenantId  string    `bson:"tenantId" json:"tenantId,omitempty"`
	Removed   bool      `json:"removed"`
}

func NewBaseEntity(ctx context.Context, logged *hctx.Logged) Entity {
	return Entity{
		CreatedAt: time.Now(),
		TenantId:  logged.Tenant,
		CreatedBy: logged.UserId,
	}
}
