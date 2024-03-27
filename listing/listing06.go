package listing

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s) // [3 2 3]
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

/*
В программе объявляется функция main(), в которой создается слайс s типа []string и инициализируется значениями "1", "2", "3".
Затем вызывается функция modifySlice(s), и в качестве аргумента передается слайс s.
Внутри функции modifySlice(), параметр i также является слайсом []string, который получает значение слайса s.
В строке i[0] = "3" значение первого элемента слайса i изменяется на "3". Таким образом, слайс i будет иметь значения ["3", "2", "3"].
Затем в строке i = append(i, "4") происходит добавление элемента "4" в конец слайса i.
Слайс i становится ["3", "2", "3", "4"]. Но это изменение не отражается на слайсе s, так как была создана копия слайса s внутри функции.
Далее в строке i[1] = "5" значение второго элемента слайса i изменяется на "5". После этого значения слайса i становятся ["3", "5", "3", "4"].
В строке i = append(i, "6") элемент "6" добавляется в конец слайса i. Теперь слайс i становится ["3", "5", "3", "4", "6"].
Затем программа выводит слайс s с помощью fmt.Println(s). В результате получаем [3 2 3], так как только первое изменение значения элемента произошло в слайсе s,
а все последующие изменения касаются копии слайса s, созданной при передаче в функцию modifySlice().
*/
