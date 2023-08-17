/*
功能：服务
说明：
*/
package xservice

import (
	"fmt"
	"sync"
)

type Service struct {
	status   Status
	rw_mutex sync.RWMutex
}

func (s *Service) Service_status() Status {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status
}

// 更新状态
func (s *Service) Update_status(st Status) {
	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()
	s.status = st
}

func (s *Service) Is_running() bool {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status == STATUS_RUNNING
}

func (s *Service) Is_dead() bool {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status == STATUS_DEAD
}

// 锁
func (s *Service) With_lock_read(handler func()) {
	if handler == nil {
		return
	}

	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	handler()
}

func (s *Service) With_lock_read_error(handler func() error) error {
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return handler()
}

func (s *Service) With_lock(handler func()) {
	if handler == nil {
		return
	}

	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()
	handler()
}

func (s *Service) With_lock_error(handler func() error) error {
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()
	return handler()
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
