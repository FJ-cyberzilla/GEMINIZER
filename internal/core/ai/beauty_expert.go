package ai

type CosmeticsExpert struct {
	makeupStyles    map[string]MakeupStyle
	beautyStandards *BeautyStandards
}

type MakeupStyle struct {
	Name        string
	Intensity   string
	Components  []string
	Occasions   []string
	AppropriateFor []string
}

type HairstyleEngine struct {
	hairstyles     map[string]HairstyleDetail
	hairPhysics    *HairPhysics
}

type HairstyleDetail struct {
	Name        string
	Length      string
	Texture     string
	Movement    string
	AppropriateFor []string
}

func NewCosmeticsExpert() *CosmeticsExpert {
	return &CosmeticsExpert{
		makeupStyles:    make(map[string]MakeupStyle),
		beautyStandards: NewBeautyStandards(),
	}
}

func NewHairstyleEngine() *HairstyleEngine {
	return &HairstyleEngine{
		hairstyles:  make(map[string]HairstyleDetail),
		hairPhysics: NewHairPhysics(),
	}
}

func (c *CosmeticsExpert) initializeMakeupStyles() {
	c.makeupStyles["natural"] = MakeupStyle{
		Name:      "Natural Makeup",
		Intensity: "light",
		Components: []string{
			"light foundation", "neutral eyeshadow", "mascara", "lip balm",
		},
		Occasions:   []string{"everyday", "professional", "sports"},
		AppropriateFor: []string{"all ages", "professional settings"},
	}
	
	c.makeupStyles["glamour"] = MakeupStyle{
		Name:      "Glamour Makeup",
		Intensity: "medium",
		Components: []string{
			"full coverage foundation", "smokey eyes", "false lashes", "contouring",
		},
		Occasions:   []string{"evening", "photoshoot", "special events"},
		AppropriateFor: []string{"fashion", "beauty photography"},
	}
	
	c.makeupStyles["editorial"] = MakeupStyle{
		Name:      "Editorial Makeup",
		Intensity: "high",
		Components: []string{
			"artistic elements", "bold colors", "graphic lines", "high fashion",
		},
		Occasions:   []string{"fashion shows", "magazine shoots", "artistic projects"},
		AppropriateFor: []string{"high fashion", "artistic expression"},
	}
}

func (h *HairstyleEngine) initializeHairstyles() {
	h.hairstyles["beach_waves"] = HairstyleDetail{
		Name:    "Beach Waves",
		Length:  "medium to long",
		Texture: "loose waves",
		Movement: "natural flow with wind",
		AppropriateFor: []string{"casual", "beach", "summer"},
	}
	
	h.hairstyles["elegant_updo"] = HairstyleDetail{
		Name:    "Elegant Updo",
		Length:  "any length",
		Texture: "sleek and polished",
		Movement: "structured and secure",
		AppropriateFor: []string{"formal", "professional", "evening"},
	}
	
	h.hairstyles["sleek_ponytail"] = HairstyleDetail{
		Name:    "Sleek Ponytail",
		Length:  "medium to long",
		Texture: "smooth and straight",
		Movement: "minimal movement, clean lines",
		AppropriateFor: []string{"athletic", "professional", "casual"},
	}
}

// EnhanceCosmetics adds professional makeup description
func (c *CosmeticsExpert) EnhanceCosmetics(prompt string, context string) string {
	appropriateStyle := c.determineAppropriateStyle(context)
	
	if appropriateStyle != "" {
		style := c.makeupStyles[appropriateStyle]
		cosmeticDescription := fmt.Sprintf("%s makeup with %s",
			style.Intensity, strings.Join(style.Components, ", "))
		
		return prompt + ", " + cosmeticDescription
	}
	
	return prompt
}

// EnhanceHairstyle adds professional hairstyle description
func (h *HairstyleEngine) EnhanceHairstyle(prompt string, context string) string {
	appropriateStyle := h.determineAppropriateStyle(context)
	
	if appropriateStyle != "" {
		style := h.hairstyles[appropriateStyle]
		hairDescription := fmt.Sprintf("%s hairstyle with %s texture and %s movement",
			style.Name, style.Texture, style.Movement)
		
		// Add hair physics
		physics := h.hairPhysics.GetHairPhysics(style.Name, context)
		
		return prompt + ", " + hairDescription + ", " + physics
	}
	
	return prompt
}
