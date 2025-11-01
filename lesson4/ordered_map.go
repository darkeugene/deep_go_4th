package main

import "cmp"

type Element[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Left  *Element[K, V]
	Right *Element[K, V]
}
type OrderedMap[K cmp.Ordered, V any] struct {
	root          *Element[K, V]
	elementsCount int
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] { // создать упорядоченный словарь
	return OrderedMap[K, V]{
		root:          nil,
		elementsCount: 0,
	}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) { // добавить элемент в словарь
	current, parent := m.findKeyAndParent(key)

	if current != nil {
		current.Value = value

		return
	}

	m.elementsCount++

	if parent == nil {
		m.root = &Element[K, V]{
			Key:   key,
			Value: value,
		}

		return
	}

	if parent.Key > key {
		parent.Left = &Element[K, V]{
			Key:   key,
			Value: value,
		}

		return
	}

	parent.Right = &Element[K, V]{
		Key:   key,
		Value: value,
	}
}

func (m *OrderedMap[K, V]) Erase(key K) { // удалить элемент из словари
	current, parent := m.findKeyAndParent(key)

	if current == nil {
		return
	}

	m.elementsCount--

	if current.Right == nil {
		if parent == nil {
			m.root = current.Left

			return
		}

		if parent.Left == current {
			parent.Left = current.Left
		} else {
			parent.Right = current.Left
		}

		return
	}

	leftest, leftestParent := findLeftestAndParent(current.Right)
	current.Key = leftest.Key
	current.Value = leftest.Value

	if leftestParent == nil {
		current.Right = leftest.Right
	} else {
		leftestParent.Left = leftest.Right
	}

}

func (m *OrderedMap[K, V]) Contains(key K) bool { // проверить существование элемента в словаре
	x, _ := m.findKeyAndParent(key)

	if x == nil {
		return false
	}

	return true
}

func (m *OrderedMap[K, V]) Size() int { // получить количество элементов в словаре
	return m.elementsCount
	//count := 0
	//recursiveWalkAndCall(m.root, func(_, _ int) {
	//	count++
	//})
	//
	//return count
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	recursiveWalkAndCall(m.root, action)
}

func recursiveWalkAndCall[K cmp.Ordered, V any](root *Element[K, V], action func(K, V)) {
	if root == nil {
		return
	}

	recursiveWalkAndCall(root.Left, action)

	action(root.Key, root.Value)

	recursiveWalkAndCall(root.Right, action)
}

func (m *OrderedMap[K, V]) findKeyAndParent(key K) (*Element[K, V], *Element[K, V]) {
	var parent *Element[K, V]

	current := m.root

	for current != nil {
		if current.Key == key {
			break
		}

		parent = current

		if current.Key > key {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return current, parent
}

func findLeftestAndParent[K cmp.Ordered, V any](root *Element[K, V]) (*Element[K, V], *Element[K, V]) {
	var parent *Element[K, V]

	current := root

	for current.Left != nil {
		parent = current
		current = current.Left
	}

	return current, parent
}
