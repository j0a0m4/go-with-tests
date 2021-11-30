package walk

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Name          string
	Input         interface{}
	ExpectedCalls []string
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []TestCase{
		{
			Name: "Struct with only one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name:          "Struct with nested fields",
			Input:         Person{"Chris", Profile{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "Pointers to things",
			Input:         &Person{"Chris", Profile{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
