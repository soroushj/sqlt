package sqlt

import (
	"testing"

	"github.com/google/uuid"
)

func TestNullUUIDScan(t *testing.T) {
	su := "11111111-1111-1111-1111-111111111111"
	u := uuid.MustParse(su)
	cases := []struct {
		in   interface{}
		ok   bool
		want NullUUID
	}{
		{nil, true, NullUUID{UUID: uuid.UUID{}, Valid: false}},
		{"", true, NullUUID{UUID: uuid.UUID{}, Valid: false}},
		{make([]byte, 0), true, NullUUID{UUID: uuid.UUID{}, Valid: false}},
		{su, true, NullUUID{UUID: u, Valid: true}},
		{"1", false, NullUUID{}},
		{[]byte(su), true, NullUUID{UUID: u, Valid: true}},
		{[]byte("1"), false, NullUUID{}},
		{u[:], true, NullUUID{UUID: u, Valid: true}},
		{1, false, NullUUID{}},
	}
	for _, c := range cases {
		got := NullUUID{}
		err := got.Scan(c.in)
		if (err == nil) != c.ok {
			t.Errorf("case: %v err: %v", c, err)
		} else if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestNullUUIDValue(t *testing.T) {
	su := "11111111-1111-1111-1111-111111111111"
	u := uuid.MustParse(su)
	cases := []struct {
		in   NullUUID
		want *string
	}{
		{NullUUID{UUID: uuid.UUID{}, Valid: false}, nil},
		{NullUUID{UUID: u, Valid: true}, &su},
	}
	for _, c := range cases {
		got, err := c.in.Value()
		if err != nil {
			t.Errorf("case: %v err: %v", c, err)
		} else if got == nil {
			if c.want != nil {
				t.Errorf("case: %v got: %v", c, got)
			}
		} else if s, ok := got.(string); !ok || s != *c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
