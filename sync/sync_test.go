package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t *testing.T, c Counter, times int) {
	t.Helper()

	if c.Value() != times {
		t.Errorf("got %d, but want %d", c.value, times)
	}
}
