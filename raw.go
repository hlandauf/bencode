package bencode

// RawMessage is a special type that will store the raw bencode data when
// encoding or decoding.
type RawMessage []byte

func (rm RawMessage) MarshalBencode() ([]byte, error) {
	return rm, nil
}

func (rm *RawMessage) UnmarshalBencode(b []byte) error {
	*rm = b
	return nil
}
