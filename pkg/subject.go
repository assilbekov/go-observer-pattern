package pkg

type Subject interface {
	Subscribe(Consumer)
	Unsubscribe()
	Notify()
}
