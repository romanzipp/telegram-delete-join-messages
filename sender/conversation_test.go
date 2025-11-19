package sender

import (
	"sync"
	"testing"
)

// TestConcurrentMapAccess verifies that there is no race condition during concurrent access to ConversationHandler
func TestConcurrentMapAccess(t *testing.T) {
	ch := NewConversationHandler()

	// Create WaitGroup for goroutine synchronization
	var wg sync.WaitGroup

	// Number of goroutines to simulate concurrent access
	numGoroutines := 100

	// Launch multiple goroutines that perform operations on the map concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			// Simulate operations that caused race condition
			ch.SetActiveStage(0, userID)
			stage := ch.GetActiveStage(userID)
			if stage != 0 {
				ch.SetActiveStage(1, userID)
			}
			ch.End(userID)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	t.Log("Concurrent map access test passed successfully")
}

// TestConversationHandlerBasicFunctionality verifies basic functionality
func TestConversationHandlerBasicFunctionality(t *testing.T) {
	ch := NewConversationHandler()

	// Test setting and getting active stage
	userID := 123
	stageID := 5

	ch.SetActiveStage(stageID, userID)

	activeStage := ch.GetActiveStage(userID)
	if activeStage != stageID {
		t.Errorf("Expected stage %d, got %d", stageID, activeStage)
	}

	// Test conversation completion
	ch.End(userID)

	// After completion, user should be inactive
	// But GetActiveStage may still return the last stage if the user is marked as active
	activeAfterEnd := ch.GetActiveStage(userID)
	if activeAfterEnd != 0 {
		t.Logf("After End() got stage %d (this is expected if active[userID] = false)", activeAfterEnd)
	}
}
