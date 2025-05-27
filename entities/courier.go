package entities

import "time"

type Courier struct {
	Success  bool       `json:"success"`
	Object   string     `json:"object"`
	Couriers []Couriers `json:"couriers"`
}

type Couriers struct {
	AvailableCollectionMethod    []AvailableCollectionMethod `json:"available_collection_method"`
	AvailableForCashOnDelivery   bool                        `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool                        `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillId bool                        `json:"available_for_instant_waybill_id"`
	CourierName                  string                      `json:"courier_name"`
	CourierCode                  string                      `json:"courier_code"`
	CourierServiceName           string                      `json:"courier_service_name"`
	CourierServiceCode           string                      `json:"courier_service_code"`
	Tier                         string                      `json:"premium"`
	Description                  string                      `json:"description"`
	ServiceType                  string                      `json:"service_type"`
	ShippingType                 string                      `json:"shipping_type"`
	ShipmentDurationRange        string                      `json:"shipment_duration_range"`
	ShipmentDurationUnit         string                      `json:"shipment_duration_unit"`
}

type AvailableCollectionMethod string

type CourierRate struct {
	OriginLatitude       string `json:"origin_latitude"`
	OriginLangitude      string `json:"origin_longitude"`
	DestinationLatitude  string `json:"destination_latitude"`
	DestinationLongitude string `json:"destination_longitude"`
}

type CreateLocation struct {
	Name         string  `json:"name"`
	ContactName  string  `json:"contact_name"`
	ContactPhone string  `json:"contact_phone"`
	Address      string  `json:"address"`
	Note         string  `json:"note"`
	PostalCode   int     `json:"postal_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Type         string  `json:"type"`
	UserId       string  `json:"user_id"`
}

type CreateLocationResponse struct {
	Success      bool      `json:"success"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Note         string    `json:"note"`
	PostalCode   int       `json:"postal_code"`
	ContactName  string    `json:"contact_name"`
	ContactPhone string    `json:"contact_phone"`
	Owned        bool      `json:"owned"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RateByCoordinate struct {
	OriginLatitude       string                  `json:"origin_latitude"`
	OriginLongitude      string                  `json:"origin_longitude"`
	DestinationLatitude  string                  `json:"destination_latitude"`
	DestinationLongitude string                  `json:"destination_longitude"`
	Couriers             string                  `json:"couriers"`
	Items                []RateByCoordinateItems `json:"items"`
}

type RateByCoordinateItems struct {
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
	Value    int    `json:"value"`
	Quantity int    `json:"quantity"`
}