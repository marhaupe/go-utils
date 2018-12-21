package pkg

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil || s.top != -1 || s.el == nil {
		t.Fatal("Received faulty stack")
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top int
		el  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
		want1  bool
	}{
		{
			name: "Pop with nothing inside",
			fields: fields{
				top: -1,
				el:  make([]interface{}, 0, 5),
			},
			want:  nil,
			want1: false,
		},
		{
			name: "Pop with one inside",
			fields: fields{
				top: 0,
				el:  []interface{}{1},
			},
			want:  1,
			want1: true,
		},
		{
			name: "Pop with two different inside",
			fields: fields{
				top: 1,
				el:  []interface{}{1, "test"},
			},
			want:  "test",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top: tt.fields.top,
				el:  tt.fields.el,
			}
			got, got1 := s.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Stack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack(t *testing.T) {
	s := New()

	// Nothing in the Stack
	val, ok := s.Peek()
	if val != nil || ok {
		t.Fatal("Peeking empty stack returned wrong result")
	}

	val, ok = s.Pop()
	if val != nil || ok {
		t.Fatal("Popping empty stack returned wrong result")
	}

	// "Test" in the Stack
	s.Push("Test")
	val, ok = s.Peek()
	if val != "Test" || !ok {
		t.Fatal("Peeking empty stack returned wrong result")
	}

	val, ok = s.Pop()
	if val != "Test" || !ok {
		t.Fatal("Peeking empty stack returned wrong result")
	}

	s.Push("Test")
	s.Push(true)
	s.Push(2)

	// "Test", true, 2 in the Stack
	val, ok = s.Pop()
	if val != 2 || !ok {
		t.Fatal("Popping filled stack returned wrong result")
	}
	// "Test", true in the Stack
	val, ok = s.Pop()
	if val != true || !ok {
		t.Fatal("Popping filled stack returned wrong result")
	}
	// "Test", true, 2 in the Stack
	val, ok = s.Pop()
	if val != "Test" || !ok {
		t.Fatal("Popping filled stack returned wrong result")
	}
	// Nothing in the Stack
	val, ok = s.Peek()
	if val != nil || ok {
		t.Fatal("Peeking empty stack returned wrong result")
	}
	val, ok = s.Pop()
	if val != nil || ok {
		t.Fatal("Popping empty stack returned wrong result")
	}
}
