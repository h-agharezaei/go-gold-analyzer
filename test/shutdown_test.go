package test

import (
	"os"
	"sync"
	"testing"
	"time"

	"gold-analyzer/shutdown"
)

func TestShutdownManagerCreation(t *testing.T) {
	mgr := shutdown.NewManager()

	if mgr == nil {
		t.Error("Expected shutdown manager, got nil")
	}
}

func TestShutdownManagerIsRunning(t *testing.T) {
	mgr := shutdown.NewManager()

	// Should be running initially
	if !mgr.IsRunning() {
		t.Error("Expected manager to be running")
	}

	// After stop, should not be running
	mgr.Stop()
	if mgr.IsRunning() {
		t.Error("Expected manager to be stopped")
	}
}

func TestShutdownManagerRegisterHook(t *testing.T) {
	mgr := shutdown.NewManager()

	hookCalled := false
	mgr.RegisterHook(func() error {
		hookCalled = true
		return nil
	})

	// Shutdown should call the hook
	if err := mgr.Shutdown(1 * time.Second); err != nil {
		t.Errorf("Shutdown returned error: %v", err)
	}

	if !hookCalled {
		t.Error("Hook was not called during shutdown")
	}
}

func TestShutdownManagerMultipleHooks(t *testing.T) {
	mgr := shutdown.NewManager()

	callCount := 0
	mu := &sync.Mutex{}

	mgr.RegisterHook(func() error {
		mu.Lock()
		callCount++
		mu.Unlock()
		return nil
	})

	mgr.RegisterHook(func() error {
		mu.Lock()
		callCount++
		mu.Unlock()
		return nil
	})

	mgr.RegisterHook(func() error {
		mu.Lock()
		callCount++
		mu.Unlock()
		return nil
	})

	if err := mgr.Shutdown(1 * time.Second); err != nil {
		t.Errorf("Shutdown returned error: %v", err)
	}

	if callCount != 3 {
		t.Errorf("Expected 3 hooks called, got %d", callCount)
	}
}

func TestShutdownManagerStart(t *testing.T) {
	mgr := shutdown.NewManager()
	mgr.Start()

	// Should still be running after start
	if !mgr.IsRunning() {
		t.Error("Expected manager to be running after Start()")
	}
}

func TestShutdownManagerTimeout(t *testing.T) {
	mgr := shutdown.NewManager()

	// Register a hook that takes longer than timeout
	mgr.RegisterHook(func() error {
		time.Sleep(2 * time.Second)
		return nil
	})

	// Use short timeout
	startTime := time.Now()
	err := mgr.Shutdown(100 * time.Millisecond)
	duration := time.Since(startTime)

	// Should return error due to timeout
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}

	// Duration should be roughly the timeout (±500ms tolerance)
	if duration < 50*time.Millisecond || duration > 600*time.Millisecond {
		t.Errorf("Shutdown took unexpected duration: %v", duration)
	}
}

func TestShutdownManagerSignalChannel(t *testing.T) {
	mgr := shutdown.NewManager()
	mgr.Start()

	// Send a signal after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		mgr.GetShutdownChan() // This would block in real scenario
	}()

	// Give time for goroutine to run
	time.Sleep(200 * time.Millisecond)

	if !mgr.IsRunning() {
		t.Error("Manager should still be running before Stop()")
	}
}

func TestShutdownManagerGetLastError(t *testing.T) {
	mgr := shutdown.NewManager()

	// Initially should have no error
	if mgr.GetLastError() != nil {
		t.Error("Expected no initial error")
	}

	// Set an error
	testErr := os.ErrNotExist
	mgr.SetLastError(testErr)

	// Should return the set error
	if mgr.GetLastError() != testErr {
		t.Error("Expected to retrieve the set error")
	}
}

func TestShutdownManagerDoubleShutdown(t *testing.T) {
	mgr := shutdown.NewManager()

	callCount := 0
	mgr.RegisterHook(func() error {
		callCount++
		return nil
	})

	// First shutdown
	if err := mgr.Shutdown(1 * time.Second); err != nil {
		t.Errorf("First shutdown returned error: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected hook called once, got %d times", callCount)
	}

	// Second shutdown should not call hooks again
	if err := mgr.Shutdown(1 * time.Second); err != nil {
		t.Errorf("Second shutdown returned error: %v", err)
	}

	if callCount != 1 {
		t.Errorf("Expected hook still called once, got %d times", callCount)
	}
}

func TestShutdownManagerSignalHandling(t *testing.T) {
	// This test would require signal mocking
	// For now, we just verify the manager can be created and used

	mgr := shutdown.NewManager()
	mgr.Start()

	// Verify it's listening for signals (implementation detail)
	if !mgr.IsRunning() {
		t.Error("Manager should be running after Start()")
	}

	// Cleanup
	mgr.Stop()
	mgr.Shutdown(1 * time.Second)
}

func BenchmarkShutdownManager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mgr := shutdown.NewManager()
		mgr.RegisterHook(func() error {
			return nil
		})
		mgr.Shutdown(1 * time.Second)
	}
}

func BenchmarkShutdownManagerWithMultipleHooks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mgr := shutdown.NewManager()

		for j := 0; j < 10; j++ {
			mgr.RegisterHook(func() error {
				return nil
			})
		}

		mgr.Shutdown(1 * time.Second)
	}
}
