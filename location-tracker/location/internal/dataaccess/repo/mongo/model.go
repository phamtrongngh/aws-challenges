package mongorepo

import "time"

type EntityType string

const (
	Device EntityType = "DEVICE"
)

type Metadata struct {
	CarSku string `json:"carSku"`
	Entity Entity `json:"entity"`
}

type Entity struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Coordinates struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Location struct {
	Metadata  Metadata    `bson:"metadata"`
	Timestamp time.Time   `bson:"timestamp"`
	Coords    Coordinates `bson:"coordinates"`
	Speed     string      `bson:"speedOverGround"`
}
