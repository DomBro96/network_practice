package internet_layer

type (
	IpDataGram struct {
		Header IpDataGramHeader		// fixed part 20 bytes, but
		Data []byte
	}
	IpDataGramHeader struct {
		Version 		int8				// 4 bits, ipv4 or ipv6
		HeaderLength 	int8				// 4 bits, 32 bits as header length unit
		Length 			int16				// byte as length unit
		Identification  int16				// Each time a datagram is generated, the counter is incremented by one.
		Flag 			int8 				// 3 bits, lowest digit MF (more fragment) is 1 means there is also a datagram such as 001， middle digit DF (don't fragment) such as 010
		SliceOffset     int16   			// 13 bits, eight bytes as offset unit, offset in src datagram / 8
		TTL    		    int8 				// time to live: In hops as a unit, how many routers can go though
		Protocol 	    int8
		HeaderCheckSum  int16				// only check header, not include data part
		SrcIPAdd	    int32				// src ip address
		DstIPAdd		int32				// dst ip address
	}
)