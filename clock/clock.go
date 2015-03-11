package clock

type Master struct {
	slaves []*Slave
}

func New() (m *Master) {
    return &Master{}
}

func (m *Master) Spawn(divider uint) (s *Slave) {
	s = &Slave{divider, 0, make(chan struct{})}
	m.slaves = append(m.slaves, s)
	return s
}

func (m *Master) Tick(ticks uint) {
	for _, s := range m.slaves {
		s.Tick(ticks)
	}
}

type Slave struct {
	divider uint
	count uint
	ch chan struct{}
}

func (s *Slave) Tick(ticks uint) {
	end := s.count + ticks
	ticks = end / s.divider
	s.count = end % s.divider
	
	for i := uint(0); i < ticks; i++ {
		s.ch <- struct{}{}
	}
}

func (s *Slave) Wait() {
    <-s.ch
}
