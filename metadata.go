package intercom

type Metadata map[string]interface{}

func CreateMetadata() Metadata {
	return make(map[string]interface{})
}

func (m Metadata) Add(key string, value interface{}) Metadata {
	m[key] = value
	return m
}
