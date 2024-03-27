package listing

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test1() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test1()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// error
/*
В программе определен пользовательский тип ошибки customError, который имеет поле msg типа string и метод Error(),
который возвращает значение поля msg.
В функции test() создается новый блок кода, в котором ничего не происходит.
В функции test() возвращается значение nil, типа *customError.
В функции main() объявляется переменная err типа error.
В переменную err присваивается возвращаемое значение функции test(), которая возвращает nil.
Затем в условии if err != nil проверяется, является ли переменная err равной nil
*/
