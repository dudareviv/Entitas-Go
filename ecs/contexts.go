package ecs

type Contexts []interface{}

func SetContexts(ebs ...interface{}) Contexts {
	var contexts Contexts
	for _, v := range ebs {
		contexts = append(contexts, v)
	}
	return contexts
}

func (c *Contexts) Add(element interface{}) {
	*c = append(*c, element)
}

func (c *Contexts) Get(id int) interface{} {
	return (*c)[id]
}
