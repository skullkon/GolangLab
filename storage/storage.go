package storage

type Storage struct {
	config *Config
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	return nil
}

func (s *Storage) Close() {

}
