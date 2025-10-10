package management

import (
	"runtime"
	"time"
)

type HealthMonitor struct {
	metricsCollector *MetricsCollector
	alertSystem      *AlertSystem
	thresholds       HealthThresholds
}

type HealthThresholds struct {
	MemoryUsage    float64
	CPULoad        float64
	ResponseTime   time.Duration
	ErrorRate      float64
	UptimeMinimum  time.Duration
}

func NewHealthMonitor() *HealthMonitor {
	return &HealthMonitor{
		metricsCollector: NewMetricsCollector(),
		alertSystem:      NewAlertSystem(),
		thresholds: HealthThresholds{
			MemoryUsage:   0.8,  // 80% max
			CPULoad:       0.75, // 75% max  
			ResponseTime:  2 * time.Second,
			ErrorRate:     0.05, // 5% max
			UptimeMinimum: 5 * time.Minute,
		},
	}
}

func (h *HealthMonitor) CheckAgentHealth(agentID string, agent Agent) HealthStatus {
	metrics := h.metricsCollector.CollectAgentMetrics(agentID)
	
	health := HealthStatus{
		Metrics:   metrics,
		LastCheck: time.Now(),
	}
	
	// Calculate health score based on multiple factors
	health.Score = h.calculateHealthScore(metrics)
	health.Status = h.determineHealthStatus(health.Score, metrics)
	
	// Trigger alerts if needed
	if health.Status != "healthy" {
		h.alertSystem.TriggerAlert(agentID, health.Status, metrics)
	}
	
	return health
}

func (h *HealthMonitor) calculateHealthScore(metrics map[string]float64) float64 {
	score := 1.0
	
	// Factor 1: Memory usage
	if memory, exists := metrics["memory_usage"]; exists {
		if memory > h.thresholds.MemoryUsage {
			penalty := (memory - h.thresholds.MemoryUsage) / (1 - h.thresholds.MemoryUsage)
			score -= penalty * 0.3
		}
	}
	
	// Factor 2: CPU load
	if cpu, exists := metrics["cpu_load"]; exists {
		if cpu > h.thresholds.CPULoad {
			penalty := (cpu - h.thresholds.CPULoad) / (1 - h.thresholds.CPULoad)
			score -= penalty * 0.3
		}
	}
	
	// Factor 3: Error rate
	if errorRate, exists := metrics["error_rate"]; exists {
		if errorRate > h.thresholds.ErrorRate {
			penalty := (errorRate - h.thresholds.ErrorRate) / h.thresholds.ErrorRate
			score -= penalty * 0.2
		}
	}
	
	// Factor 4: Response time
	if responseTime, exists := metrics["response_time"]; exists {
		thresholdMs := h.thresholds.ResponseTime.Seconds() * 1000
		if responseTime > thresholdMs {
			penalty := (responseTime - thresholdMs) / thresholdMs
			score -= penalty * 0.2
		}
	}
	
	return max(0.0, min(1.0, score))
}

func (h *HealthMonitor) determineHealthStatus(score float64, metrics map[string]float64) string {
	switch {
	case score >= 0.8:
		return "healthy"
	case score >= 0.6:
		return "degraded"
	default:
		return "unhealthy"
	}
}

// System-wide health check
func (h *HealthMonitor) CheckSystemHealth() *SystemHealth {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	metrics := map[string]float64{
		"memory_alloc":      float64(m.Alloc) / 1024 / 1024, // MB
		"memory_total":      float64(m.Sys) / 1024 / 1024,   // MB
		"goroutines":        float64(runtime.NumGoroutine()),
		"cpu_cores":         float64(runtime.NumCPU()),
		"gc_pauses":         float64(m.PauseTotalNs) / 1e9, // seconds
	}
	
	return &SystemHealth{
		Timestamp: time.Now(),
		Metrics:   metrics,
		Status:    h.assessSystemStatus(metrics),
		Score:     h.calculateSystemHealthScore(metrics),
	}
}
