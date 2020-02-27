package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type DeviceDefinition struct {
	ID         primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Version    string              `json:"version" bson:"version"`
	Name       string              `json:"name" bson:"name"`
	UID        string              `json:"uid" bson:"uid"`
	Type       string              `json:"type" bson:"type"`
	Events     []string            `json:"events,omitempty" bson:"events,omitempty"`
	Actions    []map[string]string `json:"actions,omitempty" bson:"actions,omitempty"`
	Parameters map[string]string   `json:"parameters,omitempty" bson:"omitempty"`
}
