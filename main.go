package main

import (
	"fmt"
	"github.com/Stafford1986/test_manticore/adapter/manticore"
	"github.com/Stafford1986/test_manticore/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	router := gin.Default()
	db, err := sqlx.Open("mysql", fmt.Sprintf("@tcp(%s:%d)/", "127.0.0.1", 9306))
	if err != nil {
		log.Fatal(err)
	}
	resumeRepo := manticore.NewResumeRepository(db)
	vacancyRepo := manticore.NewVacancyRepository(db)

	resumeUseCase := usecase.NewResumeUseCase(resumeRepo)
	vacancyUseCase := usecase.NewVacancyUseCase(vacancyRepo)
	resumeHandler := NewHandler(resumeUseCase, vacancyUseCase)
	initRoutes(router, resumeHandler)

	err = router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
