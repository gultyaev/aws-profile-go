package lib

type DbCollection[T interface{}] struct {
	Collection string
	Value      T
}
