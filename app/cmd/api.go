package cmd

import (
	DBConn "Astro/database"
	"Astro/repository"
	AstroRoutes "Astro/routes"
	Token "Astro/token"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	rankingDbOnce sync.Once
	usersDbOnce   sync.Once
)

func (app *app) serveApi() error {
	router := gin.Default()

	rankingDb := DBConn.InitializeDB(
		app.config.rankingDbOpts.user,
		app.config.rankingDbOpts.pass,
		app.config.rankingDbOpts.addr,
		app.config.rankingDbOpts.dbName,
		&rankingDbOnce)

	usersDb := DBConn.InitializeDB(
		app.config.usersDbOpts.user,
		app.config.usersDbOpts.pass,
		app.config.usersDbOpts.addr,
		app.config.usersDbOpts.dbName,
		&usersDbOnce)

	rankingRepo := repository.NewRankingRepository(rankingDb)
	usersRepo := repository.NewUserRepository(usersDb)

	publicRoutes := router.Group("/public")
	authRoutes := router.Group("/auth")        // Routes for Authentication
	authenticatedRoutes := router.Group("/v1") // Authenticated Route
	authenticatedRoutes.Use(Token.AuthenticatedAction())

	//region Routes
	//Public
	AstroRoutes.TokenRoutes(authRoutes)
	AstroRoutes.ApplyUserRouters(usersRepo, authRoutes)
	AstroRoutes.ApplyNewsRoutes(publicRoutes)

	//Authenticated
	AstroRoutes.ApplyRankingRoutes(authenticatedRoutes, rankingRepo)
	//endregion

	return router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
