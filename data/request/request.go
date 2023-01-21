package request

type DoctorSearchRequest struct {
	ids        []int    `json:"Ids"`
	Names      []string `json:"names"`
	Speciality []string `json:"speciality"`
	Rating     []int    `json:"ratings"`
	SortOrder  string   `json:"sort_order"`
	SortBy     string   `json:"sort_by"`
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
}
