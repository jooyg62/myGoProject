package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address // 익명 필드, 임베딩된 구조체, Subscriber 구조체로 승격됨
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
