package service

import (
	"errors"
	"reflect"
	"strconv"
)

type PersistentStore interface {
	Get(key string) (any, error)
	Set(key string, value any) error
}

type Service struct {
	Age   int
	store PersistentStore
}

func NewService(st PersistentStore) *Service {
	return &Service{store: st}
}

func (s *Service) GetAge() (int, error) {
	var r int = 0
	v, err := s.store.Get("age")
	if err == nil {
		switch t := reflect.TypeOf(v).Kind(); t {
		case reflect.Int:
			r = v.(int)
		case reflect.String:
			r, err = strconv.Atoi(v.(string))
			if err != nil {
				err = errors.New("value has type string but unable to convert")
			}
		default:
			err = errors.New("invalid value type!")
		}
	}
	return r, err
}

func (s *Service) SetAge(v int) error {
	return s.store.Set("age", v)
}
