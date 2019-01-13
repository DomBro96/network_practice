package data_link_layer

type (
	// Strictly speaking, this is an Ethernet switch.
	// Exclusive media, collision-free transmission, adapter do not need allow CSMA/CD .
	Ethernet struct {
		AddressTable map[string]int // ethernet change table, key is Mac address value is physical port
		Hosts []Adapter				// host in ethernet
		Ports []int					// physical ports
	}
)

func NewEthernet(s int) *Ethernet  {
	ports := make([]int, s)
	for i := 0; i < s; i++ {
		ports[i] = i
	}
	return &Ethernet{
		AddressTable: make(map[string]int, 0),
		Hosts: make([]Adapter, 0),
		Ports: ports,
	}
}



func (e *Ethernet) AddHost(macAdd string, port int)  {
	a := NewAdapter(macAdd)
	e.Hosts = append(e.Hosts, *a)
	select {
	case <- a.Receiver :
		
	}
}


// strictly, uniast, multicast broadcast are adapter features, but if there is no ethernet
// there is no practical meaning.
func (e *Ethernet) Uniast(srcAdd string, dstAdd string)   {

}

func (e *Ethernet) Multicast() {
	
}

func (e *Ethernet) Broadcast()  {

}







