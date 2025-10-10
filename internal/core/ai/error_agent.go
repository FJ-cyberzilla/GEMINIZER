package ai

import (
	"fmt"
	"regexp"
	"strings"
)

type ErrorHandlingAgent struct {
	patternDB    *ErrorPatternDatabase
	suggestionEngine *SuggestionEngine
}

func NewErrorHandlingAgent() *ErrorHandlingAgent {
	return &ErrorHandlingAgent{
		patternDB:    NewErrorPatternDatabase(),
		suggestionEngine: NewSuggestionEngine(),
	}
}

// AnalyzeError uses ML patterns to understand and categorize errors
func (e *ErrorHandlingAgent) AnalyzeError(userPrompt string, generatedPrompt string, resultImage []byte) *ErrorAnalysis {
	analysis := &ErrorAnalysis{
		UserPrompt:      userPrompt,
		GeneratedPrompt: generatedPrompt,
		Confidence:      0.8,
	}
	
	// Detect common issues using pattern matching
	analysis.Issues = e.detectIssues(userPrompt, generatedPrompt)
	analysis.Suggestions = e.generateSuggestions(analysis.Issues)
	analysis.Severity = e.calculateSeverity(analysis.Issues)
	
	return analysis
}

func (e *ErrorHandlingAgent) detectIssues(userPrompt, generatedPrompt string) []Issue {
	var issues []Issue
	
	// ML Pattern 1: Vague descriptions
	if e.isTooVague(userPrompt) {
		issues = append(issues, Issue{
			Type:        "VagueDescription",
			Description: "Prompt lacks specific details for professional rendering",
			Confidence:  0.9,
		})
	}
	
	// ML Pattern 2: Conflicting elements
	if e.hasConflictingElements(generatedPrompt) {
		issues = append(issues, Issue{
			Type:        "ConflictingElements", 
			Description: "Prompt contains elements that may conflict in rendering",
			Confidence:  0.7,
		})
	}
	
	// ML Pattern 3: Missing professional elements
	if e.missingProfessionalElements(userPrompt, generatedPrompt) {
		issues = append(issues, Issue{
			Type:        "MissingProfessionalElements",
			Description: "Prompt could benefit from professional photography terms",
			Confidence:  0.8,
		})
	}
	
	// ML Pattern 4: Physical impossibility
	if e.hasPhysicalImpossibility(generatedPrompt) {
		issues = append(issues, Issue{
			Type:        "PhysicalImpossibility",
			Description: "Prompt describes physically impossible scenarios",
			Confidence:  0.95,
		})
	}
	
	return issues
}

func (e *ErrorHandlingAgent) isTooVague(prompt string) bool {
	vagueIndicators := []string{
		"nice", "good", "beautiful", "awesome", "cool",
		"amazing", "great", "pretty", "looking good",
	}
	
	wordCount := len(strings.Fields(prompt))
	hasVagueTerms := false
	
	for _, term := range vagueIndicators {
		if strings.Contains(strings.ToLower(prompt), term) {
			hasVagueTerms = true
			break
		}
	}
	
	return wordCount < 10 || hasVagueTerms
}

func (e *ErrorHandlingAgent) hasConflictingElements(prompt string) bool {
	conflicts := []struct {
		element1 string
		element2 string
	}{
		{"indoor", "sunlight"},
		{"night", "bright sunshine"},
		{"underwater", "fire"},
		{"winter", "swimming"},
	}
	
	for _, conflict := range conflicts {
		if strings.Contains(prompt, conflict.element1) && strings.Contains(prompt, conflict.element2) {
			return true
		}
	}
	
	return false
}
