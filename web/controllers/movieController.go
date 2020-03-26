package controllers

import (
	"github.com/kataras/iris/mvc"
	"imooc-iris/repository"
	"imooc-iris/services"
)

type MovieController struct {
	
}

func (c *MovieController) Get() mvc.View{
	movieRepository := repository.NewMovieManger()
	movie := services.NewMovieServiceManager(movieRepository)
	name := movie.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: name,
	}
}
