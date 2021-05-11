package main

type TopicTree struct {
	m map[string]interface{}
}

func NewTopicTree() *TopicTree {
	return &TopicTree{
		m: make(map[string]interface{}),
	}
}
func (s *TopicTree) Add(key string, value interface{}) {
	if s.Contains(key) == false {
		s.m[key] = value
	}
}
func (s *TopicTree) Remove(key string) {
	delete(s.m, key)
}
func (s *TopicTree) Contains(key string) bool {
	_, exist := s.m[key]
	return exist
}
func (s *TopicTree) Get(key string) interface{} {
	value, exist := s.m[key]
	if exist {
		return value
	}
	return nil
}
