/**
 * Author: Mitch Allen
 * File: alpha_test.go
 */

package alpha

import (
	"math"
	"testing"
)

func TestCountCount(t *testing.T) {

	limit := 100
	expectedThreshold := int(math.Round(float64(limit) * 0.3))

	if got := CoinCount(limit); got[true] < expectedThreshold {
		t.Errorf("CoinCount(%d) = %v (threshold: %d)", limit, got, expectedThreshold)
	}
}
