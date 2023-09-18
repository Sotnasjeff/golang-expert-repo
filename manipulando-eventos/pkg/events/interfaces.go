package events

import "time"

// Criação do Evento
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayloads() interface{}
}

// Criação do Handler que irá lidar com o evento
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// Criação do gerenciador do evento/operação
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
