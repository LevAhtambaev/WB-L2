Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Вывод будет следующим: [77 78 79]
```
Это элементы индексом с 1 по 3 включительно
Срез по индексам работает следующим образом - Выбирается полуоткрытый диапазон, который включает первый элемент, но исключает последний.
