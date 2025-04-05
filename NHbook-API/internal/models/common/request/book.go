package request

type CreateBookRequest struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	ImageURL    string   `json:"imageUrl"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	CategoryID  uint     `json:"categoryId"`
	PublishedAt string   `json:"publishedAt"`
}
