package ai

type SafetyFilter struct {
	contentScanner    *ContentScanner
	ethicsEngine      *EthicsEngine
	loopDetector      *LoopDetector
	explicitDetector  *ExplicitDetector
	moderationHistory *ModerationHistory
}

func NewSafetyFilter() *SafetyFilter {
	return &SafetyFilter{
		contentScanner:    NewContentScanner(),
		ethicsEngine:      NewEthicsEngine(),
		loopDetector:      NewLoopDetector(),
		explicitDetector:  NewExplicitDetector(),
		moderationHistory: NewModerationHistory(),
	}
}

// ValidateStoryContent checks entire story for safety and ethics
func (s *SafetyFilter) ValidateStoryContent(outline ComicOutline) error {
	// Check for explicit content
	if s.explicitDetector.ContainsExplicitContent(outline.Premise) {
		return fmt.Errorf("story premise contains inappropriate content")
	}
	
	// Check character descriptions
	for _, char := range outline.Characters {
		if err := s.ValidateCharacterDescription(char.Description); err != nil {
			return fmt.Errorf("character '%s': %v", char.Name, err)
		}
	}
	
	// Check theme ethics
	if !s.ethicsEngine.IsThemeAppropriate(outline.Theme) {
		return fmt.Errorf("theme '%s' is not appropriate", outline.Theme)
	}
	
	return nil
}

// ValidatePanelDescription checks individual panel descriptions
func (s *SafetyFilter) ValidatePanelDescription(description string) error {
	// Detect explicit content
	if s.explicitDetector.ContainsExplicitContent(description) {
		s.moderationHistory.RecordViolation("explicit_content", description)
		return fmt.Errorf("panel description contains explicit content")
	}
	
	// Detect infinite loops
	if s.loopDetector.IsInLoop(description) {
		s.moderationHistory.RecordViolation("generation_loop", description)
		return fmt.Errorf("detected generation loop - please rephrase")
	}
	
	// Check for ethical concerns
	if !s.ethicsEngine.IsContentAppropriate(description) {
		s.moderationHistory.RecordViolation("ethical_concern", description)
		return fmt.Errorf("content raises ethical concerns")
	}
	
	// Content scanning for other issues
	scanResult := s.contentScanner.ScanContent(description)
	if !scanResult.IsSafe {
		s.moderationHistory.RecordViolation(scanResult.IssueType, description)
		return fmt.Errorf("content safety issue: %s", scanResult.IssueType)
	}
	
	return nil
}

// ValidateCharacterDescription ensures character descriptions are appropriate
func (s *SafetyFilter) ValidateCharacterDescription(description string) error {
	// Age appropriateness
	if s.explicitDetector.ContainsAgeInappropriate(description) {
		return fmt.Errorf("character description age inappropriate")
	}
	
	// Cultural sensitivity
	if !s.ethicsEngine.IsCulturallySensitive(description) {
		return fmt.Errorf("character description culturally insensitive")
	}
	
	// Stereotype detection
	if s.ethicsEngine.ContainsStereotypes(description) {
		return fmt.Errorf("character description contains harmful stereotypes")
	}
	
	return nil
}

// FinalComicReview comprehensive safety check before generation
func (s *SafetyFilter) FinalComicReview(comic *ComicStory) error {
	// Check all panels
	for i, panel := range comic.Panels {
		if err := s.ValidatePanelDescription(panel.Description); err != nil {
			return fmt.Errorf("panel %d: %v", i+1, err)
		}
		
		// Check dialogue
		if err := s.ValidateDialogue(panel.Dialogue); err != nil {
			return fmt.Errorf("panel %d dialogue: %v", i+1, err)
		}
	}
	
	// Check overall story arc
	if !s.ethicsEngine.IsStoryArcAppropriate(comic) {
		return fmt.Errorf("story arc contains inappropriate elements")
	}
	
	return nil
}
