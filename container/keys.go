package container

import (
	"sort"
	"reflect"
	"bytes"
	"fmt"
)

type CompareFunction func(interface{}, interface{}) int8

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	CompareFunc() CompareFunction
	ElemKind() reflect.Kind
}

type myKeys struct {
	Keys
	container   []interface{}
	compareFunc CompareFunction
	elemKind    reflect.Kind
}

func (keys *myKeys) Len() int {
	return len(keys.container)
}

func (keys *myKeys) Less(i, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

func (keys *myKeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil || reflect.TypeOf(k).Kind() != keys.elemKind {
		return false
	}
	return true
}

func (keys *myKeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container, k)
	sort.Sort(keys)
	return true
}

func (keys *myKeys) Remove(k interface{}) bool {
	if index, contains := keys.Search(k); !contains {
		return false
	} else {
		keys.container = append(keys.container[0:index], keys.container[index+1:]...)
		return true
	}
}

func (keys *myKeys) Clear() {
	keys.container = make([]interface{}, 0)
}

func (keys *myKeys) Get(index int) interface{} {
	if index >= keys.Len() {
		return nil
	}
	return keys.container[index]
}

func (keys *myKeys) GetAll() []interface{} {
	initialLen := len(keys.container)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, key := range keys.container {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (keys *myKeys) Search(k interface{}) (index int, contains bool) {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return
	}
	index = sort.Search(keys.Len(), func(i int) bool {
		return keys.compareFunc(keys.container[i], k) >= 0
	})
	if index < keys.Len() && keys.container[index] == k {
		contains = true
	}
	return
}

func (keys *myKeys) ElemKind() reflect.Kind {
	return keys.elemKind
}

func (keys *myKeys) CompareFunc() CompareFunction {
	return keys.compareFunc
}

func (keys *myKeys) String() string {
	var buf bytes.Buffer
	buf.WriteString("Keys<")
	buf.WriteString(keys.elemKind.String())
	buf.WriteString(">{")
	first := true
	buf.WriteString("[")
	for _, key := range keys.container {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("]")
	buf.WriteString("}")
	return buf.String()
}

func NewKeys(
	compareFunc func(interface{}, interface{}) int8,
	elemKind reflect.Kind) Keys {
	return &myKeys{
		container:   make([]interface{}, 0),
		compareFunc: compareFunc,
		elemKind:    elemKind,
	}
}

func DefaultCompare(e1 interface{}, e2 interface {}) int8 {
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		return 0
	}
	k1 := reflect.ValueOf(e1).Kind()
	switch k1 {
	case reflect.Int:
		if e1.(int) < e2.(int) {
			return -1
		} else if e1.(int) > e2.(int) {
			return 1
		} else {
			return 0
		}
	case reflect.Int8:
		if e1.(int8) < e2.(int8) {
			return -1
		} else if e1.(int8) > e2.(int8) {
			return 1
		} else {
			return 0
		}
	case reflect.Int16:
		if e1.(int16) < e2.(int16) {
			return -1
		} else if e1.(int16) > e2.(int16) {
			return 1
		} else {
			return 0
		}
	case reflect.Int32:
		if e1.(int32) < e2.(int32) {
			return -1
		} else if e1.(int32) > e2.(int32) {
			return 1
		} else {
			return 0
		}
	case reflect.Int64:
		if e1.(int64) < e2.(int64) {
			return -1
		} else if e1.(int64) > e2.(int64) {
			return 1
		} else {
			return 0
		}
	case reflect.Float32:
		if e1.(float32) < e2.(float32) {
			return -1
		} else if e1.(float32) > e2.(float32) {
			return 1
		} else {
			return 0
		}
	case reflect.Float64:
		if e1.(float64) < e2.(float64) {
			return -1
		} else if e1.(float64) > e2.(float64) {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		if e1.(string) < e2.(string) {
			return -1
		} else if e1.(string) > e2.(string) {
			return 1
		} else {
			return 0
		}
	case reflect.Bool:
		if e1.(bool) == e2.(bool) {
			return 0
		} else if e1.(bool) == false && e2.(bool) == true {
			return 1
		} else {
			return 0
		}
	case reflect.Uint:
		if e1.(uint) < e2.(uint) {
			return -1
		} else if e1.(uint) > e2.(uint) {
			return 1
		} else {
			return 0
		}
	case reflect.Uint8:
		if e1.(uint8) < e2.(uint8) {
			return -1
		} else if e1.(uint8) > e2.(uint8) {
			return 1
		} else {
			return 0
		}
	case reflect.Uint16:
		if e1.(uint16) < e2.(uint16) {
			return -1
		} else if e1.(uint16) > e2.(uint16) {
			return 1
		} else {
			return 0
		}
	case reflect.Uint32:
		if e1.(uint32) < e2.(uint32) {
			return -1
		} else if e1.(uint32) > e2.(uint32) {
			return 1
		} else {
			return 0
		}
	case reflect.Uint64:
		if e1.(uint64) < e2.(uint64) {
			return -1
		} else if e1.(uint64) > e2.(uint64) {
			return 1
		} else {
			return 0
		}
	}
	return 0
}