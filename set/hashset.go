package set

type HashSet struct {
	hash map[interface{}]interface{}
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[interface{}]interface{})}
}

func (s *HashSet) Add(val interface{}) {
	s.hash[val] = struct{}{}
}
