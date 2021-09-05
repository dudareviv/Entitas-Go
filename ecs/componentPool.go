package ecs

type componentPool []Component

func (s *componentPool) Push(e Component) {
	*s = append(*s, e)
}

func (s *componentPool) Pop() (Component, bool) {
	length := len(*s)
	if length > 0 {
		last := length - 1
		entity := (*s)[last]
		*s = (*s)[:last]
		return entity, true
	}
	return nil, false
}
