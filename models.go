package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	DeviceSchema struct {
		ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Version        string             `json:"version" bson:"version"`
		Name           string             `json:"name" bson:"name"`
		UID            string             `json:"uid" bson:"uid"`
		Type           string             `json:"type" bson:"type"`
		Public         bool               `json:"public" bson:"public"`
		Description    string             `json:"description" bson:"description"`
		LatestFirmware string             `json:"latest_firmware_version" bson:"latest_firmware_version"`
		DeviceComm     Communication      `json:"communication" bson:"communication"`
		Commands       []Command          `json:"commands,omitempty" bson:"commands,omitempty"`
		CommandsFormat CommandFormat      `json:"command_format" bson:"command_format"`
		Parameters     []Parameter        `json:"parameters" bson:"parameters"`
		CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
	}

	Communication struct {
		Protocol    string      `json:"protocol" bson:"protocol"`
		Credentials Credentials `json:"credentials" bson:"credentials"`
	}

	Command struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	}

	CommandFormat struct {
		Type    string            `json:"type" bson:"type"`
		Payload map[string]string `json:"payload" bson:"payload"`
	}

	Credentials struct {
		URL            string            `json:"url" bson:"url"`
		Authentication map[string]string `json:"authentication" bson:"authentication"`
	}

	Parameter struct {
		Name      string `json:"name" bson:"name"`
		Parameter string `json:"parameter" bson:"parameter"`
		DataType  string `json:"data_type" bson:"data_type"`
	}
)
