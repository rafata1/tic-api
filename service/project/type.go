package project

type DTOProject struct {
	Name string `json:"name"`
}

type OutputProject struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
