package container

import (
	"reflect"
	"bytes"
	"fmt"
)

type OrderedMap interface {
	Map
	//获取第一个键值，如果没有任何键值则返回nil
	FirstKey() interface{}
	//获取最后一个键值，如果没有任何键值对则返回nil
	LastKey() interface{}
	//获取由小于键值对toKey的键值对应的键值对组成的OrderedMap 类型值
	HeadMap(toKey interface{}) OrderedMap
	//获取由小于键值对toKey且大于等于键值fromKey的键值所对应的键值对组成的OrderedMap的类型值
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	//获取由大于等于键值fromKey的键值所对应的键值对组成的OrderedMap的类型值
	TailMap(fromKey interface{}) OrderedMap
}

type myOrderedMap struct {
	OrderedMap
	keys     Keys
	elemKind reflect.Kind
	m        map[interface{}]interface{}
}

func (omap *myOrderedMap) Get(key interface{}) interface{} {
	return omap.m[key]
}

func (omap *myOrderedMap) isAcceptableElem(e interface{}) bool {
	if e == nil {
		return false
	}
	if reflect.TypeOf(e).Kind() != omap.elemKind {
		return false
	}
	return true
}

func (omap *myOrderedMap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !omap.isAcceptableElem(elem) {
		return nil, false
	}
	oldElem, ok := omap.m[key]
	omap.m[key] = elem
	if !ok {
		omap.keys.Add(key)
	}
	return oldElem, true
}

func (omap *myOrderedMap) Remove(key interface{}) interface{} {
	oldElem, ok := omap.m[key]
	delete(omap.m, key)
	if ok {
		omap.keys.Remove(key)
	}
	return oldElem
}

func (omap *myOrderedMap) Clear() {
	omap.m = make(map[interface{}]interface{})
	omap.keys.Clear()
}

func (omap *myOrderedMap) Len() int {
	return len(omap.m)
}

func (omap *myOrderedMap) Contains(key interface{}) bool {
	_, ok := omap.m[key]
	return ok
}

func (omap *myOrderedMap) FirstKey() interface{} {
	if omap.Len() == 0 {
		return nil
	}
	return omap.keys.Get(0)
}

func (omap *myOrderedMap) LastKey() interface{} {
	length := omap.Len()
	if length == 0 {
		return nil
	}
	return omap.keys.Get(length - 1)
}

func (omap *myOrderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMap {
	newOmap := &myOrderedMap{
		keys:     NewKeys(omap.keys.CompareFunc(), omap.keys.ElemKind()),
		elemKind: omap.elemKind,
		m:        make(map[interface{}]interface{}),
	}
	omapLen := omap.Len()
	if omapLen == 0 {
		return newOmap
	}
	beginIndex, contains := omap.keys.Search(fromKey)
	if !contains {
		beginIndex = 0
	}
	endIndex, contains := omap.keys.Search(toKey)
	if !contains {
		endIndex = omapLen
	}
	var key, elem interface{}
	for i := beginIndex; i < endIndex; i ++ {
		key = omap.keys.Get(i)
		elem = omap.m[key]
		newOmap.Put(key, elem)
	}
	return newOmap
}

func (omap *myOrderedMap) HeadMap(toKey interface{}) OrderedMap {
	return omap.SubMap(nil, toKey)
}

func (omap *myOrderedMap) TailMap(fromKey interface{}) OrderedMap {
	return omap.SubMap(fromKey, nil)
}

func (omap *myOrderedMap) Keys() []interface{} {
	initialLen := omap.keys.Len()
	keys := make([]interface{}, initialLen)
	actualLen := 0
	for _, key := range omap.keys.GetAll() {
		if actualLen < initialLen {
			keys[actualLen] = key
		} else {
			keys = append(keys, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		keys = keys[:actualLen]
	}
	return keys
}

func (omap *myOrderedMap) Elems() []interface{} {
	initialLen := omap.Len()
	elems := make([]interface{}, initialLen)
	actualLen := 0
	for _, key := range omap.keys.GetAll() {
		elem := omap.m[key]
		if actualLen < initialLen {
			elems[actualLen] = elem
		} else {
			elems = append(elems, elem)
		}
		actualLen++
	}
	if actualLen < initialLen {
		elems = elems[:actualLen]
	}
	return elems
}

func (omap *myOrderedMap) ToMap() map[interface{}]interface{} {
	replica := make(map[interface{}]interface{})
	for k, v := range omap.m {
		replica[k] = v
	}
	return replica
}

func (omap *myOrderedMap) KeyKind() reflect.Kind {
	return omap.keys.ElemKind()
}

func (omap *myOrderedMap) ElemKind() reflect.Kind {
	return omap.elemKind
}

func (omap *myOrderedMap) String() string {
	var buf bytes.Buffer
	buf.WriteString("OrderedMap<")
	buf.WriteString(omap.keys.ElemKind().String())
	buf.WriteString(",")
	buf.WriteString(omap.elemKind.String())
	buf.WriteString(">{")
	first := true
	omapLen := omap.Len()
	for i := 0; i < omapLen; i++ {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		key := omap.keys.Get(i)
		buf.WriteString(fmt.Sprintf("%v", key))
		buf.WriteString(":")
		buf.WriteString(fmt.Sprintf("%v", omap.m[key]))
	}
	buf.WriteString("}")
	return buf.String()
}

func NewOrderedMap(keys Keys, elemKind reflect.Kind) OrderedMap {
	return &myOrderedMap{
		keys:     keys,
		elemKind: elemKind,
		m:        make(map[interface{}]interface{})}
}
