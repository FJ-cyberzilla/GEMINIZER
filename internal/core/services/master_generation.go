package services

type MasterGenerationService struct {
	masterAgent     *ai.MasterPriorityAgent
	enterprise3D    *Enterprise3DGenerationService
	qualityMonitor  *ai.QualityMonitor
}

func NewMasterGenerationService(repo HistoryRepository, logger Logger) *MasterGenerationService {
	return &MasterGenerationService{
		masterAgent:    ai.NewMasterPriorityAgent(),
		enterprise3D:   NewEnterprise3DGenerationService(repo, logger),
		qualityMonitor: ai.NewQualityMonitor(),
	}
}

// GenerateWithMasterControl is the ultimate generation endpoint
func (m *MasterGenerationService) GenerateWithMasterControl(ctx context.Context, req domain.MasterRequest) (*domain.MasterResponse, error) {
	// Step 1: Master analysis and prioritization
	masterAnalysis := m.masterAgent.AnalyzeAndPrioritize(req.UserPrompt)
	
	if !masterAnalysis.QualityCheck.Passed {
		return nil, fmt.Errorf("master quality check failed: %v", masterAnalysis.QualityCheck.Issues)
	}
	
	// Step 2: Create enterprise request with optimized prompt
	enterpriseReq := domain.Enterprise3DRequest{
		UserPrompt: masterAnalysis.OptimizedPrompt,
		Options:    req.Options,
		Style:      masterAnalysis.ArtStyle.PrimaryStyle,
		ShotType:   req.ShotType,
		Mood:       req.Mood,
		UserID:     req.UserID,
	}
	
	// Step 3: Generate with enterprise system
	enterpriseResp, err := m.enterprise3D.Generate3DProfessional(ctx, enterpriseReq)
	if err != nil {
		return nil, err
	}
	
	// Step 4: Master quality verification
	masterQuality := m.qualityMonitor.VerifyMasterQuality(enterpriseResp, masterAnalysis)
	
	return &domain.MasterResponse{
		Enterprise3DResponse: *enterpriseResp,
		MasterAnalysis:       masterAnalysis,
		MasterQuality:        masterQuality,
		OptimizationLog:      m.generateOptimizationLog(masterAnalysis),
	}, nil
}

// GetStyleRecommendations provides intelligent style suggestions
func (m *MasterGenerationService) GetStyleRecommendations(context ai.StyleContext) []ai.StyleRecommendation {
	return m.masterAgent.GetStyleRecommendations(context)
}

// ValidatePrompt checks for potential conflicts before generation
func (m *MasterGenerationService) ValidatePrompt(prompt string) *ai.ValidationResult {
	masterAnalysis := m.masterAgent.AnalyzeAndPrioritize(prompt)
	
	return &ai.ValidationResult{
		IsValid:      masterAnalysis.QualityCheck.Passed,
		Conflicts:    masterAnalysis.ConflictsDetected,
		Suggestions:  masterAnalysis.OptimizationSuggestions,
		Confidence:   masterAnalysis.FinalPriorities.Confidence,
	}
}
