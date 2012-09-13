package main

import "fmt"

// PersonInfo 是一个包含详细个人信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)
	
	// 往这个 map 里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	
	// 从这个 map 查找 key 为"1234"的信息
	person, ok := personDB["1234"]
	
	// ok 是一个返回的 bool 型,返回 true 表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}
