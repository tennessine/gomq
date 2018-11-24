package set

type (
	nothing struct{}
	Set struct {
		hash map[interface{}]nothing
	}
)

func New(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]nothing)}
	for _, v := range initial {
		s.Insert(v)
	}
	return s
}

func (s *Set) Insert(elem interface{}) {
	s.hash[elem] = nothing{}
}

func (s *Set) Len() int {
	return len(s.hash)
}

func (s *Set) Remove(elem interface{}) {
	delete(s.hash, elem)
}

func (s *Set) Do(f func(interface{})) {
	for _, v := range s.hash {
		f(v)
	}
}

func (s *Set) Has(elem interface{}) bool {
	_, exists := s.hash[elem]
	return exists
}

func (s *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]nothing)
	for key, _ := range s.hash {
		if !set.Has(key) {
			n[key] = nothing{}
		}
	}
	return &Set{n}
}

func (s *Set) Intersection(set *Set) *Set {
	n := make(map[interface{}]nothing)
	for key, _ := range s.hash {
		if set.Has(key) {
			n[key] = nothing{}
		}
	}
	return &Set{n}
}

func (s *Set) SubsetOf(set *Set) bool {
	if s.Len() > set.Len() {
		return false
	}
	for key, _ := range s.hash {
		if !set.Has(key) {
			return false
		}
	}
	return true
}

func (s *Set) Union(set *Set) *Set {
	n := make(map[interface{}]nothing)
	for key, _ := range s.hash {
		n[key] = nothing{}
	}
	for key, _ := range set.hash {
		n[key] = nothing{}
	}
	return &Set{n}
}