package main

import "github.com/gin-gonic/gin"

func initRoutes(engine *gin.Engine, handler *RequestsHandler) {
	engine.POST("/resume/create_index", handler.createResumeIndexHandler())
	engine.PUT("/resume/update_index", handler.updateResumeIndexHandler())
	engine.GET("/resume/search", handler.createSearchResumeHandler())
	engine.GET("/resume/suggest", handler.createResumeSuggestHandler())

	engine.POST("/vacancy/create_index", handler.createVacancyIndexHandler())
	engine.PUT("/vacancy/update_index", handler.updateVacanciesIndexHandler())
	engine.GET("/vacancy/search", handler.createSearchVacancyHandler())
	engine.GET("/vacancy/suggest", handler.createVacancySuggestHandler())
}
