package goim

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

type Message struct {
	Length       [4]byte
	HeaderLength [2]byte
	ProtoVersion [2]byte
	Ops          [4]byte
	SequenceID   [4]byte
	Body         []byte
}

func Encode(m *Message) ([]byte, error) {
	// Calc GOIM message length
	var length = int32(len(m.Body) + 16)
	var pkg = new(bytes.Buffer)

	// Write length into header
	if err := binary.Write(pkg, binary.BigEndian, length); err != nil {
		return nil, err
	}
	// Write Header length into header, which is at fixed size of 16 bytes
	if err := binary.Write(pkg, binary.BigEndian, int16(16)); err != nil {
		return nil, err
	}
	// Write Proto Version into header
	if err := binary.Write(pkg, binary.BigEndian, m.ProtoVersion); err != nil {
		return nil, err
	}
	// Write Ops into header
	if err := binary.Write(pkg, binary.BigEndian, m.Ops); err != nil {
		return nil, err
	}
	// Write Sequence ID into header
	if err := binary.Write(pkg, binary.BigEndian, m.SequenceID); err != nil {
		return nil, err
	}
	// Write message into body
	if err := binary.Write(pkg, binary.BigEndian, m.Body); err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(reader io.Reader) (*Message, error) {
	m := &Message{}
	if err := binary.Read(reader, binary.BigEndian, &m.Length); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &m.HeaderLength); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &m.ProtoVersion); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &m.Ops); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &m.SequenceID); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &m.Body); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Message) String() string {
	return fmt.Sprintf("length:%d header_length:%d proto_ver:%d ops:%s sequence_id:%d msg:%s",
		m.Length,
		m.HeaderLength,
		m.ProtoVersion,
		m.Ops,
		m.SequenceID,
		m.Body,
	)
}

// GOIM scanner split function
func SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	log.Printf(string(data))
	if !atEOF {
		if len(data) > 16 { // if received data bytes is more than 16 bytes
			length := int16(0)
			binary.Read(bytes.NewReader(data[:4]), binary.BigEndian, &length)
			dataLen := int(length)
			if dataLen <= len(data) {
				return dataLen, data[:dataLen], nil
			}
		}
	}
	return
}
