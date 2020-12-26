package dtos

type MultiID struct {
	Ids []string `json:"ids" binding:"required"`
}


type Pagination struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	Sort         string      `json:"sort"`
	TotalRows    int         `json:"total_rows"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
	FromRow      int         `json:"from_row"`
	ToRow        int         `json:"to_row"`
	Rows         interface{} `json:"rows"`
	Searchs      []Search    `json:"searchs"`
}
type Search struct {
	Column string `json:"column"`
	Action string `json:"action"`
	Query  string `json:"query"`
}
type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations"`
}