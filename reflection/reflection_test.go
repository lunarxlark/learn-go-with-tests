package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
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
			"Pointers",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Tokyo"},
			},
			[]string{"London", "Tokyo"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Tokyo"},
			},
			[]string{"London", "Tokyo"},
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

	// Map
	t.Run("map", func(t *testing.T) {
		data := map[string]string{
			"Foo": "foo",
			"Bar": "bar",
		}
		var got []string
		walk(data, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "foo")
		assertContains(t, got, "bar")
	})

	// Channel
	t.Run("chan", func(t *testing.T) {
		dataChannel := make(chan Profile)

		go func() {
			dataChannel <- Profile{33, "Lunar"}
			dataChannel <- Profile{34, "Lark"}
			close(dataChannel)
		}()

		var got []string
		want := []string{"Lunar", "Lark"}
		walk(dataChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// Function
	t.Run("function", func(t *testing.T) {
		dataFunc := func() (Profile, Profile) {
			return Profile{33, "Lunar"}, Profile{34, "Lark"}
		}

		var got []string
		want := []string{"Lunar", "Lark"}
		walk(dataFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, got []string, want string) {
	t.Helper()
	isContains := false
	for _, x := range got {
		if x == want {
			isContains = true
		}
	}
	if !isContains {
		t.Errorf("expected %v to contain %q, but it didn't", got, want)
	}
}