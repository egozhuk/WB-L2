Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```

test - "х" именованный возвращаемый результат.
Перед возвратом выполнится отложенная функция defer

anotherTest - "x" локальное значение, которое возвращается.
Из-за "return x" значение возвращается немедленно,
а потом уже отложенная функция defer, которая не влияет на значение,
так как уже отправлено
