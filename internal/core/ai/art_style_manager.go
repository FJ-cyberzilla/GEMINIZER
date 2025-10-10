package ai

type ArtStyleManager struct {
	styleDefinitions map[string]ArtStyleDefinition
	styleHierarchy   map[string][]string
	conflictRules    map[string]ConflictRule
	modernEras       map[string]ModernEra
}

type ArtStyleDefinition struct {
	Name           string
	Category       string
	Era            string
	KeyCharacteristics []string
	ConflictsWith  []string
	CompatibleWith []string
	Priority       int
}

type ModernEra struct {
	Name        string
	Years       string
	Aesthetics  []string
	Technology  string
	Culture     string
}

func NewArtStyleManager() *ArtStyleManager {
	manager := &ArtStyleManager{
		styleDefinitions: make(map[string]ArtStyleDefinition),
		styleHierarchy:   make(map[string][]string),
		conflictRules:    make(map[string]ConflictRule),
		modernEras:       make(map[string]ModernEra),
	}
	
	manager.initializeArtStyles()
	manager.initializeModernEras()
	manager.initializeConflictRules()
	
	return manager
}

func (a *ArtStyleManager) initializeArtStyles() {
	// Realistic Figures
	a.styleDefinitions["realistic_figure"] = ArtStyleDefinition{
		Name:           "Realistic Figure",
		Category:       "realism",
		Era:            "timeless",
		KeyCharacteristics: []string{
			"accurate human proportions", "natural lighting", "detailed textures",
			"anatomical correctness", "lifelike rendering",
		},
		ConflictsWith:  []string{"anime", "cartoon", "abstract"},
		CompatibleWith: []string{"hyper_realistic", "3d_figure", "photorealism"},
		Priority:       8,
	}
	
	a.styleDefinitions["hyper_realistic"] = ArtStyleDefinition{
		Name:           "Hyper Realistic",
		Category:       "realism",
		Era:            "contemporary", 
		KeyCharacteristics: []string{
			"extreme detail", "micro-textures", "subsurface scattering",
			"perfect lighting", "flawless rendering", "beyond reality",
		},
		ConflictsWith:  []string{"impressionist", "expressionist", "minimalist"},
		CompatibleWith: []string{"realistic_figure", "3d_figure", "digital_art"},
		Priority:       9,
	}
	
	a.styleDefinitions["3d_figure"] = ArtStyleDefinition{
		Name:           "3D Figure",
		Category:       "digital",
		Era:            "contemporary",
		KeyCharacteristics: []string{
			"3d rendering", "cg quality", "digital modeling",
			"perfect lighting", "clean surfaces", "computer generated",
		},
		ConflictsWith:  []string{"traditional", "painterly", "sketch"},
		CompatibleWith: []string{"hyper_realistic", "realistic_figure", "digital_art"},
		Priority:       7,
	}
	
	// Historical Styles
	a.styleDefinitions["baroque"] = ArtStyleDefinition{
		Name:           "Baroque",
		Category:       "historical",
		Era:            "17th-18th century",
		KeyCharacteristics: []string{
			"dramatic lighting", "rich colors", "emotional intensity",
			"theatrical compositions", "ornate details", "chiaroscuro",
		},
		ConflictsWith:  []string{"minimalist", "modern", "cyberpunk"},
		CompatibleWith: []string{"classical", "renaissance", "dramatic"},
		Priority:       6,
	}
	
	// Modern & Contemporary
	a.styleDefinitions["modern"] = ArtStyleDefinition{
		Name:           "Modern",
		Category:       "modern",
		Era:            "20th century", 
		KeyCharacteristics: []string{
			"clean lines", "simplified forms", "experimental",
			"abstract elements", "geometric shapes", "minimal color",
		},
		ConflictsWith:  []string{"baroque", "rococo", "traditional"},
		CompatibleWith: []string{"contemporary", "minimalist", "abstract"},
		Priority:       5,
	}
	
	a.styleDefinitions["post_modern"] = ArtStyleDefinition{
		Name:           "Post Modern",
		Category:       "contemporary",
		Era:            "late 20th century",
		KeyCharacteristics: []string{
			"eclectic mixing", "ironic", "self-referential",
			"breaking rules", "cultural references", "hybrid styles",
		},
		ConflictsWith:  []string{"classical", "pure styles", "traditional"},
		CompatibleWith: []string{"modern", "contemporary", "mixed_media"},
		Priority:       4,
	}
	
	// Cyber & Futuristic
	a.styleDefinitions["cyberpunk"] = ArtStyleDefinition{
		Name:           "Cyberpunk",
		Category:       "futuristic",
		Era:            "contemporary",
		KeyCharacteristics: []string{
			"neon lighting", "high tech low life", "dystopian",
			"urban decay", "technological fusion", "rainy nights",
		},
		ConflictsWith:  []string{"historical", "natural", "pastoral"},
		CompatibleWith: []string{"futuristic", "sci_fi", "urban"},
		Priority:       7,
	}
	
	a.styleDefinitions["y2k"] = ArtStyleDefinition{
		Name:           "Y2K",
		Category:       "nostalgic_futuristic", 
		Era:            "1990s-2000s",
		KeyCharacteristics: []string{
			"futuristic optimism", "metallic surfaces", "translucent plastics",
			"bright colors", "digital patterns", "retro future",
		},
		ConflictsWith:  []string{"medieval", "historical", "naturalistic"},
		CompatibleWith: []string{"retro", "futuristic", "digital"},
		Priority:       5,
	}
}

