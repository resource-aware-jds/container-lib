package model

type Task struct {
	ID         string
	Attributes TaskAttributes
}

type TaskAttributes map[string]interface{}

func (t TaskAttributes) Get(key string) (interface{}, bool) {
	value, ok := t[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func GetTaskAttributeWithSpecificType[T any](attribute TaskAttributes, key string) (*T, bool) {
	value, ok := attribute.Get(key)
	if !ok {
		return nil, false
	}

	parsedValue, ok := value.(T)
	if !ok {
		return nil, false
	}

	return &parsedValue, true
}
