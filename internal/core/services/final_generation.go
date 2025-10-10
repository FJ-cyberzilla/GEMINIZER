package services

type FinalGenerationService struct {
	imageGenerator  *EnhancedImageGenerator
	finalReview     *ai.FinalReviewAgent
	qualityAssurance *ai.QualityAssurance
}

func NewFinalGenerationService(repo HistoryRepository, logger Logger) *FinalGenerationService {
	return &FinalGenerationService{
		imageGenerator:   NewEnhancedImageGenerator(repo, logger),
		finalReview:      ai.NewFinalReviewAgent(),
		qualityAssurance: ai.NewQualityAssurance(),
	}
}

// GenerateWithFinalReview is the ultimate generation endpoint
func (f *FinalGenerationService) GenerateWithFinalReview(ctx context.Context, req domain.GenerationRequest) (*domain.FinalGenerationResponse, error) {
	// Step 1: Initial enhancement
	enhancedResponse, err := f.imageGenerator.GenerateWithAnalysis(ctx, req)
	if err != nil {
		return nil, err
	}
	
	// Step 2: Final review by expert agent
	finalPrompt, err := f.finalReview.ReviewAndFinalizePrompt(
		enhancedResponse.EnrichedPrompt,
		ai.GenerationContext{
			Style:    req.Options.Style,
			Intent:   f.detectIntent(req.UserPrompt),
			UserLevel: f.assessUserLevel(ctx, req.UserID),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("final review failed: %v", err)
	}
	
	if !finalPrompt.Review.Approved {
		return nil, fmt.Errorf("prompt not approved: %v", finalPrompt.Review.GetIssues())
	}
	
	// Step 3: Quality assurance
	qualityCheck := f.qualityAssurance.VerifyQuality(finalPrompt.Final)
	if !qualityCheck.Passed {
		return nil, fmt.Errorf("quality assurance failed: %v", qualityCheck.Issues)
	}
	
	// Step 4: Final generation with approved prompt
	finalRequest := req
	finalRequest.UserPrompt = finalPrompt.Final
	
	finalResult, err := f.imageGenerator.GenerateImage(finalRequest.UserPrompt, finalRequest.Options)
	if err != nil {
		return nil, err
	}
	
	return &domain.FinalGenerationResponse{
		GenerationResponse: *enhancedResponse,
		FinalPrompt:        finalPrompt.Final,
		Review:             finalPrompt.Review,
		QualityCheck:       qualityCheck,
		Confidence:         finalPrompt.Confidence,
	}, nil
}
