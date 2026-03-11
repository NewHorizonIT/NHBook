package presentation

import "github.com/gin-gonic/gin"

// RegisterBookRoutes đăng ký các routes cho book module
func RegisterBookRoutes(r *gin.RouterGroup, handler *BookHandler) {
	bookRouter := r.Group("/books")
	{
		// Book routes
		bookRouter.POST("", handler.CreateBook)                           // Create book
		bookRouter.GET("", handler.GetBooks)                             // Get all books with filtering
		bookRouter.GET("/search", handler.SearchBooks)                   // Search books
		bookRouter.GET("/:id", handler.GetBook)                          // Get book by ID
		bookRouter.PUT("/:id", handler.UpdateBook)                       // Update book
		bookRouter.DELETE("/:id", handler.DeleteBook)                    // Delete book
		bookRouter.PATCH("/:id/adjust-stock", handler.AdjustStock)       // Adjust stock
		bookRouter.PATCH("/:id/change-price", handler.ChangePrice)       // Change price
		bookRouter.PATCH("/:id/toggle-active", handler.ToggleActive)     // Toggle active status

		// Author routes
		authorRouter := bookRouter.Group("/authors")
		{
			authorRouter.POST("", handler.CreateAuthor)                   // Create author
			authorRouter.GET("/:id", handler.GetAuthor)                   // Get author by ID
			authorRouter.PUT("/:id", handler.UpdateAuthor)                // Update author
			authorRouter.DELETE("/:id", handler.DeleteAuthor)             // Delete author
		}

		// Category routes
		categoryRouter := bookRouter.Group("/categories")
		{
			categoryRouter.POST("", handler.CreateCategory)               // Create category
			categoryRouter.GET("", handler.GetCategories)                 // Get all categories
			categoryRouter.GET("/:id", handler.GetCategory)               // Get category by ID
			categoryRouter.PUT("/:id", handler.UpdateCategory)            // Update category
			categoryRouter.DELETE("/:id", handler.DeleteCategory)         // Delete category
		}

		// Publisher routes
		publisherRouter := bookRouter.Group("/publishers")
		{
			publisherRouter.POST("", handler.CreatePublisher)             // Create publisher
			publisherRouter.GET("", handler.GetPublishers)                // Get all publishers
			publisherRouter.GET("/:id", handler.GetPublisher)             // Get publisher by ID
			publisherRouter.PUT("/:id", handler.UpdatePublisher)          // Update publisher
			publisherRouter.DELETE("/:id", handler.DeletePublisher)       // Delete publisher
		}
	}
}
