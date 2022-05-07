package models

type DbCollection[T interface{}] struct {
	Collection string
	Value      T
}
