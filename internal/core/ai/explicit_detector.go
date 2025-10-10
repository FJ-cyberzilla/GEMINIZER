package ai

type ExplicitDetector struct {
	bannedTerms      map[string]bool
	suspiciousPatterns []*regexp.Regexp
	contextAnalyzer  *ContextAnalyzer
	culturalFilters  *CulturalFilters
}

func NewExplicitDetector() *ExplicitDetector {
	detector := &ExplicitDetector{
		bannedTerms: make(map[string]bool),
		suspiciousPatterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)nude|naked|bare`),
			regexp.MustCompile(`(?i)sexual|erotic|sensual`),
			regexp.MustCompile(`(?i)violence|blood|gore`),
			regexp.MustCompile(`(?i)hate|racist|discriminat`),
		},
		contextAnalyzer: NewContextAnalyzer(),
		culturalFilters: NewCulturalFilters(),
	}
	
	detector.initializeBannedTerms()
	return detector
}

func (e *ExplicitDetector) initializeBannedTerms() {
	// Comprehensive list of inappropriate terms
	bannedLists := []string{
		"explicit", "porn", "xxx", "adult", "nsfw",
		"violence", "gore", "blood", "weapon",
		"hate", "racist", "discriminate", "offensive",
		"illegal", "criminal", "harmful",
	}
	
	for _, term := range bannedLists {
		e.bannedTerms[strings.ToLower(term)] = true
	}
}

// ContainsExplicitContent comprehensive explicit content detection
func (e *ExplicitDetector) ContainsExplicitContent(text string) bool {
	lowerText := strings.ToLower(text)
	
	// Check banned terms
	for term := range e.bannedTerms {
		if strings.Contains(lowerText, term) {
			return true
		}
	}
	
	// Check suspicious patterns
	for _, pattern := range e.suspiciousPatterns {
		if pattern.MatchString(text) {
			// Analyze context to avoid false positives
			if !e.contextAnalyzer.IsContextAppropriate(text, pattern.String()) {
				return true
			}
		}
	}
	
	// Cultural sensitivity check
	if !e.culturalFilters.IsCulturallyAppropriate(text) {
		return true
	}
	
	return false
}

// ContainsAgeInappropriate checks for content unsuitable for all ages
func (e *ExplicitDetector) ContainsAgeInappropriate(text string) bool {
	ageInappropriatePatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)minor|underage|teen`),
		regexp.MustCompile(`(?i)school|classroom`),
		regexp.MustCompile(`(?i)child|kid|baby`),
	}
	
	for _, pattern := range ageInappropriatePatterns {
		if pattern.MatchString(text) {
			// Check if context is inappropriate
			if e.contextAnalyzer.IsAgeInappropriateContext(text) {
				return true
			}
		}
	}
	
	return false
}

// SafeAlternative generates safe alternative when explicit content detected
func (e *ExplicitDetector) SafeAlternative(originalText string) string {
	return e.contextAnalyzer.GenerateSafeAlternative(originalText)
}
