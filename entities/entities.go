package entities

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obu_id"`
	Unix  int64   `json:"unix"`
}

type Invoice struct {
	OBUID         int     `json:"obu_id"`
	TotalDistance float64 `json:"total_distance"`
	TotalCost     float64 `json:"total_cost"`
}
