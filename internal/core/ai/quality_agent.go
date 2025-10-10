package ai

type QualityAssessmentAgent struct {
	qualityMetrics *QualityMetrics
	improvementEngine *ImprovementEngine
}

func NewQualityAssessmentAgent() *QualityAssessmentAgent {
	return &QualityAssessmentAgent{
		qualityMetrics: NewQualityMetrics(),
		improvementEngine: NewImprovementEngine(),
	}
}

// AssessPromptQuality evaluates prompt quality and suggests improvements
func (q *QualityAssessmentAgent) AssessPromptQuality(prompt string, intent string) *QualityAssessment {
	assessment := &QualityAssessment{
		Prompt: prompt,
		Intent: intent,
	}
	
	assessment.Score = q.calculateQualityScore(prompt, intent)
	assessment.Strengths = q.identifyStrengths(prompt, intent)
	assessment.Weaknesses = q.identifyWeaknesses(prompt, intent)
	assessment.Improvements = q.improvementEngine.SuggestImprovements(prompt, intent, assessment.Weaknesses)
	assessment.ProfessionalLevel = q.determineProfessionalLevel(assessment.Score)
	
	return assessment
}

func (q *QualityAssessmentAgent) calculateQualityScore(prompt string, intent string) float64 {
	score := 0.0
	
	// Factor 1: Specificity (0-30 points)
	specificityScore := q.assessSpecificity(prompt)
	score += specificityScore * 0.3
	
	// Factor 2: Professional terminology (0-25 points)  
	professionalScore := q.assessProfessionalTerminology(prompt)
	score += professionalScore * 0.25
	
	// Factor 3: Technical completeness (0-20 points)
	technicalScore := q.assessTechnicalCompleteness(prompt, intent)
	score += technicalScore * 0.2
	
	// Factor 4: Clarity and coherence (0-15 points)
	clarityScore := q.assessClarity(prompt)
	score += clarityScore * 0.15
	
	// Factor 5: Creativity and uniqueness (0-10 points)
	creativityScore := q.assessCreativity(prompt)
	score += creativityScore * 0.1
	
	return score / 100.0
}

func (q *QualityAssessmentAgent) assessSpecificity(prompt string) float64 {
	wordCount := len(strings.Fields(prompt))
	specificTerms := []string{
		"specific", "detailed", "precise", "exact", "particular",
		"8K", "4K", "professional", "studio", "lighting", "angle",
	}
	
	specificTermCount := 0
	for _, term := range specificTerms {
		if strings.Contains(strings.ToLower(prompt), term) {
			specificTermCount++
		}
	}
	
	// Score based on word count and specific terms
	baseScore := float64(wordCount) * 0.5
	termBonus := float64(specificTermCount) * 10
	
	return min(baseScore+termBonus, 30)
}

func (q *QualityAssessmentAgent) identifyStrengths(prompt string, intent string) []string {
	var strengths []string
	
	if q.hasGoodLightingDescription(prompt) {
		strengths = append(strengths, "Good lighting description")
	}
	
	if q.hasClearSubject(prompt) {
		strengths = append(strengths, "Clear subject focus")
	}
	
	if q.hasAppropriateDetail(prompt) {
		strengths = append(strengths, "Appropriate level of detail")
	}
	
	if q.hasConsistentStyle(prompt, intent) {
		strengths = append(strengths, "Consistent artistic style")
	}
	
	return strengths
}
