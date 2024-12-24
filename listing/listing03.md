Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```

У нас возвращается интерфейс с двумя полями: тип, значение.
Тип у нас PathError, Значение nil.
Проверка "== nil" проверяет оба значения на nil, а у нас не так.
