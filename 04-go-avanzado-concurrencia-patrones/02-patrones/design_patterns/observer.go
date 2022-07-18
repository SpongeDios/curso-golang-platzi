package main

import "fmt"

type Topic interface {
	registrar(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

//Item -> No disponible
//El item deberia avisar a todos los observadores que esta disponible

type Item struct {
	observers []Observer
	name string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name:      name,
	}
}

func (i *Item) UpdateAvailable()  {
	fmt.Printf("Item %s is available \n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) registrar(observer Observer)  {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast()  {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(value string)  {
	fmt.Printf("Sending email = %s available from client %s\n", value, e.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.registrar(firstObserver)
	nvidiaItem.registrar(secondObserver)
	nvidiaItem.UpdateAvailable()
}
