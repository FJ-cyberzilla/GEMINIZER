package ai

type CharacterExpressionEngine struct {
	emotionLibrary   *EmotionLibrary
	vibeAnalyzer     *VibeAnalyzer
	cosmeticsExpert  *CosmeticsExpert
	hairstyleEngine  *HairstyleEngine
	bodyTypeExpert   *BodyTypeExpert
}

type CharacterVibe struct {
	Mood        string
	Intensity   float64
	Expression  string
	BodyLanguage string
	EnergyLevel string
}

type EmotionProfile struct {
	PrimaryEmotion   string
	SecondaryEmotion string
	Intensity        float64
	PhysicalSigns    []string
	FacialExpression string
}

func NewCharacterExpressionEngine() *CharacterExpressionEngine {
	return &CharacterExpressionEngine{
		emotionLibrary:   NewEmotionLibrary(),
		vibeAnalyzer:     NewVibeAnalyzer(),
		cosmeticsExpert:  NewCosmeticsExpert(),
		hairstyleEngine:  NewHairstyleEngine(),
		bodyTypeExpert:   NewBodyTypeExpert(),
	}
}

// AnalyzeCharacterVibe understands and enhances character emotions
func (c *CharacterExpressionEngine) AnalyzeCharacterVibe(prompt string) *CharacterAnalysis {
	analysis := &CharacterAnalysis{
		OriginalPrompt: prompt,
	}
	
	// Extract emotional content
	analysis.Emotion = c.emotionLibrary.ExtractEmotion(prompt)
	analysis.Vibe = c.vibeAnalyzer.DetectVibe(prompt)
	analysis.Cosmetics = c.cosmeticsExpert.AnalyzeMakeup(prompt)
	analysis.Hairstyle = c.hairstyleEngine.AnalyzeHairstyle(prompt)
	analysis.BodyType = c.bodyTypeExpert.AnalyzeBodyDescription(prompt)
	
	// Ensure appropriateness
	analysis.IsAppropriate = c.ensureAppropriateness(analysis)
	
	return analysis
}

// EnhanceCharacterDescription adds professional emotional and visual depth
func (c *CharacterExpressionEngine) EnhanceCharacterDescription(prompt string) string {
	analysis := c.AnalyzeCharacterVibe(prompt)
	
	if !analysis.IsAppropriate {
		prompt = c.makeAppropriate(prompt, analysis)
	}
	
	// Add emotional depth
	enhanced := c.addEmotionalDepth(prompt, analysis.Emotion)
	
	// Enhance vibe description
	enhanced = c.addVibeDescription(enhanced, analysis.Vibe)
	
	// Refine cosmetics description
	enhanced = c.refineCosmetics(enhanced, analysis.Cosmetics)
	
	// Enhance hairstyle description
	enhanced = c.enhanceHairstyle(enhanced, analysis.Hairstyle)
	
	// Ensure body type appropriateness
	enhanced = c.ensureBodyAppropriateness(enhanced, analysis.BodyType)
	
	return enhanced
}

func (c *CharacterExpressionEngine) addEmotionalDepth(prompt string, emotion EmotionProfile) string {
	emotionalEnhancements := []string{}
	
	// Add primary emotion
	if emotion.PrimaryEmotion != "" {
		emotionalEnhancements = append(emotionalEnhancements,
			fmt.Sprintf("%s expression", emotion.PrimaryEmotion))
	}
	
	// Add physical signs of emotion
	for _, sign := range emotion.PhysicalSigns {
		emotionalEnhancements = append(emotionalEnhancements, sign)
	}
	
	// Add facial expression
	if emotion.FacialExpression != "" {
		emotionalEnhancements = append(emotionalEnhancements,
			fmt.Sprintf("%s facial expression", emotion.FacialExpression))
	}
	
	if len(emotionalEnhancements) > 0 {
		return prompt + ", " + strings.Join(emotionalEnhancements, ", ")
	}
	
	return prompt
}
