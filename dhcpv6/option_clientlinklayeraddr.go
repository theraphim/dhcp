package dhcpv6

import (
	"fmt"
	"net"

	"github.com/u-root/u-root/pkg/uio"
)

// OptClientLinkLayerAddr represents an OptionClientLinkLayerAddr.
//
// This module defines the OptClientLinkLayerAddr structure.
// https://www.ietf.org/rfc/rfc6939.txt
type OptClientLinkLayerAddr struct {
	Addr net.HardwareAddr
}

// Code returns the option's code
func (op *OptClientLinkLayerAddr) Code() OptionCode {
	return OptionClientLinkLayerAddr
}

// ToBytes serializes the option and returns it as a sequence of bytes
func (op *OptClientLinkLayerAddr) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	buf.WriteBytes([]byte{0, 1})
	buf.WriteBytes(op.Addr)
	return buf.Data()
}

func (op *OptClientLinkLayerAddr) String() string {
	return fmt.Sprintf("OptClientLinkLayerAddr{mac=%s}", op.Addr)
}

// ParseOptClientLinkLayerAddr builds an OptClientLinkLayerAddr structure from a sequence
// of bytes. The input data does not include option code and length
// bytes.
func ParseOptClientLinkLayerAddr(data []byte) (*OptClientLinkLayerAddr, error) {
	var opt OptClientLinkLayerAddr
	if len(data) != 8 || data[0] != 0 || data[1] != 1 {
		return nil, fmt.Errorf("can't parse non-mac option 79 yet")
	}
	opt.Addr = append([]byte{}, data[2:]...)
	return &opt, nil
}
