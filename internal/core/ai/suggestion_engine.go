package ai

type SuggestionEngine struct {
	knowledgeBase *ProfessionalKnowledgeBase
	styleMatcher  *StyleMatchingEngine
}

func NewSuggestionEngine() *SuggestionEngine {
	return &SuggestionEngine{
		knowledgeBase: NewProfessionalKnowledgeBase(),
		styleMatcher:  NewStyleMatchingEngine(),
	}
}

// GenerateSuggestions provides intelligent improvements for prompts
func (s *SuggestionEngine) GenerateSuggestions(userPrompt string, currentOptions GenerationOptions) []Suggestion {
	var suggestions []Suggestion
	
	// Suggestion 1: Enhance with professional terminology
	if s.needsProfessionalTerms(userPrompt) {
		suggestions = append(suggestions, s.suggestProfessionalTerms(userPrompt))
	}
	
	// Suggestion 2: Improve composition
	if s.canImproveComposition(userPrompt) {
		suggestions = append(suggestions, s.suggestCompositionImprovements(userPrompt))
	}
	
	// Suggestion 3: Optimize lighting
	if s.canOptimizeLighting(userPrompt, currentOptions) {
		suggestions = append(suggestions, s.suggestLightingOptimizations(userPrompt, currentOptions))
	}
	
	// Suggestion 4: Material enhancements
	if s.canEnhanceMaterials(userPrompt) {
		suggestions = append(suggestions, s.suggestMaterialEnhancements(userPrompt))
	}
	
	return suggestions
}

func (s *SuggestionEngine) suggestProfessionalTerms(prompt string) Suggestion {
	professionalReplacements := map[string]string{
		"good lighting":    "professional studio lighting with softboxes",
		"nice background":  "minimalist studio backdrop with subtle gradient",
		"looking at camera": "engaging directly with the viewer with confident eye contact",
		"standing":         "dynamic pose with balanced weight distribution",
		"smiling":          "authentic expression with natural facial muscles",
	}
	
	var replacements []string
	for simple, professional := range professionalReplacements {
		if strings.Contains(strings.ToLower(prompt), simple) {
			replacements = append(replacements, fmt.Sprintf("'%s' → '%s'", simple, professional))
		}
	}
	
	return Suggestion{
		Type:        "ProfessionalTerminology",
		Description: "Enhance with professional photography terms",
		Examples:    replacements,
		Confidence:  0.85,
	}
}

func (s *SuggestionEngine) suggestCompositionImprovements(prompt string) Suggestion {
	compositionTips := []string{
		"Consider using rule of thirds for better framing",
		"Add negative space to create visual breathing room", 
		"Use leading lines to guide viewer's attention",
		"Consider different camera angles for dynamic composition",
		"Balance elements for visual harmony",
	}
	
	return Suggestion{
		Type:        "CompositionImprovement",
		Description: "Enhance image composition",
		Examples:    compositionTips,
		Confidence:  0.75,
	}
}
