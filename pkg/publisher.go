package pkg

import "fmt"

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (p *Publisher) Subscribe(c Consumer) {
	fmt.Println("Subscribed", c.GetName(), "to", p.Name)
	p.Consumers[c.GetName()] = c
}

func (p *Publisher) Unsubscribe(c Consumer) {
	fmt.Println("Unsubscribed", c.GetName(), "from", p.Name)
	delete(p.Consumers, c.GetName())
}

func (p *Publisher) Notify() {
	fmt.Println("Notifying consumers of", p.Name)
	for _, c := range p.Consumers {
		c.Update(p.Name)
	}
}
