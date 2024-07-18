package model

/* Request body:
{
	"info": {
		"car_sku": "MG5_001",
		"entityId": {
			"entityType": "DEVICE",
			"id": "4beac0f0-eb6e-11ee-9c0c-a7e0377c6340"
		}
	},
	"value": [
		{
		"ts": 1715184377186,
		"longitude": "10743.594002",
		"speedOverGround": "0",
		"latitude": "1038.587898"
		}
	]
}
*/

type LocationUpdateValue struct {
	Timestamp       int64  `json:"ts" example:"1622547800"`
	Longitude       string `json:"longitude" example:"10743.594002"`
	Latitude        string `json:"latitude" example:"1038.587898"`
	SpeedOverGround string `json:"speedOverGround" example:"0"`
}

type Entity struct {
	Id   string `json:"id" example:"4beac0f0-eb6e-11ee-9c0c-a7e0377c6340"`
	Type string `json:"entityType" example:"DEVICE"`
}

type Info struct {
	CarSku string `json:"car_sku" example:"MG5_001"`
	Entity Entity `json:"entityId"`
}

type LocationUpdate struct {
	Info   Info                  `json:"info"`
	Values []LocationUpdateValue `json:"value"`
}

type LocationFilter struct {
	Limit *int64 `form:"limit"`
}
