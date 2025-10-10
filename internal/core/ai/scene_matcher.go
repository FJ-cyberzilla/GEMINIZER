package ai

type SceneMatchingEngine struct {
	backgroundLibrary *BackgroundLibrary
	lightingMatcher   *LightingMatcher
	compositionRules  *CompositionRules
	colorHarmony      *ColorHarmony
	styleConsistency  *StyleConsistency
}

type BackgroundLibrary struct {
	studioBackgrounds map[string]StudioBackground
	environmentScenes map[string]EnvironmentScene
	decorStyles       map[string]DecorationStyle
}

type StudioBackground struct {
	Name        string
	Type        string // gradient, seamless, textured, set
	Colors      []string
	Lighting    string
	Complexity  string
	BestFor     []string
}

type EnvironmentScene struct {
	Name          string
	Location      string
	TimeOfDay     string
	Weather       string
	KeyElements   []string
	Mood          string
	AppropriateFor []string
}

func NewSceneMatchingEngine() *SceneMatchingEngine {
	return &SceneMatchingEngine{
		backgroundLibrary: NewBackgroundLibrary(),
		lightingMatcher:   NewLightingMatcher(),
		compositionRules:  NewCompositionRules(),
		colorHarmony:      NewColorHarmony(),
		styleConsistency:  NewStyleConsistency(),
	}
}

// MatchBackgroundToCharacter finds perfect background for character and pose
func (s *SceneMatchingEngine) MatchBackgroundToCharacter(characterDesc string, poseDesc string, style string) *SceneMatch {
	match := &SceneMatch{
		CharacterDescription: characterDesc,
		PoseDescription:      poseDesc,
		Style:                style,
	}
	
	// Analyze character and pose
	characterAnalysis := s.analyzeCharacter(characterDesc)
	poseAnalysis := s.analyzePose(poseDesc)
	
	// Find matching background
	match.Background = s.findBestBackground(characterAnalysis, poseAnalysis, style)
	
	// Match lighting
	match.Lighting = s.lightingMatcher.MatchLighting(poseAnalysis, characterAnalysis, match.Background)
	
	// Ensure color harmony
	match.ColorPalette = s.colorHarmony.CreateHarmoniousPalette(characterAnalysis, match.Background)
	
	// Check composition
	match.Composition = s.compositionRules.SuggestComposition(poseAnalysis, match.Background)
	
	// Calculate match score
	match.MatchScore = s.calculateMatchScore(match)
	
	return match
}

// EnhanceSceneDescription adds professional background and lighting details
func (s *SceneMatchingEngine) EnhanceSceneDescription(prompt string) string {
	// Extract character and pose information
	character, pose := s.extractCharacterAndPose(prompt)
	
	// Get scene match
	match := s.MatchBackgroundToCharacter(character, pose, "professional")
	
	// Build enhanced description
	enhanced := prompt
	
	// Add background if missing
	if !s.containsBackground(prompt) {
		enhanced += ", " + match.Background.Description
	}
	
	// Add lighting if missing  
	if !s.containsLighting(prompt) {
		enhanced += ", " + match.Lighting.Description
	}
	
	// Add composition guidance
	enhanced += ", " + match.Composition.Guidance
	
	return enhanced
}

func (s *SceneMatchingEngine) findBestBackground(character CharacterAnalysis, pose PoseAnalysis, style string) Background {
	var bestBackground Background
	highestScore := 0.0
	
	allBackgrounds := s.backgroundLibrary.GetAllBackgrounds()
	
	for _, bg := range allBackgrounds {
		score := s.calculateBackgroundScore(bg, character, pose, style)
		if score > highestScore {
			highestScore = score
			bestBackground = bg
		}
	}
	
	return bestBackground
}
