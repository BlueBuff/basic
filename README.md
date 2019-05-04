
```
func main() {
	keys1 := container.NewKeys(container.DefaultCompare, reflect.Int)
	keys1.Add(5)
	keys1.Add(2)
	keys1.Add(3)
	keys1.Add(1)
	keys1.Add(4)
	keys1.Add(6)
	fmt.Println(keys1)

	keys2 := container.NewKeys(container.DefaultCompare, reflect.String)
	keys2.Add("e")
	keys2.Add("f")
	keys2.Add("a")
	keys2.Add("b")
	keys2.Add("g")
	keys2.Add("b")
	fmt.Println(keys2)

	cmap := container.NewConcurrentMap(reflect.String,reflect.Int)
	cmap.Put("a",1)
	cmap.Put("b",2)
	cmap.Put("c",3)
	fmt.Println(cmap)
	fmt.Println("get:",cmap.Get("a"))
	val:=cmap.Remove("a")
	fmt.Println("val:",val)
	fmt.Println("get:",cmap.Get("a"))

	omap:=container.NewOrderedMap(container.NewKeys(container.DefaultCompare,reflect.Int),reflect.String)

	omap.Put(1,"a")
	omap.Put(6,"b")
	omap.Put(5,"c")
	omap.Put(3,"d")

	fmt.Println(omap)

	omap.Remove(1)

	fmt.Println(omap)
}
```