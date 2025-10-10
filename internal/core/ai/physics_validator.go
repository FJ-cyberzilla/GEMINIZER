package ai

type PhysicsValidator struct {
	anatomyKnowledge *AnatomyKnowledge
	movementPhysics  *MovementPhysics
	clothingPhysics  *ClothingPhysics
}

func NewPhysicsValidator() *PhysicsValidator {
	return &PhysicsValidator{
		anatomyKnowledge: NewAnatomyKnowledge(),
		movementPhysics:  NewMovementPhysics(),
		clothingPhysics:  NewClothingPhysics(),
	}
}

// ValidatePhysics checks physical plausibility and handles sensitive movements
func (p *PhysicsValidator) ValidatePhysics(prompt string) *PhysicsValidation {
	validation := &PhysicsValidation{
		IsValid: true,
		Issues:  []string{},
	}
	
	// Check for 180-degree leg positions
	if p.containsWideLegPosition(prompt) {
		if !p.isWideLegPositionAppropriate(prompt) {
			validation.IsValid = false
			validation.Issues = append(validation.Issues, "Wide leg position requires careful framing and context")
		}
	}
	
	// Check anatomy plausibility
	if !p.anatomyKnowledge.IsAnatomicallyPlausible(prompt) {
		validation.IsValid = false
		validation.Issues = append(validation.Issues, "Anatomical position may not be physically possible")
	}
	
	// Check clothing physics
	if !p.clothingPhysics.IsClothingPhysicallyAccurate(prompt) {
		validation.IsValid = false
		validation.Issues = append(validation.Issues, "Clothing behavior may not be physically accurate")
	}
	
	// Check movement physics
	if !p.movementPhysics.IsMovementPhysicallyAccurate(prompt) {
		validation.IsValid = false
		validation.Issues = append(validation.Issues, "Movement may not be physically accurate")
	}
	
	return validation
}

func (p *PhysicsValidator) containsWideLegPosition(prompt string) bool {
	wideLegIndicators := []string{
		"180 degree", "180°", "wide stance", "full split",
		"middle split", "side split", "jump split",
	}
	
	lowerPrompt := strings.ToLower(prompt)
	for _, indicator := range wideLegIndicators {
		if strings.Contains(lowerPrompt, indicator) {
			return true
		}
	}
	return false
}

func (p *PhysicsValidator) isWideLegPositionAppropriate(prompt string) bool {
	// Allow wide leg positions in appropriate contexts
	appropriateContexts := []string{
		"gymnastics", "yoga", "dance", "martial arts", "stretching",
		"athletic", "sports", "training", "flexibility",
	}
	
	// Check for mini-skirt context that requires careful handling
	hasMiniSkirt := strings.Contains(strings.ToLower(prompt), "mini skirt") ||
				   strings.Contains(strings.ToLower(prompt), "short skirt")
	
	lowerPrompt := strings.ToLower(prompt)
	
	// If mini-skirt with wide legs, require specific safe framing
	if hasMiniSkirt {
		safeFraming := []string{
			"side angle", "back view", "three quarter", "high angle",
			"strategic framing", "tasteful composition", "modest",
			"leggings underneath", "sports shorts", "appropriate",
		}
		
		for _, framing := range safeFraming {
			if strings.Contains(lowerPrompt, framing) {
				return true
			}
		}
		return false
	}
	
	// Check for appropriate athletic context
	for _, context := range appropriateContexts {
		if strings.Contains(lowerPrompt, context) {
			return true
		}
	}
	
	return false
}

// CorrectPhysics modifies prompts to maintain physical accuracy while being appropriate
func (p *PhysicsValidator) CorrectPhysics(prompt string, issues []string) string {
	corrected := prompt
	
	for _, issue := range issues {
		switch {
		case strings.Contains(issue, "Wide leg position"):
			corrected = p.correctWideLegPosition(corrected)
		case strings.Contains(issue, "Anatomical position"):
			corrected = p.correctAnatomy(corrected)
		case strings.Contains(issue, "Clothing behavior"):
			corrected = p.correctClothingPhysics(corrected)
		case strings.Contains(issue, "Movement"):
			corrected = p.correctMovementPhysics(corrected)
		}
	}
	
	return corrected
}

func (p *PhysicsValidator) correctWideLegPosition(prompt string) string {
	// Add appropriate context and framing for wide leg positions
	enhancements := []string{
		"athletic context",
		"professional sports photography",
		"tasteful composition",
		"appropriate framing",
		"focus on flexibility and strength",
	}
	
	// Add specific framing for mini-skirt situations
	if strings.Contains(strings.ToLower(prompt), "mini skirt") {
		enhancements = append(enhancements,
			"side angle view",
			"strategic camera positioning",
			"modest representation",
			"sports-appropriate attire",
		)
	}
	
	return prompt + ", " + strings.Join(enhancements, ", ")
}
