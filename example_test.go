package nestedmap

import "fmt"

func ExampleNestedMapIntIntInt() {
	nm := NewNestedMap[int, int, int]()
	nm.Set(1, 2, 3)

	ret1, ok := nm.GetOuter(1)
	fmt.Println(ret1)
	fmt.Println(ok)

	ret2, ok := nm.GetOuter(100)
	fmt.Println(ret2)
	fmt.Println(ok)

	ret3, ok := nm.GetInner(1, 2)
	fmt.Println(ret3)
	fmt.Println(ok)

	ret4, ok := nm.GetInner(1, 100)
	fmt.Println(ret4)
	fmt.Println(ok)

	// Output:
	// map[2:3]
	// true
	// map[]
	// false
	// 3
	// true
	// 0
	// false
}

func ExampleNestedMapStringIntBool() {
	nm := NewNestedMap[string, int, bool]()
	nm.Set("gopher", 1, true)

	ret1, ok := nm.GetOuter("gopher")
	fmt.Println(ret1)
	fmt.Println(ok)

	ret2, ok := nm.GetOuter("java")
	fmt.Println(ret2)
	fmt.Println(ok)

	ret3, ok := nm.GetInner("gopher", 1)
	fmt.Println(ret3)
	fmt.Println(ok)

	ret4, ok := nm.GetInner("gopher", 100)
	fmt.Println(ret4)
	fmt.Println(ok)

	// Output:
	// map[1:true]
	// true
	// map[]
	// false
	// true
	// true
	// false
	// false
}
