Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Будет выведено error, поскольку функция test() возвращает не пустой интерфейс. 
В переменной err будет храниться nil, однако тип данных будет указывать на customError.
Поэтому err != nil.
```
