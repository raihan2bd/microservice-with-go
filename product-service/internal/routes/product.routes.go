package routes

func (r *Routes) productRoutes() {
	r.mux.GET("/products", r.handlers.GetProductByID)
}
