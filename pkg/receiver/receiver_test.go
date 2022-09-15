package receiver

import (
	"errors"
	"testing"
)

type ReaderMock struct{}

func (r *ReaderMock) Read(p []byte) (n int, err error) {
	copy(p[:0], []byte("hello, world\n"))
	return 10, nil
}

type BufioReaderMock struct {
	ret_bytes      []byte
	ret_bytes_err  error
	ret_string     string
	ret_string_err error
}

func (r *BufioReaderMock) ReadBytes(delim byte) ([]byte, error) {
	return r.ret_bytes, r.ret_bytes_err
}

func (r *BufioReaderMock) ReadString(delim byte) (string, error) {
	return r.ret_string, r.ret_string_err
}

func Test_SocketReceiver_NewSockerReceiver(t *testing.T) {
	mock_connection := ReaderMock{}
	receiver := NewSocketReceiver(&mock_connection)
	if receiver.Delim != '\n' {
		t.Errorf("receiver.Delim = %v; want '\\n'", receiver.Delim)
	}
	// if receiver.reader != &mock_connection {
	// 	t.Errorf("receiver.connection = %v; want '%v'", receiver.connection, mock_connection)
	// }
}

func Test_SocketReceiver_ReadBytes_ReturnsProperResponse_NoError(t *testing.T) {
	mock_bytes := []byte("hello, world")
	receiver := SocketReceiver{
		reader: &BufioReaderMock{
			ret_bytes: mock_bytes,
		},
		Delim: '\n',
	}
	res, err := receiver.ReadBytes()
	if string(res) != string(mock_bytes) {
		t.Errorf("Wrong response, got: %v want: %v", res, mock_bytes)
	}
	if err != nil {
		t.Errorf("Wrong err, got: %v want: nil", err)
	}
}

func Test_SocketReceiver_ReadBytes_ReturnsNoBytes_WithError(t *testing.T) {
	mock_err := errors.New("units errors")
	receiver := SocketReceiver{
		reader: &BufioReaderMock{
			ret_bytes:     []byte{},
			ret_bytes_err: mock_err,
		},
		Delim: '\n',
	}
	res, err := receiver.ReadBytes()
	if string(res) != "" {
		t.Errorf("Wrong response, got: %v want: \"\"", res)
	}
	if err != mock_err {
		t.Errorf("Wrong err, got: %v want: %v", err, mock_err)
	}
}

func Test_SocketReceiver_ReadString_ReturnsProperResponse_NoError(t *testing.T) {
	mock_string := "hello, world"
	receiver := SocketReceiver{
		reader: &BufioReaderMock{
			ret_string: mock_string,
		},
		Delim: '\n',
	}
	res, err := receiver.ReadString()
	if res != mock_string {
		t.Errorf("Wrong response, got: %v want: %v", res, mock_string)
	}
	if err != nil {
		t.Errorf("Wrong err, got: %v want: nil", err)
	}
}

func Test_SocketReceiver_ReadString_ReturnsNoBytes_WithError(t *testing.T) {
	mock_err := errors.New("units errors")
	receiver := SocketReceiver{
		reader: &BufioReaderMock{
			ret_string:     "",
			ret_string_err: mock_err,
		},
		Delim: '\n',
	}
	res, err := receiver.ReadString()
	if string(res) != "" {
		t.Errorf("Wrong response, got: %v want: \"\"", res)
	}
	if err != mock_err {
		t.Errorf("Wrong err, got: %v want: %v", err, mock_err)
	}
}
