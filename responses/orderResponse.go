package responses

type OrderResponse struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	OrderID  int      `json:"orderId"`
}
