package services

type CountService interface {
	Hits() int
	Increment()
}

type countService struct {
	container *Container
	hits      int
}

func (c *countService) Hits() int {
	return c.hits
}

func (c *countService) Increment() {
	c.hits++
}

func NewCountService(c *Container) CountService {
	return &countService{
		container: c,
	}
}
