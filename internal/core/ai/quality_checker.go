package ai

type ProfessionalQualityChecker struct {
	professionalTerms []string
	requiredElements  []string
	commonMistakes   []string
}

func NewProfessionalQualityChecker() *ProfessionalQualityChecker {
	return &ProfessionalQualityChecker{
		professionalTerms: []string{
			"photography", "studio", "lighting", "aperture", "shutter", 
			"composition", "framing", "depth of field", "bokeh", "chiaroscuro",
			"softbox", "reflector", "key light", "fill light", "rim light",
		},
		requiredElements: []string{
			"subject", "lighting", "composition", "style", "quality",
		},
		commonMistakes: []string{
			"plastic skin", "deformed hands", "blurry artifacts",
			"oversaturated", "unnatural lighting", "poor composition",
		},
	}
}

func (q *ProfessionalQualityChecker) AnalyzePromptQuality(prompt string) *QualityReport {
	report := &QualityReport{
		Prompt: prompt,
		Score:  q.calculateQualityScore(prompt),
	}
	
	report.HasProfessionalTerms = q.hasProfessionalTerminology(prompt)
	report.HasLightingSpec = q.hasLightingSpecification(prompt)
	report.HasCameraDetails = q.hasCameraDetails(prompt)
	report.HasRealismRequirements = q.hasRealismRequirements(prompt)
	report.Suggestions = q.generateQualitySuggestions(report)
	
	return report
}

func (q *ProfessionalQualityChecker) calculateQualityScore(prompt string) float64 {
	score := 0.0
	
	// Length and specificity (30%)
	wordCount := len(strings.Fields(prompt))
	score += min(float64(wordCount)/50.0, 0.3) // Max 30% for length
	
	// Professional terminology (25%)
	if q.hasProfessionalTerminology(prompt) {
		score += 0.25
	}
	
	// Technical completeness (20%)
	if q.hasTechnicalCompleteness(prompt) {
		score += 0.2
	}
	
	// Realism requirements (15%)
	if q.hasRealismRequirements(prompt) {
		score += 0.15
	}
	
	// Clarity and structure (10%)
	if q.hasGoodStructure(prompt) {
		score += 0.1
	}
	
	return score
}

func (q *ProfessionalQualityChecker) hasProfessionalTerminology(prompt string) bool {
	lowerPrompt := strings.ToLower(prompt)
	
	for _, term := range q.professionalTerms {
		if strings.Contains(lowerPrompt, strings.ToLower(term)) {
			return true
		}
	}
	return false
}

func (q *ProfessionalQualityChecker) generateQualitySuggestions(report *QualityReport) []string {
	var suggestions []string
	
	if !report.HasProfessionalTerms {
		suggestions = append(suggestions, "Add professional photography terminology (aperture, lighting, composition)")
	}
	
	if !report.HasLightingSpec {
		suggestions = append(suggestions, "Specify lighting setup (studio, natural, dramatic)")
	}
	
	if !report.HasCameraDetails {
		suggestions = append(suggestions, "Include camera details (lens, angle, framing)")
	}
	
	if !report.HasRealismRequirements {
		suggestions = append(suggestions, "Add realism requirements (skin texture, hair details, fabric physics)")
	}
	
	if report.Score < 0.7 {
		suggestions = append(suggestions, "Increase specificity and technical details for professional results")
	}
	
	return suggestions
}
