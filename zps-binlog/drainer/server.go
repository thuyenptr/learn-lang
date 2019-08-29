package drainer

import "context"

type Server struct {
	collector *Collector
	ctx context.Context

	latestTS int64
}

func NewServer() *Server {

	collector := NewCollector()
	return &Server{
		collector: collector,

	}
}

func (s *Server) Start() {
	s.collector.Start(s.ctx)
}