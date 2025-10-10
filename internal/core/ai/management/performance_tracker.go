package management

type PerformanceTracker struct {
	metricsDB    *PerformanceDatabase
	optimizer    *PerformanceOptimizer
	trendAnalyzer *TrendAnalyzer
}

func NewPerformanceTracker() *PerformanceTracker {
	return &PerformanceTracker{
		metricsDB:    NewPerformanceDatabase(),
		optimizer:    NewPerformanceOptimizer(),
		trendAnalyzer: NewTrendAnalyzer(),
	}
}

func (p *PerformanceTracker) EvaluatePerformance(agentID string) float64 {
	metrics := p.metricsDB.GetAgentMetrics(agentID)
	trends := p.trendAnalyzer.AnalyzeTrends(agentID, metrics)
	
	// Calculate performance score (0.0 - 1.0)
	performance := p.calculatePerformanceScore(metrics, trends)
	
	// Apply optimizations if performance is suboptimal
	if performance < 0.7 {
		p.optimizer.OptimizeAgent(agentID, metrics, trends)
	}
	
	return performance
}

func (p *PerformanceTracker) calculatePerformanceScore(metrics AgentMetrics, trends TrendAnalysis) float64 {
	score := 0.0
	
	// Success rate (40% weight)
	successRate := metrics.SuccessfulGenerations / metrics.TotalGenerations
	score += successRate * 0.4
	
	// Response time (25% weight)
	responseScore := 1.0 - min(metrics.AverageResponseTime/5000.0, 1.0) // Normalize to 5 seconds
	score += responseScore * 0.25
	
	// Quality score (20% weight)
	score += metrics.AverageQuality * 0.2
	
	// Trend analysis (15% weight)
	trendScore := p.calculateTrendScore(trends)
	score += trendScore * 0.15
	
	return score
}

// Record successful generation for learning
func (p *PerformanceTracker) RecordSuccess(agentID string, generationData GenerationData) {
	p.metricsDB.RecordSuccess(agentID, generationData)
	
	// Learn from successful generations
	p.optimizer.LearnFromSuccess(agentID, generationData)
}

// Record failure for improvement
func (p *PerformanceTracker) RecordFailure(agentID string, errorData ErrorData) {
	p.metricsDB.RecordFailure(agentID, errorData)
	
	// Learn from failures
	p.optimizer.LearnFromFailure(agentID, errorData)
}
