package ai

type ConflictResolver struct {
	conflictRules []ConflictRule
	priorityRules []PriorityRule
	styleConflicts map[string][]string
	culturalConflicts map[string][]string
}

type ConflictRule struct {
	Name        string
	Condition   string
	Resolution  string
	Priority    int
}

type PriorityRule struct {
	Element     string
	Priority    int
	Overrides   []string
}

func NewConflictResolver() *ConflictResolver {
	cr := &ConflictResolver{
		conflictRules:    []ConflictRule{},
		priorityRules:    []PriorityRule{},
		styleConflicts:   make(map[string][]string),
		culturalConflicts: make(map[string][]string),
	}
	
	cr.initializeConflictRules()
	cr.initializePriorityRules()
	cr.initializeStyleConflicts()
	cr.initializeCulturalConflicts()
	
	return cr
}

func (c *ConflictResolver) initializeConflictRules() {
	// Art style conflicts
	c.conflictRules = append(c.conflictRules, ConflictRule{
		Name: "Realism vs Anime",
		Condition: "realistic_figure and anime detected",
		Resolution: "prioritize realistic_figure for human characters",
		Priority: 9,
	})
	
	c.conflictRules = append(c.conflictRules, ConflictRule{
		Name: "Historical vs Futuristic", 
		Condition: "baroque and cyberpunk detected",
		Resolution: "choose based on character description priority",
		Priority: 8,
	})
	
	c.conflictRules = append(c.conflictRules, ConflictRule{
		Name: "Minimalist vs Busy",
		Condition: "minimalist background and busy character",
		Resolution: "adjust background complexity to match character",
		Priority: 7,
	})
	
	// Cultural conflicts
	c.conflictRules = append(c.conflictRules, ConflictRule{
		Name: "Mixed Cultural Physics",
		Condition: "multiple regional physics detected",
		Resolution: "prioritize cultural context from character description",
		Priority: 9,
	})
	
	// Background conflicts
	c.conflictRules = append(c.conflictRules, ConflictRule{
		Name: "Background Overwhelms Subject",
		Condition: "busy background with detailed subject",
		Resolution: "simplify background or enhance subject separation",
		Priority: 8,
	})
}

func (c *ConflictResolver) initializePriorityRules() {
	c.priorityRules = []PriorityRule{
		{
			Element:   "character_description",
			Priority:  10,
			Overrides: []string{"background", "lighting", "style"},
		},
		{
			Element:   "cultural_context", 
			Priority:  9,
			Overrides: []string{"art_style", "physics"},
		},
		{
			Element:   "art_style",
			Priority:  8,
			Overrides: []string{"background", "lighting"},
		},
		{
			Element:   "physics_accuracy",
			Priority:  7,
			Overrides: []string{"stylization"},
		},
		{
			Element:   "background",
			Priority:  6,
			Overrides: []string{},
		},
	}
}

// RemoveConflicts applies conflict resolution to prompt
func (c *ConflictResolver) RemoveConflicts(prompt string, priorities *GenerationPriority) string {
	resolved := prompt
	
	// Apply style conflicts resolution
	resolved = c.resolveStyleConflicts(resolved, priorities.ArtStyle)
	
	// Apply cultural conflicts resolution  
	resolved = c.resolveCulturalConflicts(resolved, priorities.CulturalContext)
	
	// Apply background conflicts resolution
	resolved = c.resolveBackgroundConflicts(resolved, priorities.BackgroundType)
	
	// Apply physics conflicts resolution
	resolved = c.resolvePhysicsConflicts(resolved, priorities.PhysicsRegion)
	
	return resolved
}

func (c *ConflictResolver) resolveStyleConflicts(prompt string, primaryStyle string) string {
	resolved := prompt
	
	// Remove conflicting style mentions
	for style, conflicts := range c.styleConflicts {
		if style != primaryStyle {
			for _, conflict := range conflicts {
				if strings.Contains(resolved, conflict) {
					resolved = strings.ReplaceAll(resolved, conflict, "")
				}
			}
		}
	}
	
	return resolved
}
