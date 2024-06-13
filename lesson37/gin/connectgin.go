package gin

type Handler struct {
	User *psql.UserRepo
}

func ConnectGin(handler Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/users", handler.GetAllUser)
	return r
}
