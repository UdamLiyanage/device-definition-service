package main

type DeviceDefinition struct {
	ID            string            `json:"_id,omitempty" bson:"_id,omitempty"`
	Version       string            `json:"version" bson:"version"`
	Name          string            `json:"name" bson:"name"`
	Type          string            `json:"type" bson:"type"`
	Organisations []interface{}     `json:"organisations,omitempty" bson:"organisation,omitempty"`
	Users         []interface{}     `json:"users,omitempty" bson:"users,omitempty"`
	Events        []string          `json:"events,omitempty" bson:"events,omitempty"`
	Actions       []string          `json:"actions,omitempty" bson:"events,omitempty"`
	Parameters    map[string]string `json:"parameters,omitempty" bson:"omitempty"`
}
