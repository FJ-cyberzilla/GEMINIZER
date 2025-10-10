package services

type EnterpriseGenerationService struct {
	finalGeneration *FinalGenerationService
	expressionEngine *ai.CharacterExpressionEngine
	artStyleEngine  *ai.ArtStyleEngine
	qualityControl  *ai.EnterpriseQualityControl
}

func NewEnterpriseGenerationService(repo HistoryRepository, logger Logger) *EnterpriseGenerationService {
	return &EnterpriseGenerationService{
		finalGeneration:  NewFinalGenerationService(repo, logger),
		expressionEngine: ai.NewCharacterExpressionEngine(),
		artStyleEngine:   ai.NewArtStyleEngine(),
		qualityControl:   ai.NewEnterpriseQualityControl(),
	}
}

// GenerateEnterpriseGrade is the ultimate enterprise generation endpoint
func (e *EnterpriseGenerationService) GenerateEnterpriseGrade(ctx context.Context, req domain.EnterpriseRequest) (*domain.EnterpriseResponse, error) {
	// Step 1: Enhance character expressions and emotions
	characterEnhanced := e.expressionEngine.EnhanceCharacterDescription(req.UserPrompt)
	
	// Step 2: Apply art style and rendering
	styleEnhanced := e.artStyleEngine.ApplyArtStyle(characterEnhanced, req.Style)
	
	if req.Filter != "" {
		styleEnhanced = e.artStyleEngine.ApplyFilter(styleEnhanced, req.Filter)
	}
	
	// Step 3: Create generation request
	genRequest := domain.GenerationRequest{
		UserPrompt: styleEnhanced,
		Options:    req.Options,
		UserID:     req.UserID,
	}
	
	// Step 4: Final review and generation
	finalResponse, err := e.finalGeneration.GenerateWithFinalReview(ctx, genRequest)
	if err != nil {
		return nil, err
	}
	
	// Step 5: Enterprise quality control
	qualityReport := e.qualityControl.VerifyEnterpriseQuality(finalResponse)
	
	return &domain.EnterpriseResponse{
		FinalGenerationResponse: *finalResponse,
		QualityReport:           qualityReport,
		EnterpriseGrade:         qualityReport.Passed,
		Confidence:              qualityReport.Confidence,
	}, nil
}