func (a *ArtStyleManager) initializeModernEras() {
	a.modernEras["modern"] = ModernEra{
		Name:       "Modern",
		Years:      "1900-1960",
		Aesthetics: []string{"minimalism", "functionalism", "abstraction"},
		Technology: "industrial age",
		Culture:    "western avant-garde",
	}
	
	a.modernEras["post_modern"] = ModernEra{
		Name:       "Post Modern", 
		Years:      "1960-1990",
		Aesthetics: []string{"eclecticism", "pastiche", "deconstruction"},
		Technology: "information age beginning",
		Culture:    "global multicultural",
	}
	
	a.modernEras["contemporary"] = ModernEra{
		Name:       "Contemporary",
		Years:      "1990-present", 
		Aesthetics: []string{"digital", "global", "hybrid"},
		Technology: "digital revolution",
		Culture:    "global internet culture",
	}
}

// DetectAndPrioritizeStyle identifies and ranks art styles in prompt
func (a *ArtStyleManager) DetectAndPrioritizeStyle(prompt string) *StyleAnalysis {
	analysis := &StyleAnalysis{
		DetectedStyles: make(map[string]float64),
	}
	
	// Detect all mentioned styles
	for styleName, styleDef := range a.styleDefinitions {
		confidence := a.detectStyleConfidence(prompt, styleName, styleDef)
		if confidence > 0.3 {
			analysis.DetectedStyles[styleName] = confidence
		}
	}
	
	// Resolve conflicts and set primary style
	analysis.PrimaryStyle = a.resolveStyleConflicts(analysis.DetectedStyles)
	analysis.Confidence = analysis.DetectedStyles[analysis.PrimaryStyle]
	
	// Set secondary styles
	analysis.SecondaryStyles = a.getCompatibleStyles(analysis.PrimaryStyle, analysis.DetectedStyles)
	
	return analysis
}

func (a *ArtStyleManager) resolveStyleConflicts(detectedStyles map[string]float64) string {
	var primaryStyle string
	highestPriority := -1
	
	for styleName, confidence := range detectedStyles {
		styleDef := a.styleDefinitions[styleName]
		
		// Weight by both confidence and priority
		weightedScore := float64(styleDef.Priority) * confidence
		
		if weightedScore > float64(highestPriority) {
			highestPriority = int(weightedScore)
			primaryStyle = styleName
		}
	}
	
	return primaryStyle
}
