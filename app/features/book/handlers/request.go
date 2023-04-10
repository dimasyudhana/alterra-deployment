package handlers

type BookRequest struct {
	Title     string `json:"title"`
	Year      string `json:"year"`
	Publisher string `json:"publisher"`
}

type UpdateBook struct {
	Title     string `json:"title"`
	Year      string `json:"year"`
	Publisher string `json:"publisher"`
}
