package services

import (
	"imooc-iris/repository"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
    repo repository.MovieRepository
}

func NewMovieServiceManager(repo repository.MovieRepository) MovieService{
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string{
	return "获取的名称：" + m.repo.GetMovieName()
}

