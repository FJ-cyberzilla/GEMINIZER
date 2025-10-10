package ai

type PromptEnhancer struct {
	poseLibrary    *ProfessionalPoseLibrary
	styleEngine    *StyleEngine
	qualityChecker *ProfessionalQualityChecker
}

func NewPromptEnhancer() *PromptEnhancer {
	return &PromptEnhancer{
		poseLibrary:    NewProfessionalPoseLibrary(),
		styleEngine:    NewStyleEngine(),
		qualityChecker: NewProfessionalQualityChecker(),
	}
}

// EnhanceProfessionalPrompt takes amateur prompts and makes them professional
func (p *PromptEnhancer) EnhanceProfessionalPrompt(userPrompt string) *EnhancedPrompt {
	// Step 1: Analyze current prompt quality
	qualityReport := p.qualityChecker.AnalyzePromptQuality(userPrompt)
	
	// Step 2: Detect professional elements
	professionalAnalysis := p.poseLibrary.AnalyzeProfessionalPrompt(userPrompt)
	
	// Step 3: Apply professional enhancements
	enhanced := p.applyProfessionalEnhancements(userPrompt, professionalAnalysis, qualityReport)
	
	// Step 4: Ensure technical completeness
	enhanced = p.ensureTechnicalCompleteness(enhanced)
	
	return &EnhancedPrompt{
		OriginalPrompt: userPrompt,
		EnhancedPrompt: enhanced,
		QualityScore:   qualityReport.Score,
		Improvements:   qualityReport.Suggestions,
		ProfessionalAnalysis: professionalAnalysis,
	}
}

func (p *PromptEnhancer) applyProfessionalEnhancements(prompt string, analysis *ProfessionalAnalysis, quality *QualityReport) string {
	enhanced := prompt
	
	// Add professional photography terminology if missing
	if !quality.HasProfessionalTerms {
		enhanced = p.addProfessionalTerminology(enhanced)
	}
	
	// Enhance pose description if detected
	if analysis.DetectedPose != nil {
		enhanced = p.enhancePoseDescription(enhanced, analysis.DetectedPose)
	}
	
	// Add lighting specifications if missing
	if !quality.HasLightingSpec {
		enhanced = p.addLightingSpecification(enhanced, analysis.DetectedStyle)
	}
	
	// Add camera technical details
	if !quality.HasCameraDetails {
		enhanced = p.addCameraDetails(enhanced)
	}
	
	// Ensure realism requirements
	if !quality.HasRealismRequirements {
		enhanced = p.addRealismRequirements(enhanced)
	}
	
	return enhanced
}

func (p *PromptEnhancer) addProfessionalTerminology(prompt string) string {
	professionalAdditions := []string{
		"8K professional photography",
		"ultra-realistic rendering", 
		"studio quality",
		"editorial grade",
	}
	
	return professionalAdditions[0] + ", " + prompt
}

func (p *PromptEnhancer) enhancePoseDescription(prompt string, pose *ProfessionalPose) string {
	enhancements := []string{
		"professionally posed",
		"anatomically correct posture",
		"natural body positioning",
	}
	
	// Add specific pose enhancements
	if pose.Category == "dynamic" {
		enhancements = append(enhancements, "frozen motion capture", "dynamic energy")
	}
	
	return prompt + ", " + strings.Join(enhancements, ", ")
}

func (p *PromptEnhancer) addLightingSpecification(prompt string, style *PhotographyStyle) string {
	if style != nil {
		return prompt + ", " + style.Lighting
	}
	
	// Default professional lighting
	return prompt + ", professional studio lighting with softboxes and reflectors"
}

func (p *PromptEnhancer) ensureTechnicalCompleteness(prompt string) string {
	technicalRequirements := []string{
		"razor sharp focus",
		"accurate color reproduction", 
		"proper white balance",
		"no lens distortion",
		"professional grade retouching",
	}
	
	return prompt + ", " + strings.Join(technicalRequirements, ", ")
}
