package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Shape interface {
	GetId() int
	GetType() string
	SetId() int
}

type Manager struct {
	Items map[string]Cloneable
}

func NewPrototypeManage() *Manager {
	return &Manager{Items: make(map[string]Cloneable)}
}

func (m *Manager) Get(name string) Cloneable {
	c, ok := m.Items[name]
	if !ok {
		return nil
	}
	return c.Clone()
}

func (m *Manager) Set(name string, cloneable Cloneable)  {
	m.Items[name] = cloneable
}


type Circle struct {
	Id int
	Type string
}

func (circle *Circle)GetId() int {
	return circle.Id
}

func (circle *Circle)GetType() string {
	return "circle"
}

func (circle *Circle) SetId(id int) int {
	circle.Id = id
	return circle.Id
}

func (circle *Circle) Clone() Cloneable {
	circle2 := *circle
	return &circle2
}

type Rectangle struct {
	Id int
	Type string
}


func (rectangle *Rectangle)GetId() int {
	return rectangle.Id
}

func (rectangle *Rectangle)GetType() string {
	return "circle"
}

func (rectangle *Rectangle) SetId(id int) int {
	rectangle.Id = id
	return rectangle.Id
}

func (rectangle *Rectangle) Clone() Cloneable {
	rectangle2 := *rectangle
	return &rectangle2
}
