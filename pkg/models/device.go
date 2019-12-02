package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UID        string             `json:uid"`
	Name       string             `json:"name"`
	Identity   map[string]string  `json:"identity"`
	Attributes map[string]string  `json:"attributes"`
	PublicKey  string             `json:"public_key" bson:"public_key"`
	TenantID   string             `json:"tenant_id" bson:"tenant_id"`
	LastSeen   time.Time          `json:"last_seen" bson:"last_seen"`
	Online     bool               `json:"online" bson:"-,omitempty"`
}

type DeviceAuthRequest struct {
	Identity   map[string]string `json:"identity"`
	Attributes map[string]string `json:"attributes"`
	PublicKey  string            `json:"public_key"`
	TenantID   string            `json:"tenant_id"`
	Sessions   []string          `json:"sessions,omitempty"`
}
