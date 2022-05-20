package services

type injectable[T interface{}] func(c *Container) T

type Container struct {
	CountService
}

func CreateContainer(
	countService injectable[CountService],
) *Container {

	c := &Container{}

	c.CountService = countService(c)

	return c
}
