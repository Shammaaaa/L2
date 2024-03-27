package pattern

type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

type Place interface {
	Accept(v Visitor) string
}

type People struct {
}

func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

type BurgerBar struct {
}

func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

/*
Паттерн "Посетитель" применяется, когда необходимо добавить новую операцию (функциональность) к группе объектов без изменения самих объектов.
Он позволяет разделить алгоритмы от структуры объектов, что упрощает добавление новых операций и обеспечивает открытую гибкую систему.

Плюсы использования паттерна "Посетитель" включают:

1. Расширяемость: Паттерн позволяет добавлять новые операции (посетители) к существующим классам, не изменяя их самостоятельно.
Это позволяет легко добавлять новую функциональность к объектам, не нарушая принцип открытости/закрытости.
2. Разделение алгоритмов и структур: Паттерн позволяет разделить алгоритмы и операции над объектами от самих объектов, что делает систему более гибкой и модульной.
Это упрощает добавление новых операций и алгоритмов без изменения существующих структур.
3. Полиморфизм: Паттерн "Посетитель" использует полиморфизм для различных типов объектов и методов посетителей,
что позволяет обрабатывать объекты в зависимости от их типа и выполнять соответствующие операции.

Однако, есть и некоторые минусы:

1. Сложность добавления новых классов: При добавлении нового класса в структуру объектов может потребоваться внести изменения в интерфейсы и методы посетителей.
Это может быть неудобно и сложно при работе со сложными иерархиями объектов.
2. Нарушение инкапсуляции: Внедрение посетителя может нарушить инкапсуляцию объектов, поскольку посетитель получает доступ к внутреннему состоянию объектов.
Это может привести к увеличению связанности между объектами.

Необходимость использования паттерна "Посетитель" следует оценивать в контексте конкретной задачи. Он особенно полезен,
когда требуется добавить новую функциональность к группе объектов без изменения самих объектов или когда структура объектов стабильна,
но операции над ними могут изменяться.
*/
