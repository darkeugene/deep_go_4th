package main

import "errors"

type Container struct {
	registered map[string]func() interface{}
}

func NewContainer() *Container { // создать DI контейнер
	return &Container{
		registered: make(map[string]func() interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) { // зарегистрировать конструктор по созданию типа
	c.registered[name] = constructor.(func() interface{})
}

func (c *Container) RegisterSingletonType(name string, constructor interface{}) {
	singleton := constructor.(func() interface{})()
	c.registered[name] = func() interface{} { return singleton }
}

func (c *Container) Resolve(name string) (interface{}, error) { // создать объект с использованием конструктора
	if constructor, ok := c.registered[name]; ok {
		return constructor(), nil
	}

	return nil, errors.New("no constructor found")
}
