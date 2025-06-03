package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	helper "superapps/helpers"
	models "superapps/models"
)

func CourierList() (map[string]any, error) {
	url := os.Getenv("URL_BITESHIP") + "/v1/couriers"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		helper.Logger("error", "In Server: Failed to create request - "+err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTHORIZATION_BITESHIP"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.Logger("error", "In Server: Failed to send request - "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.Logger("error", "In Server: API returned status "+resp.Status)
		return nil, err
	}

	var courier models.Courier

	if err := json.NewDecoder(resp.Body).Decode(&courier); err != nil {
		helper.Logger("error", "In Server: Failed to decode response - "+err.Error())
		return nil, err
	}

	return map[string]any{
		"data": courier,
	}, nil
}

func CreateLocation(cl *models.CreateLocation) (map[string]any, error) {
	url := os.Getenv("URL_BITESHIP") + "/v1/locations"

	payloadData := map[string]any{
		"name":          cl.Name,
		"contact_name":  cl.ContactName,
		"contact_phone": cl.ContactPhone,
		"address":       cl.Address,
		"note":          cl.Note,
		"postal_code":   cl.PostalCode,
		"latitude":      cl.Latitude,
		"longitude":     cl.Longitude,
		"type":          cl.Type,
	}

	payload, err := json.Marshal(payloadData)

	if err != nil {
		helper.Logger("error", "Failed to marshal request body - "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		helper.Logger("error", "In Server: Failed to create request - "+err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTHORIZATION_BITESHIP"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.Logger("error", "In Server: Failed to send request - "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.Logger("error", "In Server: API returned status "+resp.Status)
		return nil, err
	}

	return map[string]any{
		"data": "",
	}, nil
}

func RateByCoordinate(rbc *models.RateByCoordinate) (map[string]any, error) {
	url := os.Getenv("URL_BITESHIP") + "/v1/rates/couriers"

	payloadData := map[string]any{
		"origin_latitude":       rbc.OriginLatitude,
		"origin_longitude":      rbc.OriginLongitude,
		"destination_latitude":  rbc.DestinationLatitude,
		"destination_longitude": rbc.DestinationLongitude,
		"couriers":              rbc.Couriers,
		"items":                 rbc.Items,
	}

	payload, err := json.Marshal(payloadData)

	if err != nil {
		helper.Logger("error", "Failed to marshal request body - "+err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		helper.Logger("error", "In Server: Failed to create request - "+err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTHORIZATION_BITESHIP"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.Logger("error", "In Server: Failed to send request - "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.Logger("error", "In Server: API returned status "+resp.Status)
		return nil, err
	}

	var rateByCoordinate models.RateByCoordinateResponse

	if err := json.NewDecoder(resp.Body).Decode(&rateByCoordinate); err != nil {
		helper.Logger("error", "In Server: Failed to decode response - "+err.Error())
		return nil, err
	}

	return map[string]any{
		"data": rateByCoordinate,
	}, nil
}

func OrderByCoordinate(obc *models.OrderByCoordinate) (map[string]any, error) {
	url := os.Getenv("URL_BITESHIP") + "/v1/orders"

	payloadData := map[string]any{
		"shipper_contact_name":  obc.ShipperContactName,
		"shipper_contact_phone": obc.ShipperContactPhone,
		"shipper_contact_email": obc.ShipperContactEmail,
		"shipper_organization":  obc.ShipperOrganization,
		"origin_contact_name":   obc.OriginContactName,
		"origin_contact_phone":  obc.OriginContactPhone,
		"origin_address":        obc.OriginAddress,
		"origin_note":           obc.OriginNote,
		"origin_coordinate": map[string]float64{
			"latitude":  obc.OriginCoordinate.Latitude,
			"longitude": obc.OriginCoordinate.Longitude,
		},
		"destination_contact_name":  obc.DestinationContactName,
		"destination_contact_phone": obc.DesinationContactPhone,
		"destination_contact_email": obc.DestinationContactEmail,
		"destination_address":       obc.DestinationAddress,
		"destination_note":          obc.DestinationNote,
		"destination_coordinate": map[string]float64{
			"latitude":  obc.DestinationCoordinate.Latitude,
			"longitude": obc.DestinationCoordinate.Longitude,
		},
		"courier_company": obc.CourierCompany,
		"courier_type":    obc.CourierType,
		"delivery_type":   obc.DeliveryType,
		"order_note":      obc.OrderNote,
		"items":           obc.Items,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("AUTHORIZATION_BITESHIP"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		helper.Logger("error", "In Server: Failed to send request - "+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.Logger("error", "In Server: API returned status "+resp.Status)
		return nil, err
	}

	var orderByCoordinateResponse models.OrderByCoordinateResponse

	if err := json.NewDecoder(resp.Body).Decode(&orderByCoordinateResponse); err != nil {
		helper.Logger("error", "In Server: Failed to decode response - "+err.Error())
		return nil, err
	}
	return map[string]any{
		"data": orderByCoordinateResponse,
	}, nil
}
