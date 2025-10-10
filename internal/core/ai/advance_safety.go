package ai

type AdvancedSafetyAnalyzer struct {
	contextAnalyzer *ContextAnalyzer
	intentDetector  *IntentDetector
	culturalNorms   *CulturalNorms
	moderationRules *ModerationRules
}

func NewAdvancedSafetyAnalyzer() *AdvancedSafetyAnalyzer {
	return &AdvancedSafetyAnalyzer{
		contextAnalyzer: NewContextAnalyzer(),
		intentDetector:  NewIntentDetector(),
		culturalNorms:   NewCulturalNorms(),
		moderationRules: NewModerationRules(),
	}
}

// AnalyzeAdvancedSafety performs deep context analysis for safety
func (a *AdvancedSafetyAnalyzer) AnalyzeAdvancedSafety(prompt string) *SafetyAnalysis {
	analysis := &SafetyAnalysis{
		IsSafe: true,
		Issues: []string{},
		Context: a.contextAnalyzer.AnalyzeContext(prompt),
	}
	
	// Check intent
	intent := a.intentDetector.DetectIntent(prompt)
	if !a.moderationRules.IsIntentAppropriate(intent) {
		analysis.IsSafe = false
		analysis.Issues = append(analysis.Issues, "Intent may not be appropriate")
	}
	
	// Check cultural appropriateness
	if !a.culturalNorms.IsCulturallyAppropriate(prompt, analysis.Context) {
		analysis.IsSafe = false
		analysis.Issues = append(analysis.Issues, "Content may not be culturally appropriate")
	}
	
	// Check for gymnastics and wide leg positions
	if a.containsGymnasticsOrWideLegs(prompt) {
		if !a.isGymnasticsAppropriate(prompt, analysis.Context) {
			analysis.IsSafe = false
			analysis.Issues = append(analysis.Issues, "Gymnastics position requires appropriate context")
		}
	}
	
	// Check clothing and activity combination
	if !a.isClothingActivityAppropriate(prompt) {
		analysis.IsSafe = false
		analysis.Issues = append(analysis.Issues, "Clothing may not be appropriate for activity")
	}
	
	return analysis
}

func (a *AdvancedSafetyAnalyzer) containsGymnasticsOrWideLegs(prompt string) bool {
	gymnasticsTerms := []string{
		"gymnastics", "split", "180", "flexibility", "contortion",
		"acrobatics", "tumbling", "floor exercise", "balance beam",
	}
	
	lowerPrompt := strings.ToLower(prompt)
	for _, term := range gymnasticsTerms {
		if strings.Contains(lowerPrompt, term) {
			return true
		}
	}
	return false
}

func (a *AdvancedSafetyAnalyzer) isGymnasticsAppropriate(prompt string, context ContextAnalysis) bool {
	// Gymnastics is appropriate in sports contexts
	sportsContexts := []string{
		"competition", "training", "gym", "sports", "athletic",
		"olympic", "professional", "practice", "coaching",
	}
	
	lowerPrompt := strings.ToLower(prompt)
	for _, ctx := range sportsContexts {
		if strings.Contains(lowerPrompt, ctx) {
			return true
		}
	}
	
	// Check for appropriate clothing in gymnastics context
	if strings.Contains(lowerPrompt, "gymnastics") {
		appropriateClothing := []string{
			"leotard", "unitard", "sports wear", "athletic wear",
			"gymnastics attire", "competition outfit",
		}
		
		for _, clothing := range appropriateClothing {
			if strings.Contains(lowerPrompt, clothing) {
				return true
			}
		}
	}
	
	return false
}

// MakeSafe modifies prompts to ensure safety while preserving intent
func (a *AdvancedSafetyAnalyzer) MakeSafe(prompt string, issues []string) (string, error) {
	safePrompt := prompt
	
	for _, issue := range issues {
		switch {
		case strings.Contains(issue, "Gymnastics position"):
			safePrompt = a.makeGymnasticsSafe(safePrompt)
		case strings.Contains(issue, "Clothing may not be appropriate"):
			safePrompt = a.makeClothingAppropriate(safePrompt)
		case strings.Contains(issue, "Intent may not be appropriate"):
			safePrompt, err := a.correctIntent(safePrompt)
			if err != nil {
				return "", err
			}
		case strings.Contains(issue, "Culturally appropriate"):
			safePrompt = a.makeCulturallyAppropriate(safePrompt)
		}
	}
	
	return safePrompt, nil
}

func (a *AdvancedSafetyAnalyzer) makeGymnasticsSafe(prompt string) string {
	// Add appropriate context and framing
	safeEnhancements := []string{
		"professional sports context",
		"athletic competition setting",
		"appropriate sports attire",
		"tasteful sports photography",
		"focus on athletic achievement",
	}
	
	// Handle mini-skirt in gymnastics context
	if strings.Contains(strings.ToLower(prompt), "mini skirt") {
		safeEnhancements = append(safeEnhancements,
			"sports shorts underneath",
			"athletic wear consideration",
			"modest sports representation",
		)
	}
	
	return prompt + ", " + strings.Join(safeEnhancements, ", ")
}
