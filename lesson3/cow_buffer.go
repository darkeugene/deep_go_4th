package main

import (
	"runtime"
	"sync"
	"unsafe"
)

type COWBuffer struct {
	data    []byte
	counter *int
	mx      *sync.RWMutex
}

func NewCOWBuffer(data []byte) COWBuffer { // создать буффер с определенными данными
	counter := 1

	cowBuffer := &COWBuffer{
		data:    data,
		counter: &counter,
		mx:      &sync.RWMutex{},
	}

	runtime.SetFinalizer(cowBuffer, (*COWBuffer).Close)

	return *cowBuffer
}
func (b *COWBuffer) Clone() COWBuffer { // создать новую копию буфера
	b.mx.Lock()
	defer b.mx.Unlock()

	*b.counter++

	return COWBuffer{
		data:    b.data,
		counter: b.counter,
		mx:      b.mx,
	}
}

func (b *COWBuffer) Close() { // перестать использовать копию буффера
	b.mx.Lock()
	defer b.mx.Unlock()

	*b.counter--

	zeroCounter := 0
	b.data = []byte{}
	b.counter = &zeroCounter
}

func (b *COWBuffer) Update(index int, value byte) bool { // изменить определенный байт в буффере
	b.mx.Lock()
	defer b.mx.Unlock()

	if *b.counter == 0 || index < 0 || index >= len(b.data) {
		return false
	}

	if *b.counter > 1 {
		*b.counter--

		*b = NewCOWBuffer(b.data)

		newData := make([]byte, len(b.data))
		copy(newData, b.data)
		b.data = newData
	}

	b.data[index] = value

	return true
}

func (b *COWBuffer) String() string { // сконвертировать буффер в строку
	b.mx.RLock()
	defer b.mx.RUnlock()

	return *(*string)(unsafe.Pointer(&b.data))
}
