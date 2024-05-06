package pkg

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (p *Publisher) Subscribe(c Consumer) {
	p.Consumers[c.GetName()] = c
}

func (p *Publisher) Unsubscribe(c Consumer) {
	delete(p.Consumers, c.GetName())
}

func (p *Publisher) Notify() {
	for _, c := range p.Consumers {
		c.Update(p.Name)
	}
}
