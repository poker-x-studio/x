/*
功能：服务
说明：
*/
package xservice

import (
	"fmt"
	"sync"
)

// Service 服务
type Service struct {
	status   Status
	rw_mutex sync.RWMutex
}

// Service_status 状态
func (s *Service) Service_status() Status {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status
}

// Update_status 更新状态
func (s *Service) Update_status(st Status) {
	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()
	s.status = st
}

// Is_running 是否在运行
func (s *Service) Is_running() bool {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status == STATUS_RUNNING
}

// Is_dead 是否关闭
func (s *Service) Is_dead() bool {
	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return s.status == STATUS_DEAD
}

// With_lock_read 锁
func (s *Service) With_lock_read(handler func()) {
	if handler == nil {
		return
	}

	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	handler()
}

// With_lock_read_error 锁
func (s *Service) With_lock_read_error(handler func() error) error {
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	s.rw_mutex.RLock()
	defer s.rw_mutex.RUnlock()
	return handler()
}

// With_lock 锁
func (s *Service) With_lock(handler func()) {
	if handler == nil {
		return
	}

	s.rw_mutex.Lock()
	defer s.rw_mutex.Unlock()
	handler()
}

// With_lock_error 锁
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
