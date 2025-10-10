package ai

type BackgroundIntelligence struct {
	backgroundTypes map[string]BackgroundType
	noisePatterns   map[string]NoisePattern
	busyBackgrounds map[string]BusyBackground
	modernSettings  map[string]ModernSetting
}

type BackgroundType struct {
	Name        string
	Complexity  string
	Mood        string
	BestFor     []string
	AvoidWith   []string
	Description string
}

type NoisePattern struct {
	Type        string
	Intensity   string
	VisualEffect string
	AppropriateUse []string
}

type BusyBackground struct {
	Type        string
	Elements    []string
	Organization string
	VisualWeight string
	Handling     string
}

func NewBackgroundIntelligence() *BackgroundIntelligence {
	bi := &BackgroundIntelligence{
		backgroundTypes: make(map[string]BackgroundType),
		noisePatterns:   make(map[string]NoisePattern),
		busyBackgrounds: make(map[string]BusyBackground),
		modernSettings:  make(map[string]ModernSetting),
	}
	
	bi.initializeBackgroundTypes()
	bi.initializeNoisePatterns() 
	bi.initializeBusyBackgrounds()
	bi.initializeModernSettings()
	
	return bi
}

func (b *BackgroundIntelligence) initializeBackgroundTypes() {
	// Simple backgrounds
	b.backgroundTypes["minimalist"] = BackgroundType{
		Name:        "Minimalist",
		Complexity:  "low",
		Mood:        "clean, focused, modern",
		BestFor:     []string{"product shots", "portraits", "fashion"},
		AvoidWith:   []string{"busy characters", "complex outfits"},
		Description: "clean simple background with single color or subtle gradient",
	}
	
	b.backgroundTypes["gradient"] = BackgroundType{
		Name:        "Gradient",
		Complexity:  "low-medium", 
		Mood:        "professional, elegant, modern",
		BestFor:     []string{"professional portraits", "commercial", "fashion"},
		AvoidWith:   []string{"gradient outfits", "rainbow themes"},
		Description: "smooth color transition background",
	}
	
	// Noisy/Busy backgrounds
	b.backgroundTypes["noisy_busy"] = BackgroundType{
		Name:        "Noisy Busy",
		Complexity:  "high",
		Mood:        "urban, lively, chaotic",
		BestFor:     []string{"street photography", "documentary", "urban fashion"},
		AvoidWith:   []string{"detailed outfits", "complex hairstyles"},
		Description: "background with visual noise and multiple elements",
	}
	
	b.backgroundTypes["artistic_busy"] = BackgroundType{
		Name:        "Artistic Busy", 
		Complexity:  "high",
		Mood:        "creative, dynamic, expressive",
		BestFor:     []string{"art projects", "creative portraits", "editorial"},
		AvoidWith:   []string{"minimalist aesthetics", "corporate themes"},
		Description: "intentionally busy background for artistic effect",
	}
	
	// Modern backgrounds
	b.backgroundTypes["modern"] = BackgroundType{
		Name:        "Modern",
		Complexity:  "medium",
		Mood:        "contemporary, sleek, sophisticated",
		BestFor:     []string{"modern fashion", "tech products", "contemporary art"},
		AvoidWith:   []string{"historical themes", "vintage styles"},
		Description: "clean modern elements with contemporary design",
	}
	
	b.backgroundTypes["post_modern"] = BackgroundType{
		Name:        "Post Modern",
		Complexity:  "high",
		Mood:        "eclectic, ironic, mixed",
		BestFor:     []string{"conceptual art", "fashion editorial", "art photography"},
		AvoidWith:   []string{"traditional themes", "pure minimalism"},
		Description: "mixed elements from different eras and styles",
	}
}

func (b *BackgroundIntelligence) initializeNoisePatterns() {
	b.noisePatterns["urban_noise"] = NoisePattern{
		Type:        "Urban Noise",
		Intensity:   "high",
		VisualEffect: "gritty texture, random patterns",
		AppropriateUse: []string{"street photography", "urban fashion", "documentary"},
	}
	
	b.noisePatterns["digital_noise"] = NoisePattern{
		Type:        "Digital Noise",
		Intensity:   "medium", 
		VisualEffect: "pixel patterns, glitch effects",
		AppropriateUse: []string{"cyber themes", "digital art", "tech fashion"},
	}
	
	b.noisePatterns["artistic_noise"] = NoisePattern{
		Type:        "Artistic Noise",
		Intensity:   "variable",
		VisualEffect: "paint textures, brush strokes",
		AppropriateUse: []string{"art portraits", "painterly styles", "mixed media"},
	}
}

// AnalyzeBackgroundType determines optimal background for prompt
func (b *BackgroundIntelligence) AnalyzeBackgroundType(prompt string) string {
	// Extract key elements from prompt
	elements := b.extractBackgroundElements(prompt)
	
	// Determine complexity needs
	complexity := b.determineRequiredComplexity(prompt)
	
	// Match background type to content
	backgroundType := b.matchBackgroundToContent(elements, complexity)
	
	return backgroundType
}

// HandleBusyBackgrounds manages noisy/complex backgrounds
func (b *BackgroundIntelligence) HandleBusyBackgrounds(prompt string, backgroundType string) string {
	if backgroundType != "noisy_busy" && backgroundType != "artistic_busy" {
		return prompt
	}
	
	enhanced := prompt
	
	// Add background handling instructions
	enhanced += ", carefully composed to balance busy background with clear subject focus"
	enhanced += ", strategic depth of field to separate subject from background"
	enhanced += ", professional background blur where appropriate"
	
	// Add noise pattern handling
	if strings.Contains(prompt, "noisy") || strings.Contains(prompt, "busy") {
		enhanced += ", controlled visual noise for artistic effect"
		enhanced += ", balanced composition preventing visual chaos"
	}
	
	return enhanced
}
