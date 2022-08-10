package codec

import (
	"bufio"
	"encoding/json"
	"io"
)

type JsonCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *json.Decoder
	enc  *json.Encoder
}

func NewJsonCodec(conn io.ReadWriteCloser) *JsonCodec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn: conn,
		buf:  buf,
		dec:  json.NewDecoder(conn),
		enc:  json.NewEncoder(buf),
	}
}

func (j *JsonCodec) Write(data any) error {
	defer func() {
		_ = j.buf.Flush()
	}()
	if err := j.enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func (j *JsonCodec) Read(data any) error {
	return j.dec.Decode(data)
}

func (j *JsonCodec) Close() error {
	return j.conn.Close()
}
