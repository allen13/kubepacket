package handlers

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init() error {
	return nil
}

func (d *Default) ObjectCreated(obj interface{}) {

}

func (d *Default) ObjectDeleted(obj interface{}) {

}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {

}
