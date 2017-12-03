package main

import "time"

type ServerInfo struct {
	startedAt time.Time
	requests  int
}

func (s *ServerInfo) getUptime() string {
	duration := time.Since(s.startedAt)
	return duration.String()
}

func (s *ServerInfo) getStartedAt() string {
	return s.startedAt.String()
}

func (s *ServerInfo) init() {
	s.startedAt = time.Now()
	s.requests = 0
}

func (s *ServerInfo) setCounter() {
	s.requests++
}

func (s *ServerInfo) getCounter() int {
	return s.requests
}
