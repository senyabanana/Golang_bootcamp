package usecases

import "warehouse/internal/repository"

type ClientUseСase struct {
	repo repository.ClientRepository
}

func NewClientUsecase() *ClientUseСase {
	return &ClientUseСase{
		repo: repository.NewHTTPRepository(),
	}
}

func (c *ClientUseСase) Get(key string) (string, error) {
	return c.repo.Get(key)
}

func (c *ClientUseСase) Set(key string, value string) error {
	return c.repo.Set(key, value)
}

func (c *ClientUseСase) Delete(key string) error {
	return c.repo.Delete(key)
}

func (c *ClientUseСase) JoinCluster(address, port string) error {
	return c.repo.JoinCluster(address, port)
}
