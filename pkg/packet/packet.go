package packet

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/allen13/kubepacket/pkg/prom"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Capture - Start capturing packets on an interface
func Capture(iface string, filter string, snaplen int32, flushDuration time.Duration) {
	log.Printf("starting capture on interface %q", iface)
	// Set up pcap packet capture
	handle, err := pcap.OpenLive(iface, snaplen, true, flushDuration/2)
	if err != nil {
		log.Fatal("error opening pcap handle: ", err)
	}
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal("error setting BPF filter: ", err)
	}

	log.Println("reading in packets")

	// We use a DecodingLayerParser here instead of a simpler PacketSource.
	// This approach should be measurably faster, but is also more rigid.
	// PacketSource will handle any known type of packet safely and easily,
	// but DecodingLayerParser will only handle those packet types we
	// specifically pass in.  This trade-off can be quite useful, though, in
	// high-throughput situations.
	var eth layers.Ethernet
	var ip4 layers.IPv4
	var tcp layers.TCP
	var udp layers.UDP

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet,
		&eth, &ip4, &tcp, &udp)
	parser.DecodingLayerParserOptions.IgnoreUnsupported = true
	decoded := make([]gopacket.LayerType, 0, 4)

loop:
	for {
		data, _, err := handle.ZeroCopyReadPacketData()

		if err != nil {
			log.Printf("error getting packet: %v", err)
			continue
		}

		err = parser.DecodeLayers(data, &decoded)
		if err != nil {
			log.Printf("error decoding packet: %v", err)
			continue
		}

		foundNetLayer := false
		for _, typ := range decoded {
			switch typ {
			case layers.LayerTypeIPv4:
				foundNetLayer = true
			case layers.LayerTypeTCP:
				if foundNetLayer {
					recordPacket(
						"tcp",
						ip4.SrcIP,
						ip4.DstIP,
						uint16(tcp.SrcPort),
						uint16(tcp.DstPort))
					continue loop
				}
			case layers.LayerTypeUDP:
				if foundNetLayer {

					recordPacket(
						"udp",
						ip4.SrcIP,
						ip4.DstIP,
						uint16(udp.SrcPort),

						uint16(udp.DstPort))
					continue loop
				}
			}
		}
	}
}

func recordPacket(protocol string, srcIP, dstIP net.IP, srcPort, dstPort uint16) {
	var srcPortStr, dstPortStr string

	if srcPort >= 32768 {
		srcPortStr = "dynamic"
	} else {
		srcPortStr = strconv.FormatUint(uint64(srcPort), 10)
	}

	if dstPort >= 32768 {
		dstPortStr = "dynamic"
	} else {
		dstPortStr = strconv.FormatUint(uint64(dstPort), 10)
	}

	prom.PacketCount.WithLabelValues(protocol, srcIP.String(), srcPortStr, dstIP.String(), dstPortStr).Inc()
}
