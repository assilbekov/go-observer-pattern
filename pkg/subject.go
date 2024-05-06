package pkg

type Subject interface {
	Subscribe(Consumer)
	Unsubscribe(Consumer)
	Notify()
}
