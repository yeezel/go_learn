package basics

import (
	"fmt"
	"sort"
)

func TestMapDemo() {

	/* 声明变量，默认 map 是 nil */
	// var map_variable map[key_data_type]value_data_type

	/* 直接初始化map */
	// mapLit := map[string]int{"one": 1, "two": 2}
	/* 使用 make 函数创建map */
	// map_variable := make(map[key_data_type]value_data_type)

	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	/*如果确定是真实的,则存在,否则不存在 */
	if capital, ok := countryCapitalMap["American"]; ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	//delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
	delete(countryCapitalMap, "France1")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	for country, name := range countryCapitalMap {
		fmt.Println(country, "首都是", name)
	}

	// map的切片，必须使用两次 make() 函数，第一次分配切片，第二次分配切片中每个 map 元素
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	//排序需要将 key（或者 value）拷贝到一个切片，再对切片排序
	keys := make([]string, len(countryCapitalMap))
	i := 0
	for k, _ := range countryCapitalMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Print("sorted:")
	for _, k := range keys {
		fmt.Printf("%v=%v / ", k, countryCapitalMap[k])
	}
	println()
}
