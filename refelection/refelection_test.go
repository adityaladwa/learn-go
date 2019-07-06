package refelection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struch with one string field",
			struct {
				Name string
			}{"Aditya"},
			[]string{"Aditya"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Aditya", "Bangalore"},
			[]string{"Aditya", "Bangalore"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Aditya", 33},
			[]string{"Aditya"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Aditya",
				Profile{33, "Bangalore"},
			},
			[]string{"Aditya", "Bangalore"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Aditya"},
				{34, "Test"},
			},
			[]string{"Aditya", "Test"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	expected := "Aditya"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("Wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("Got '%s', want '%s'", got[0], expected)
	}
}
