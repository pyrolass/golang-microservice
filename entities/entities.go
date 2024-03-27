package entities

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}
