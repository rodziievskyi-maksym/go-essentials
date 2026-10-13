package main

import (
	"context"
	"sync"
	"time"
)

// InventoryTask the unit that represents one SKU of product
type InventoryTask struct {
	//ProductID - id of product on shopify
	ProductID string

	//VariantID - specific variant of product(size,  color, package)
	VariantID string

	AvailableQuantity int

	//LocationID warehouse id
	LocationID string

	// Individual context of task in case request comes from HTTP endpoint with Timeout
	Ctx context.Context

	// Time when got in queue for latency calculation
	/*
		CreatedAt time.Time: У production-метриках критично бачити різницю time.Since(task.CreatedAt) — це метрика Queue Wait Time.
		Якщо вона росте, значить обраного тарифу в 20 RPS уже недостатньо під обсяг складу.
	*/
	CreatedAt time.Time
}

type TokenBucketLimiter struct {
	mu sync.Mutex

	rate  float64 // Allowed RPS to external API (Shopify) -> RECOVERY RATE HOW QUICKLY BUCKET IS RECOVERING -> 20 RPS
	burst float64 // Maximum requests allowed from Shopify

	tokens     float64   // how many tokens in bucket now,  current dynamic state of limiter
	lastUpdate time.Time // tokens amount last update time
}

func NewTokenBucketLimiter(rps float64, burst float64) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		rate:       rps,
		burst:      burst,
		tokens:     burst, //start with full bucket for immediate readiness
		lastUpdate: time.Now(),
	}
}

func (tb *TokenBucketLimiter) SetRate(newRate float64, newBurst int) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// First add tokens from old tariff to current state
	tb.replenish(time.Now())

	//set new rules
	tb.rate = newRate
	tb.burst = float64(newBurst)

	if tb.tokens > tb.burst {
		tb.tokens = tb.burst
	}
}

func (tb *TokenBucketLimiter) replenish(now time.Time) { //04 00 -> 03 00
	// Now = 3:01:20 = -234234234 * 40 = -29834723984729384234
	//lastUpdate -> 3:58:40
	if now.Before(tb.lastUpdate) {
		tb.lastUpdate = now
		return
	}

	elapsed := now.Sub(tb.lastUpdate).Seconds()
	tb.lastUpdate = now

	// adding tokens which generated from lastUpdate
	tb.tokens = elapsed * tb.rate

	//The bucket can't be more than burst
	if tb.tokens > tb.burst {
		tb.tokens = tb.burst
	}
}

func (tb *TokenBucketLimiter) Wait(ctx context.Context) error {
	withDuration, err := tb.reserve()
	if err != nil {
		return err
	}

	if withDuration <= 0 {
		return nil
	}

	timer := time.NewTimer(withDuration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		// context is canceled or timeout or service stopped
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (tb *TokenBucketLimiter) reserve() (time.Duration, error) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	tb.replenish(now)

	if tb.tokens > 1.0 {
		tb.tokens -= 1.0
		return 0, nil
	}

	missing := 1.0 - tb.tokens                                    // 0.7
	waitSec := missing / tb.rate                                  // 20.0 -> 0.7 / 20.0 = 0,035
	waitDuration := time.Duration(waitSec * float64(time.Second)) //0,035 * 23420342934.0

	tb.tokens -= 1.0 // 70

	return waitDuration, nil
}
