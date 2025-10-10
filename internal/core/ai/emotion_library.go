package ai

type EmotionLibrary struct {
	emotions map[string]EmotionDetail
	vibes    map[string]VibeDetail
}

type EmotionDetail struct {
	IntensityRange   [2]float64
	FacialExpressions []string
	BodyLanguage     []string
	EyeExpressions   []string
	MouthExpressions []string
	AppropriateContexts []string
}

type VibeDetail struct {
	EnergyLevel      string
	ColorPalette     []string
	LightingStyle    string
	Composition      string
	MoodKeywords     []string
}

func NewEmotionLibrary() *EmotionLibrary {
	lib := &EmotionLibrary{
		emotions: make(map[string]EmotionDetail),
		vibes:    make(map[string]VibeDetail),
	}
	
	lib.initializeEmotions()
	lib.initializeVibes()
	
	return lib
}

func (e *EmotionLibrary) initializeEmotions() {
	// Confident emotions
	e.emotions["confident"] = EmotionDetail{
		IntensityRange: [2]float64{0.6, 1.0},
		FacialExpressions: []string{
			"direct eye contact", "slight smirk", "raised chin",
			"relaxed brow", "self-assured gaze",
		},
		BodyLanguage: []string{
			"shoulders back", "upright posture", "open stance",
			"hands on hips", "balanced weight distribution",
		},
		AppropriateContexts: []string{"professional", "leadership", "fashion"},
	}
	
	e.emotions["playful"] = EmotionDetail{
		IntensityRange: [2]float64{0.4, 0.8},
		FacialExpressions: []string{
			"wink", "grin", "head tilt", "sparkling eyes",
			"mischievous smile", "raised eyebrows",
		},
		BodyLanguage: []string{
			"dynamic pose", "weight shift", "arm gestures",
			"playful lean", "energetic stance",
		},
		AppropriateContexts: []string{"casual", "youthful", "beach"},
	}
	
	e.emotions["serene"] = EmotionDetail{
		IntensityRange: [2]float64{0.3, 0.6},
		FacialExpressions: []string{
			"soft gaze", "gentle smile", "relaxed features",
			"peaceful expression", "half-closed eyes",
		},
		BodyLanguage: []string{
			"calm posture", "slow movements", "graceful lines",
			"balanced pose", "minimal tension",
		},
		AppropriateContexts: []string{"yoga", "meditation", "nature"},
	}
	
	e.emotions["melancholy"] = EmotionDetail{
		IntensityRange: [2]float64{0.5, 0.9},
		FacialExpressions: []string{
			"distant gaze", "slight frown", "pensive look",
			"soft sadness", "contemplative expression",
		},
		BodyLanguage: []string{
			"slightly slumped", "arms crossed", "protective posture",
			"slow movements", "introspective stance",
		},
		AppropriateContexts: []string{"dramatic", "artistic", "storytelling"},
	}
	
	e.emotions["joyful"] = EmotionDetail{
		IntensityRange: [2]float64{0.7, 1.0},
		FacialExpressions: []string{
			"bright smile", "laughing eyes", "cheerful expression",
			"beaming face", "happy grin",
		},
		BodyLanguage: []string{
			"energetic pose", "open arms", "jumping motion",
			"dynamic movement", "expressive gestures",
		},
		AppropriateContexts: []string{"celebration", "beach", "sports"},
	}
}

func (e *EmotionLibrary) initializeVibes() {
	e.vibes["glamour"] = VibeDetail{
		EnergyLevel:   "sophisticated",
		ColorPalette:  []string{"metallic", "jewel tones", "black and white"},
		LightingStyle: "dramatic studio lighting",
		Composition:   "elegant poses, clean backgrounds",
		MoodKeywords:  []string{"sophisticated", "elegant", "polished"},
	}
	
	e.vibes["vintage"] = VibeDetail{
		EnergyLevel:   "nostalgic",
		ColorPalette:  []string{"sepia tones", "faded colors", "warm hues"},
		LightingStyle: "soft natural light",
		Composition:   "film grain texture, classic framing",
		MoodKeywords:  []string{"nostalgic", "timeless", "classic"},
	}
	
	e.vibes["ethereal"] = VibeDetail{
		EnergyLevel:   "dreamy",
		ColorPalette:  []string{"pastels", "soft hues", "luminous colors"},
		LightingStyle: "soft glow, backlighting",
		Composition:   "flowing elements, soft focus",
		MoodKeywords:  []string{"dreamy", "magical", "soft"},
	}
	
	e.vibes["edgy"] = VibeDetail{
		EnergyLevel:   "high",
		ColorPalette:  []string{"high contrast", "bold colors", "monochrome"},
		LightingStyle: "dramatic shadows",
		Composition:   "dynamic angles, urban settings",
		MoodKeywords:  []string{"bold", "urban", "contemporary"},
	}
}
