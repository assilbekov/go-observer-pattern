package main

import "observer/pkg"

func main() {
	sub1 := &pkg.Subscriber{Name: "sub1"}
	sub2 := &pkg.Subscriber{Name: "sub2"}
	sub3 := &pkg.Subscriber{Name: "sub3"}
	channel := pkg.Publisher{
		Name: "Channel",
		Consumers: map[string]pkg.Consumer{
			sub1.GetName(): sub1,
		},
	}
	channel.Subscribe(sub3)
	channel.Subscribe(sub2)

	channel.Notify()

	channel.Unsubscribe(sub2)

	channel.Notify()
}
