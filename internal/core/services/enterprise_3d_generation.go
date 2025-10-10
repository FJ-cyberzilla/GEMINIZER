package services

type Enterprise3DGenerationService struct {
	enterpriseGen   *EnterpriseGenerationService
	anatomyEngine   *ai.ThreeDAnatomyEngine
	poseLibrary     *ai.ProfessionalPoseLibrary
	sceneMatcher    *ai.SceneMatchingEngine
	studioKnowledge *ai.StudioKnowledge
}

func NewEnterprise3DGenerationService(repo HistoryRepository, logger Logger) *Enterprise3DGenerationService {
	return &Enterprise3DGenerationService{
		enterpriseGen:   NewEnterpriseGenerationService(repo, logger),
		anatomyEngine:   ai.NewThreeDAnatomyEngine(),
		poseLibrary:     ai.NewProfessionalPoseLibrary(),
		sceneMatcher:    ai.NewSceneMatchingEngine(),
		studioKnowledge: ai.NewStudioKnowledge(),
	}
}

// Generate3DProfessional handles complete 3D-aware generation
func (e *Enterprise3DGenerationService) Generate3DProfessional(ctx context.Context, req domain.Enterprise3DRequest) (*domain.Enterprise3DResponse, error) {
	// Step 1: 3D Pose Analysis and Enhancement
	poseEnhanced := e.anatomyEngine.EnhancePoseDescription(req.UserPrompt)
	
	// Step 2: Scene Background Matching
	sceneEnhanced := e.sceneMatcher.EnhanceSceneDescription(poseEnhanced)
	
	// Step 3: Studio Setup Application
	studioSetup := e.studioKnowledge.GetProfessionalStudioSetup(req.ShotType, req.Mood)
	studioEnhanced := e.applyStudioSetup(sceneEnhanced, studioSetup)
	
	// Step 4: Enterprise Generation
	enterpriseReq := domain.EnterpriseRequest{
		UserPrompt: studioEnhanced,
		Options:    req.Options,
		Style:      req.Style,
		Filter:     req.Filter,
		UserID:     req.UserID,
	}
	
	enterpriseResp, err := e.enterpriseGen.GenerateEnterpriseGrade(ctx, enterpriseReq)
	if err != nil {
		return nil, err
	}
	
	// Step 5: 3D Quality Verification
	quality3D := e.verify3DQuality(enterpriseResp, req)
	
	return &domain.Enterprise3DResponse{
		EnterpriseResponse: *enterpriseResp,
		Quality3D:          quality3D,
		PoseAnalysis:       e.anatomyEngine.AnalyzePose3D(req.UserPrompt),
		SceneMatch:         e.sceneMatcher.MatchBackgroundToCharacter(req.UserPrompt, "", req.Style),
	}, nil
}

// GetPoseRecommendations suggests professional poses
func (e *Enterprise3DGenerationService) GetPoseRecommendations(context ai.PoseContext) []ai.ProfessionalPose {
	return e.poseLibrary.GetPoseRecommendation(context)
}

// VerifyPosePhysics checks pose physical validity
func (e *Enterprise3DGenerationService) VerifyPosePhysics(poseDescription string) *ai.PhysicsValidation {
	return e.anatomyEngine.ValidatePosePhysics(poseDescription)
}
