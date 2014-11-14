package intercom

type CustomData map[string]interface{}

func CreateCustomData() CustomData {
	return make(map[string]interface{})
}

func (c CustomData) Add(key string, value interface{}) CustomData {
	c[key] = value
	return c
}
