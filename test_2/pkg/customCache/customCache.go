package customCache

type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
	Clear()
}
