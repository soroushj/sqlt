package sqlt

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNullRawMessageScan(t *testing.T) {
	cases := []struct {
		in   interface{}
		ok   bool
		want NullRawMessage
	}{
		{nil, true, NullRawMessage{RawMessage: json.RawMessage{}, Valid: false}},
		{"1", true, NullRawMessage{RawMessage: json.RawMessage("1"), Valid: true}},
		{[]byte("1"), true, NullRawMessage{RawMessage: json.RawMessage("1"), Valid: true}},
		{1, false, NullRawMessage{}},
	}
	for _, c := range cases {
		got := NullRawMessage{}
		err := got.Scan(c.in)
		if (err == nil) != c.ok {
			t.Errorf("case: %v err: %v", c, err)
		} else if !bytes.Equal(got.RawMessage, c.want.RawMessage) || got.Valid != c.want.Valid {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestNullRawMessageValue(t *testing.T) {
	cases := []struct {
		in   NullRawMessage
		want []byte
	}{
		{NullRawMessage{RawMessage: json.RawMessage{}, Valid: false}, nil},
		{NullRawMessage{RawMessage: json.RawMessage("1"), Valid: true}, []byte("1")},
	}
	for _, c := range cases {
		got, err := c.in.Value()
		if err != nil {
			t.Errorf("case: %v err: %v", c, err)
		} else if got == nil {
			if c.want != nil {
				t.Errorf("case: %v got: %v", c, got)
			}
		} else if b, ok := got.([]byte); !ok || !bytes.Equal(b, c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
