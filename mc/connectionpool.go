package mc

import (
	"sync"
)

type ConnectionPool struct {
	pool []Connection
	mu   sync.Mutex
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		pool: make([]Connection, 0),
		mu:   sync.Mutex{},
	}
}

func (cp *ConnectionPool) GetConnection() *Connection {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if len(cp.pool) > 0 {
		conn := cp.pool[0]
		cp.pool = cp.pool[1:]
		return conn
	}

	newConn := NewConnection()
	cp.pool = append(cp.pool, newConn)
	return newConn
}

func (cp *ConnectionPool) ReturnConnection(conn *Connection) {
	cp.mu.Lock()
	cp.pool = append(cp.pool, conn)
	cp.mu.Unlock()
}
