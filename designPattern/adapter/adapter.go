package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

type AdapteeImpl struct {
}

func (adapteeImpl *AdapteeImpl) SpecificRequest() string {
	return "adaptee specific method"
}

type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		Adaptee: adaptee,
	}
}