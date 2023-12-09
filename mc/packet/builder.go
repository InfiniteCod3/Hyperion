package packet

import (
	"bytes"
)

type Builder struct {
	buf bytes.Buffer
}

func (p *Builder) WriteField(fields ...FieldEncoder) {
	w := bufio.NewWriter(&p.buf)
	defer w.Flush()
	for _, f := range fields {
		_, err := f.WriteTo(w)
		if err != nil {
			panic(err)
		}
	}
}

func (p *Builder) Packet(id int32) Packet {
	return Packet{ID: id, Data: p.buf.Bytes()}
}
