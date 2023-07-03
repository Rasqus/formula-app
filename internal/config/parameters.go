package config

// swagger:parameters GetRequest
type GetRequest struct {
	// ID of an item
	// in: path
	ID int `json:"id"`
}

// swagger:parameters DeleteRequest
type DeleteRequest struct {
	// ID of an item
	// in: path
	ID int `json:"id"`
}

// swagger:parameters PostDriverRequest
type PostDriverRequest struct {
	// Name of the Driver
	// in: string
	Name string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	// Surname of the Driver
	// in: string
	Surname string `json:"surname" validate:"required"`
	// Country of the Driver
	// in: string
	Country string `json:"country" validate:"required"`
	// StartNumber of the Driver
	// in: int64
	StartNumber int `json:"start_number"  validate:"required"`
	// Team of the Driver
	// in: string
	Team string `json:"team"  validate:"required"`
}

// swagger:parameters PostDriverRequestBody
type PostDriverRequestBody struct {
	// name: driver
	// in: body
	// description: new driver
	// schema:
	// type: object
	// "$ref": "#/definitions/PostDriverRequest"
	// required: true
	Body PostDriverRequest `json:"body"`
}

// swagger:parameters PutDriverRequest
type PutDriverRequest struct {
	// ID of the Driver
	// in: int64
	ID int `json:"id"`
	// Name of the Driver
	// in: string
	Name string `json:"name"`
	// Surname of the Driver
	// in: string
	Surname string `json:"surname"`
	// Country of the Driver
	// in: string
	Country string `json:"country"`
	// StartNumber of the Driver
	// in: int64
	StartNumber int `json:"start_number"`
	// Team of the Driver
	// in: string
	Team string `json:"team"`
}

// swagger:parameters PutDriverRequestBody
type PutDriverRequestBody struct {
	// name: driver
	// in: body
	// description: update driver
	// schema:
	// type: object
	// "$ref": "#/definitions/PutDriverRequest"
	// required: true
	Body PutDriverRequest `json:"body"`
}
