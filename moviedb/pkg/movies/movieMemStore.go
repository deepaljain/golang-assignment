package movies

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type MemStore struct {
	list map[string]Movie
}

func NewMemStore() *MemStore {
	list := make(map[string]Movie)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, movie Movie) error {
	m.list[name] = movie
	return nil
}

func (m MemStore) List() (map[string]Movie, error) {
	return m.list, nil
}

func (m MemStore) Get(name string) (Movie, error) {
	movie, ok := m.list[name]
	if !ok {
		return Movie{}, ErrNotFound
	}
	return movie, nil
}

func (m MemStore) Delete(name string) error {
	_, ok := m.list[name]
	if !ok {
		return ErrNotFound
	}
	delete(m.list, name)
	return nil
}

func (m MemStore) Update(name string, movie Movie) error {
	_, ok := m.list[name]
	if !ok {
		return ErrNotFound
	}
	m.list[name] = movie
	return nil
}