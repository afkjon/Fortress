package tutorial

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Matt"},
			result: "Hello, Matt!",
		},
		{
			items:  []string{"Matt", "Josh"},
			result: "Hello, Matt, Josh!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s",
				st.result, st.items, s)
		}
	}
}
