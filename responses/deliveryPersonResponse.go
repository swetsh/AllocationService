package responses

type Location struct {
	GeoCoordinate string `json:"geoCoordinate"`
}

type DeliveryPerson struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	OrderID  int      `json:"orderId"`
}
