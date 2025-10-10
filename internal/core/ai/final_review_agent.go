package ai

type FinalReviewAgent struct {
	knowledgeBase    *SportsKnowledgeBase
	physicsValidator *PhysicsValidator
	safetyAnalyzer   *AdvancedSafetyAnalyzer
	consistencyChecker *ConsistencyChecker
	styleEnforcer    *StyleEnforcer
}

func NewFinalReviewAgent() *FinalReviewAgent {
	return &FinalReviewAgent{
		knowledgeBase:    NewSportsKnowledgeBase(),
		physicsValidator: NewPhysicsValidator(),
		safetyAnalyzer:   NewAdvancedSafetyAnalyzer(),
		consistencyChecker: NewConsistencyChecker(),
		styleEnforcer:    NewStyleEnforcer(),
	}
}

// ReviewAndFinalizePrompt is the final gatekeeper before generation
func (f *FinalReviewAgent) ReviewAndFinalizePrompt(prompt string, context GenerationContext) (*FinalPrompt, error) {
	review := &PromptReview{
		OriginalPrompt: prompt,
		Context:        context,
	}
	
	// Step 1: Knowledge Base Validation
	review.KnowledgeCheck = f.knowledgeBase.ValidateSportsKnowledge(prompt)
	
	// Step 2: Physics Validation
	review.PhysicsCheck = f.physicsValidator.ValidatePhysics(prompt)
	
	// Step 3: Advanced Safety Analysis
	review.SafetyCheck = f.safetyAnalyzer.AnalyzeAdvancedSafety(prompt)
	
	// Step 4: Style Consistency
	review.StyleCheck = f.styleEnforcer.EnforceStyleConsistency(prompt, context.Style)
	
	// Step 5: Generate final optimized prompt
	finalPrompt, err := f.generateFinalPrompt(prompt, review)
	if err != nil {
		return nil, err
	}
	
	review.FinalPrompt = finalPrompt
	review.Approved = f.isApproved(review)
	
	return &FinalPrompt{
		Original:   prompt,
		Final:      finalPrompt,
		Review:     review,
		Confidence: f.calculateConfidence(review),
	}, nil
}

func (f *FinalReviewAgent) generateFinalPrompt(original string, review *PromptReview) (string, error) {
	prompt := original
	
	// Apply knowledge base corrections
	if !review.KnowledgeCheck.IsAccurate {
		prompt = f.knowledgeBase.CorrectTerminology(prompt, review.KnowledgeCheck.Issues)
	}
	
	// Apply physics corrections
	if !review.PhysicsCheck.IsValid {
		prompt = f.physicsValidator.CorrectPhysics(prompt, review.PhysicsCheck.Issues)
	}
	
	// Apply safety modifications
	if !review.SafetyCheck.IsSafe {
		safePrompt, err := f.safetyAnalyzer.MakeSafe(prompt, review.SafetyCheck.Issues)
		if err != nil {
			return "", err
		}
		prompt = safePrompt
	}
	
	// Apply style enhancements
	if !review.StyleCheck.IsConsistent {
		prompt = f.styleEnforcer.EnhanceStyle(prompt, review.StyleCheck.Recommendations)
	}
	
	return prompt, nil
}

func (f *FinalReviewAgent) isApproved(review *PromptReview) bool {
	return review.KnowledgeCheck.IsAccurate &&
		   review.PhysicsCheck.IsValid &&
		   review.SafetyCheck.IsSafe &&
		   review.StyleCheck.IsConsistent
}
