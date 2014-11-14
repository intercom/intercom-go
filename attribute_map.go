package intercom

type AttributeMap map[string]interface{}

func (am AttributeMap) Add(key string, value interface{}) AttributeMap {
	am[key] = value
	return am
}

func CreateCustomData() AttributeMap {
	return make(AttributeMap)
}

func CreateMetadata() AttributeMap {
	return make(AttributeMap)
}
