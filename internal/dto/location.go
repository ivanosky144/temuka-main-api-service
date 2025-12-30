package dto

type AddLocationRequest struct {
	Name string `json:"name"`
}

type UpdateLocationRequest struct {
	Name string `json:"name"`
}

type LocationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
