package data_link_layer

import (
	"network_practice/common"
	"sync"
)

type (
	// Strictly speaking, this is an Ethernet switch.
	// Exclusive media, collision-free transmission, adapter do not need allow CSMA/CD .
	Ethernet struct {
		AddressTable map[string]int // ethernet change table, key is Mac address value is physical port
		PortMap map[int]*Adapter 	// analog hardware level Ethernet port
		Hosts []Adapter				// host in ethernet
		lock sync.RWMutex			// atomic update change table
	}
)

func NewEthernet() *Ethernet {
	return &Ethernet{
		AddressTable: make(map[string]int, 0),
		PortMap: make(map[int]*Adapter, 0),
		Hosts: make([]Adapter, 0),
	}
}

func (e *Ethernet) DstHosts(srcAdd string) []string  {
	dsts := make([]string, 0)
	for i := 0; i < len(e.Hosts); i++ {
		if e.Hosts[i].MacAdd != srcAdd {
			dsts = append(dsts, e.Hosts[i].MacAdd)
		}
	}
	return dsts
}

func (e *Ethernet) AddHost(macAdd string, port int)  {
	a := NewAdapter(macAdd)
	e.PortMap[port] = a
	e.Hosts = append(e.Hosts, *a)
	go func() {
		for {
			a.ReceiveFrame()		// adapter always receiving data
		}
	}()
}

// Self-learning Ethernet switch
// search host in address table, if there is no element, broadcast on ethernet, return dst port
func (e *Ethernet) SelfLearningEthernetSwitch(port int, dstAdd string) int {
	a := e.PortMap[port]
	if len(e.AddressTable) == 0 {
		e.lock.Lock()
		e.AddressTable[a.MacAdd] = port
		e.lock.Unlock()
	}
	dstPort := e.AddressTable[dstAdd]
	if dstPort == 0 {
		return e.BroadcastOnEthernet(a, port, dstAdd)
	}
	return dstPort
}

// send frame to other hosts except current host
func (e *Ethernet) BroadcastOnEthernet(ad *Adapter, port int, dstAdd string) int  {
	data := common.StringToSlice(dstAdd)
	ad.Broadcast(data, IP)
	for p, a := range e.PortMap {
		if p != port {
			dataBuf, _ := a.ReceiveFrame() // Host compares the received destination address with its own address
			if a.MacAdd == common.SliceToString(dataBuf) {
				ad.Uniast(a.MacAdd, common.StringToSlice(dstAdd), IP)	//dst host uniast to src host
				e.lock.Lock()
				e.AddressTable[ad.MacAdd] = p
				e.lock.Unlock()
				return p
			}
		}
	}
	return 0
}






