package counter

import "testing"

func TestCount(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		
		assertCount(t, counter, 3)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func NewCounter() *Counter {
	return &Counter{}
}