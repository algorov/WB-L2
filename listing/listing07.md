Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Создаются 2 канала и третий канал, в который будут записываться выводы из первых двух.
Будут выводиться числа с 1 до 8, далее бесконечно нули.

Проблема в том, что в merge в констукции select не предусмотрена ситуация, когда каналы закрыты
и поэтому в канал с записываются бесконечно дефолтные значения.

В главной горутине тоже есть неточность, в for не предусмотрена ситуация, когда канала c закрыт.
Поэтому переменная v получает значение по умолчанию, то есть 0.

```
