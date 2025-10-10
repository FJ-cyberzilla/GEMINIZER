package services

// EnhancedImageGenerator with AI agents
type EnhancedImageGenerator struct {
	*ImageGenerator
	errorAgent    *ai.ErrorHandlingAgent
	nluEngine     *ai.NLUEngine
	qualityAgent  *ai.QualityAssessmentAgent
	suggestionAgent *ai.SuggestionEngine
}

func NewEnhancedImageGenerator(repo HistoryRepository, logger Logger) *EnhancedImageGenerator {
	baseGenerator := NewImageGenerator(repo, logger)
	
	return &EnhancedImageGenerator{
		ImageGenerator: baseGenerator,
		errorAgent:     ai.NewErrorHandlingAgent(),
		nluEngine:      ai.NewNLUEngine(),
		qualityAgent:   ai.NewQualityAssessmentAgent(),
		suggestionAgent: ai.NewSuggestionEngine(),
	}
}

// GenerateWithAnalysis provides enhanced generation with AI analysis
func (e *EnhancedImageGenerator) GenerateWithAnalysis(ctx context.Context, req domain.GenerationRequest) (*domain.EnhancedGenerationResponse, error) {
	// Step 1: Natural Language Understanding
	promptUnderstanding := e.nluEngine.UnderstandPrompt(req.UserPrompt)
	
	// Step 2: Quality Assessment
	qualityAssessment := e.qualityAgent.AssessPromptQuality(req.UserPrompt, promptUnderstanding.Intent)
	
	// Step 3: Generate image (original logic)
	response, err := e.ImageGenerator.Generate(ctx, req)
	if err != nil {
		// Use error agent to provide helpful error messages
		errorAnalysis := e.errorAgent.AnalyzeError(req.UserPrompt, "", nil)
		return nil, domain.NewAppError(err, errorAnalysis.Suggestions[0].Description, "ENHANCED_ERROR")
	}
	
	// Step 4: Generate intelligent suggestions
	suggestions := e.suggestionAgent.GenerateSuggestions(req.UserPrompt, req.Options)
	
	// Step 5: Create enhanced response
	enhancedResponse := &domain.EnhancedGenerationResponse{
		GenerationResponse: *response,
		Analysis: &domain.GenerationAnalysis{
			Intent:            promptUnderstanding.Intent,
			Entities:          promptUnderstanding.Entities,
			QualityScore:      qualityAssessment.Score,
			Strengths:         qualityAssessment.Strengths,
			ImprovementAreas:  qualityAssessment.Weaknesses,
			Suggestions:       suggestions,
			ProfessionalLevel: qualityAssessment.ProfessionalLevel,
		},
	}
	
	return enhancedResponse, nil
}
