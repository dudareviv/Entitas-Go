package ecs

type entityPool []int

func (s *entityPool) Push(e int) {
	*s = append(*s, e)
}

func (s *entityPool) Pop() (int, bool) {
	length := len(*s)
	if length > 0 {
		last := length - 1
		entity := (*s)[last]
		*s = (*s)[:last]
		return entity, true
	}
	return 0, false
}
