package ai

type LoopDetector struct {
	recentPrompts    []string
	maxHistory       int
	similarityThreshold float64
	recoveryEngine   *RecoveryEngine
}

func NewLoopDetector() *LoopDetector {
	return &LoopDetector{
		recentPrompts:    make([]string, 0),
		maxHistory:       10,
		similarityThreshold: 0.8,
		recoveryEngine:   NewRecoveryEngine(),
	}
}

// IsInLoop detects if generation is stuck in repetitive loop
func (l *LoopDetector) IsInLoop(currentPrompt string) bool {
	// Check if current prompt is too similar to recent ones
	for _, pastPrompt := range l.recentPrompts {
		similarity := l.calculateSimilarity(currentPrompt, pastPrompt)
		if similarity > l.similarityThreshold {
			return true
		}
	}
	
	// Add to history
	l.recentPrompts = append(l.recentPrompts, currentPrompt)
	if len(l.recentPrompts) > l.maxHistory {
		l.recentPrompts = l.recentPrompts[1:]
	}
	
	return false
}

// RecoverFromLoop provides alternative prompts when loop detected
func (l *LoopDetector) RecoverFromLoop(stuckPrompt string, context string) []string {
	return l.recoveryEngine.GenerateAlternatives(stuckPrompt, context)
}

func (l *LoopDetector) calculateSimilarity(prompt1, prompt2 string) float64 {
	// Simple word-based similarity (can be enhanced with ML)
	words1 := strings.Fields(strings.ToLower(prompt1))
	words2 := strings.Fields(strings.ToLower(prompt2))
	
	commonWords := 0
	wordSet := make(map[string]bool)
	
	for _, word := range words1 {
		wordSet[word] = true
	}
	
	for _, word := range words2 {
		if wordSet[word] {
			commonWords++
		}
	}
	
	totalWords := len(words1) + len(words2)
	if totalWords == 0 {
		return 0.0
	}
	
	return float64(2*commonWords) / float64(totalWords)
}
