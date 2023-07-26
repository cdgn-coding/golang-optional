package golang_optional

import (
	"encoding/json"
)

type Optional[T any] struct {
	value    T
	hasValue bool
}

func (s *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.hasValue = false
		return nil
	}
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	s.value = t
	s.hasValue = true
	return nil
}

func (s Optional[T]) MarshalJSON() ([]byte, error) {
	if s.IsEmpty() {
		return []byte("null"), nil
	}
	return json.Marshal(s.value)
}

func NewEmpty[T any]() Optional[T] {
	return Optional[T]{}
}

func (s *Optional[T]) IsEmpty() bool {
	return !s.hasValue
}

func (s *Optional[T]) Value() (T, bool) {
	return s.value, s.hasValue
}

func NewOptionalOf[T any](val T) Optional[T] {
	return Optional[T]{value: val, hasValue: true}
}
