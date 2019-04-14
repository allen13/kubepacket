package kubepacket

//Kubepacket handler
// Start/Stop packet captures using kubewatch events
type Kubepacket struct {
}

// Init handler
func (k *Kubepacket) Init() error {
	return nil
}

// ObjectCreated handler
func (k *Kubepacket) ObjectCreated(obj interface{}) {

}

// ObjectDeleted handler
func (k *Kubepacket) ObjectDeleted(obj interface{}) {

}

// ObjectUpdated handler
func (k *Kubepacket) ObjectUpdated(oldObj, newObj interface{}) {

}
