package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Manager handles graceful shutdown of the application
type Manager struct {
	stopChan      chan os.Signal
	shutdownChan  chan bool
	mu            sync.Mutex
	isRunning     bool
	shutdownTime  time.Time
	lastError     error
	shutdownHooks []func() error
	shutdownDone  bool
}

// NewManager creates a new shutdown manager
func NewManager() *Manager {
	return &Manager{
		stopChan:      make(chan os.Signal, 1),
		shutdownChan:  make(chan bool, 1),
		isRunning:     true,
		shutdownHooks: make([]func() error, 0),
	}
}

// RegisterHook registers a function to be called during shutdown
func (m *Manager) RegisterHook(hook func() error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.shutdownHooks = append(m.shutdownHooks, hook)
}

// Start initializes signal handlers
func (m *Manager) Start() {
	m.mu.Lock()
	m.shutdownTime = time.Now()
	m.mu.Unlock()

	// Register for SIGINT (Ctrl+C) and SIGTERM
	signal.Notify(m.stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
}

// Stop triggers graceful shutdown
func (m *Manager) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.shutdownDone {
		return
	}

	m.isRunning = false
	m.shutdownDone = true
}

// IsRunning returns whether the application is still running
func (m *Manager) IsRunning() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.isRunning
}

// WaitForShutdown waits for shutdown signal
func (m *Manager) WaitForShutdown() os.Signal {
	sig := <-m.stopChan
	return sig
}

// Shutdown performs graceful shutdown with cleanup
func (m *Manager) Shutdown(timeout time.Duration) error {
	m.mu.Lock()
	if m.shutdownDone {
		m.mu.Unlock()
		return nil
	}

	m.isRunning = false
	m.shutdownDone = true
	m.mu.Unlock()

	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}))
	fmt.Println("🔄 شروع خاتمهٔ نرم برنامه...")
	fmt.Println(string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}))

	// Run shutdown hooks with timeout
	shutdownCtx := time.Now()
	hooksCompleted := make(chan error, len(m.shutdownHooks))

	for _, hook := range m.shutdownHooks {
		go func(h func() error) {
			if err := h(); err != nil {
				hooksCompleted <- err
			} else {
				hooksCompleted <- nil
			}
		}(hook)
	}

	// Wait for all hooks to complete or timeout
	shutdownTimer := time.NewTimer(timeout)
	defer shutdownTimer.Stop()

	for i := 0; i < len(m.shutdownHooks); i++ {
		select {
		case err := <-hooksCompleted:
			if err != nil {
				fmt.Printf("⚠️  خطا در shutdown hook: %v\n", err)
				m.lastError = err
			}
		case <-shutdownTimer.C:
			fmt.Printf("⏱️  مهلت زمانی خاتمه تمام شد\n")
			return fmt.Errorf("shutdown timeout exceeded")
		}
	}

	// Calculate shutdown duration
	shutdownDuration := time.Since(shutdownCtx)

	fmt.Println("\n📊 آمار نهایی:")
	fmt.Printf("   ✓ تمام hooks تکمیل شدند\n")
	fmt.Printf("   ✓ مدت زمان Shutdown: %v\n", shutdownDuration)
	fmt.Printf("   ✓ زمان خاتمه: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}))
	fmt.Println("✅ برنامه با موفقیت متوقف شد")
	fmt.Println(string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}))

	return m.lastError
}

// GetShutdownChan returns the shutdown channel for select statements
func (m *Manager) GetShutdownChan() <-chan bool {
	return m.shutdownChan
}

// SignalShutdownComplete signals that shutdown is complete
func (m *Manager) SignalShutdownComplete() {
	m.shutdownChan <- true
}

// GetLastError returns the last error that occurred
func (m *Manager) GetLastError() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.lastError
}

// SetLastError sets the last error
func (m *Manager) SetLastError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lastError = err
}
