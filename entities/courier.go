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

type RateByCoordinateResponse struct {
	Success     bool                                `json:"success"`
	Object      string                              `json:"object"`
	Message     string                              `json:"message"`
	Code        int64                               `json:"code"`
	Origin      RateByCoordinateResponseOrigin      `json:"origin"`
	Stops       []any                               `json:"stops"`
	Destination RateByCoordinateResponseDestination `json:"destination"`
	Pricing     []RateByCoordinateResponsePricing   `json:"pricing"`
}

type RateByCoordinateResponseOrigin struct {
	LocationID                        any    `json:"location_id"`
	Latitude                          string `json:"latitude"`
	Longitude                         string `json:"longitude"`
	PostalCode                        any    `json:"postal_code"`
	AdministrativeDivisionLevel1_Name any    `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1_Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2_Name any    `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2_Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3_Name any    `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3_Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4_Name any    `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4_Type string `json:"administrative_division_level_4_type"`
	Address                           any    `json:"address"`
}

type RateByCoordinateResponseDestination struct {
	LocationID                        any    `json:"location_id"`
	Latitude                          string `json:"latitude"`
	Longitude                         string `json:"longitude"`
	PostalCode                        any    `json:"postal_code"`
	AdministrativeDivisionLevel1_Name any    `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1_Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2_Name any    `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2_Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3_Name any    `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3_Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4_Name any    `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4_Type string `json:"administrative_division_level_4_type"`
	Address                           any    `json:"address"`
}

type RateByCoordinateResponsePricing struct {
	AvailableCollectionMethod    []string `json:"available_collection_method"`
	AvailableForCashOnDelivery   bool     `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool     `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillID bool     `json:"available_for_instant_waybill_id"`
	AvailableForInsurance        bool     `json:"available_for_insurance"`
	Company                      string   `json:"company"`
	CourierName                  string   `json:"courier_name"`
	CourierCode                  string   `json:"courier_code"`
	CourierServiceName           string   `json:"courier_service_name"`
	CourierServiceCode           string   `json:"courier_service_code"`
	Currency                     string   `json:"currency"`
	Description                  string   `json:"description"`
	Duration                     string   `json:"duration"`
	ShipmentDurationRange        string   `json:"shipment_duration_range"`
	ShipmentDurationUnit         string   `json:"shipment_duration_unit"`
	ServiceType                  string   `json:"service_type"`
	ShippingType                 string   `json:"shipping_type"`
	Price                        int64    `json:"price"`
	TaxLines                     []any    `json:"tax_lines"`
	Type                         string   `json:"type"`
}

type OrderByCoordinate struct {
	ShipperContactName      string                 `json:"shipper_contact_name"`
	ShipperContactPhone     string                 `json:"shipper_contact_phone"`
	ShipperContactEmail     string                 `json:"shipper_contact_email"`
	ShipperOrganization     string                 `json:"shipper_organization"`
	OriginContactName       string                 `json:"origin_contact_name"`
	OriginContactPhone      string                 `json:"origin_contact_phone"`
	OriginAddress           string                 `json:"origin_address"`
	OriginNote              string                 `json:"origin_note"`
	OriginCoordinate        OriginCoordinate       `json:"origin_coordinate"`
	DestinationContactName  string                 `json:"destination_contact_name"`
	DesinationContactPhone  string                 `json:"destination_contact_phone"`
	DestinationContactEmail string                 `json:"destination_contact_email"`
	DestinationAddress      string                 `json:"destination_address"`
	DestinationNote         string                 `json:"destination_note"`
	DestinationCoordinate   DestinationCoordinate  `json:"destination_coordinate"`
	CourierCompany          string                 `json:"courier_company"`
	CourierType             string                 `json:"courier_type"`
	CourierInsurance        string                 `json:"courier_insurance"`
	DeliveryType            string                 `json:"delivery_type"`
	OrderNote               string                 `json:"order_note"`
	MetaData                any                    `json:"metadata"`
	Items                   OrderByCoordinateItems `json:"items"`
}

type OriginCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DestinationCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OrderByCoordinateItems struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Value       int    `json:"value"`
	Quantity    int    `json:"quantity"`
	Weight      int    `json:"weight"`
}
