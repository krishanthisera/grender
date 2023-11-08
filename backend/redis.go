package backend

type Redis struct {
	Chanel string
}

func (r Redis) Put() error {
	// Redis put logic
	return nil
}

func (r Redis) Get() ([]byte, error) {
	// Redis get logic
	return nil, nil
}
