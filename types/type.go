package types

// swagger:model
type ToyReq struct {
	// the name for this toy
	// required: true
	// min length: 3
	Name string `json:"name"`
	// the price for this toy
	// required: true
	// min: 1
	Price int `json:"price"`
}
