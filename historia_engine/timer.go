package historia_engine

import "time"

type Timer struct {
	startTime time.Time
}

func (t *Timer) Restart() time.Duration {
	elapsedTime := time.Now().Sub(t.startTime)
	t.startTime = time.Now()
	return elapsedTime
}

func (t *Timer) GetElapsedTime() time.Duration {
	return time.Now().Sub(t.startTime)
}
