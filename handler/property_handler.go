type PropertyHandler struct {
	propertyUseCase *usecase.PropertyUseCase
}

func NewPropertyHandler(r *gin.Engine, pu *usecase.PropertyUseCase, middleware *middleware.AuthMiddleware) {
	handler := &PropertyHandler{
		propertyUseCase: pu,
	}

	propertyRouter := r.Group("/api/v1/properties")
	propertyRouter.Use(middleware.AuthRequired())
	{
		propertyRouter.POST("/", handler.CreateProperty)
		propertyRouter.GET("/:id", handler.GetProperty)
		propertyRouter.PUT("/:id", handler.UpdateProperty)
		propertyRouter.DELETE("/:id", handler.DeleteProperty)
		propertyRouter.GET("/", handler.ListProperties)
	}
}