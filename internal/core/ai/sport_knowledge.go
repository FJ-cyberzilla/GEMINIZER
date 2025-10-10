package ai

type SportsKnowledgeBase struct {
	yogaPoses      map[string]YogaPose
	swimmingStyles map[string]SwimmingStyle
	gymnasticsMoves map[string]GymnasticsMove
	beachSettings  map[string]BeachEnvironment
	clothingTypes  map[string]ClothingKnowledge
}

type YogaPose struct {
	Name          string
	Difficulty    string
	Angles        []string // front, side, back, top
	Modifications []string
	SafetyNotes   string
}

type SwimmingStyle struct {
	Name        string
	Environment string // pool, sea, ocean
	Physics     string
	Clothing    string
}

func NewSportsKnowledgeBase() *SportsKnowledgeBase {
	kb := &SportsKnowledgeBase{
		yogaPoses:      make(map[string]YogaPose),
		swimmingStyles: make(map[string]SwimmingStyle),
		gymnasticsMoves: make(map[string]GymnasticsMove),
		beachSettings:  make(map[string]BeachEnvironment),
		clothingTypes:  make(map[string]ClothingKnowledge),
	}
	
	kb.initializeYogaPoses()
	kb.initializeSwimmingStyles()
	kb.initializeGymnasticsMoves()
	kb.initializeBeachEnvironments()
	kb.initializeClothingKnowledge()
	
	return kb
}

func (s *SportsKnowledgeBase) initializeYogaPoses() {
	// 7+ Yoga poses with different angles and variations
	s.yogaPoses["downward_dog"] = YogaPose{
		Name:       "Downward-Facing Dog",
		Difficulty: "beginner",
		Angles:     []string{"side", "three_quarter", "front_low_angle"},
		Modifications: []string{
			"bent knees for beginners",
			"heels lifted for tight hamstrings",
		},
		SafetyNotes: "Keep spine straight, shoulders away from ears",
	}
	
	s.yogaPoses["warrior_ii"] = YogaPose{
		Name:       "Warrior II",
		Difficulty: "beginner",
		Angles:     []string{"front", "side", "low_angle", "high_angle"},
		Modifications: []string{
			"shorter stance for stability",
			"arms parallel to ground",
		},
		SafetyNotes: "Front knee aligned over ankle, back leg straight",
	}
	
	s.yogaPoses["tree_pose"] = YogaPose{
		Name:       "Tree Pose",
		Difficulty: "beginner",
		Angles:     []string{"front", "side", "three_quarter"},
		Modifications: []string{
			"foot on calf instead of thigh",
			"using wall for balance",
		},
		SafetyNotes: "Avoid placing foot on knee joint",
	}
	
	s.yogaPoses["cobra_pose"] = YogaPose{
		Name:       "Cobra Pose",
		Difficulty: "beginner", 
		Angles:     []string{"side", "three_quarter", "high_angle"},
		Modifications: []string{
			"low cobra with elbows bent",
			"gentle backbend",
		},
		SafetyNotes: "Keep pubic bone on floor, shoulders relaxed",
	}
	
	s.yogaPoses["bridge_pose"] = YogaPose{
		Name:       "Bridge Pose",
		Difficulty: "intermediate",
		Angles:     []string{"side", "low_angle", "three_quarter_low"},
		Modifications: []string{
			"supported bridge with block under sacrum",
			"one-legged variation",
		},
		SafetyNotes: "Keep knees hip-width apart, neck relaxed",
	}
	
	s.yogaPoses["dancers_pose"] = YogaPose{
		Name:       "Dancer's Pose",
		Difficulty: "intermediate",
		Angles:     []string{"front", "side", "three_quarter_dynamic"},
		Modifications: []string{
			"using strap for foot hold",
			"standing near wall for balance",
		},
		SafetyNotes: "Maintain standing leg stability, open chest",
	}
	
	s.yogaPoses["lotus_pose"] = YogaPose{
		Name:       "Lotus Pose",
		Difficulty: "advanced",
		Angles:     []string{"front", "high_angle", "eye_level"},
		Modifications: []string{
			"half lotus with one foot up",
			"easy pose cross-legged",
		},
		SafetyNotes: "Only if hips are open, avoid knee strain",
	}
}

func (s *SportsKnowledgeBase) initializeSwimmingStyles() {
	s.swimmingStyles["freestyle"] = SwimmingStyle{
		Name:        "Freestyle/Front Crawl",
		Environment: "pool",
		Physics:     "streamlined body position, alternating arm strokes, flutter kick",
		Clothing:    "competitive swimsuit, swim cap, goggles",
	}
	
	s.swimmingStyles["breaststroke"] = SwimmingStyle{
		Name:        "Breaststroke", 
		Environment: "pool",
		Physics:     "frog-like kick, simultaneous arm movement, glide phase",
		Clothing:    "training swimsuit, goggles",
	}
	
	s.swimmingStyles["backstroke"] = SwimmingStyle{
		Name:        "Backstroke",
		Environment: "pool",
		Physics:     "supine position, alternating arms, flutter kick",
		Clothing:    "racerback swimsuit, swim cap",
	}
	
	s.swimmingStyles["ocean_swimming"] = SwimmingStyle{
		Name:        "Ocean Swimming",
		Environment: "ocean",
		Physics:     "wave adaptation, current awareness, sighting technique",
		Clothing:    "bright colored swimsuit, safety buoy",
	}
	
	s.swimmingStyles["sea_swimming"] = SwimmingStyle{
		Name:        "Sea Swimming", 
		Environment: "sea",
		Physics:     "tidal awareness, temperature adaptation, marine life caution",
		Clothing:    "wetsuit for cold water, neoprene cap",
	}
}

func (s *SportsKnowledgeBase) initializeClothingKnowledge() {
	// Bikini vs Swimsuits
	s.clothingTypes["bikini"] = ClothingKnowledge{
		Type:        "bikini",
		Description: "Two-piece swimwear, various styles and coverage levels",
		Variations:  []string{"triangle", "bandeau", "high-waisted", "sport"},
		Physics:     "separate top and bottom, adjustable coverage",
	}
	
	s.clothingTypes["one_piece"] = ClothingKnowledge{
		Type:        "one_piece_swimsuit",
		Description: "Single-piece swimwear, athletic or fashion styles",
		Variations:  []string{"racerback", "monokini", "competitive", "plunge"},
		Physics:     "unified coverage, streamlined for swimming",
	}
	
	// Skirt and top combinations
	s.clothingTypes["mini_skirt"] = ClothingKnowledge{
		Type:        "mini_skirt",
		Description: "Short skirt with various lengths and styles",
		Physics:     "fabric drape, movement dynamics, coverage considerations",
		ActivityNotes: map[string]string{
			"yoga": "consider longer shorts or leggings for inversions",
			"gymnastics": "sports shorts recommended for dynamic movements",
			"beach": "appropriate for casual wear, consider wind",
		},
	}
	
	s.clothingTypes["off_shoulder"] = ClothingKnowledge{
		Type:        "off_shoulder_top",
		Description: "Top with shoulders exposed, various sleeve styles",
		Physics:     "may restrict arm movement, stays in place with elastic",
		ActivityNotes: map[string]string{
			"yoga": "may slip during arm balances",
			"swimming": "not suitable for water activities",
			"beach": "fashionable for casual wear",
		},
	}
}

// ValidateSportsKnowledge checks prompt against sports knowledge
func (s *SportsKnowledgeBase) ValidateSportsKnowledge(prompt string) *KnowledgeValidation {
	validation := &KnowledgeValidation{
		IsAccurate: true,
		Issues:     []string{},
	}
	
	// Check yoga pose accuracy
	if s.containsYogaPose(prompt) {
		if !s.isYogaPoseAccurate(prompt) {
			validation.IsAccurate = false
			validation.Issues = append(validation.Issues, "Yoga pose description may be inaccurate")
		}
	}
	
	// Check swimming environment accuracy
	if s.containsSwimming(prompt) {
		if !s.isSwimmingEnvironmentAccurate(prompt) {
			validation.IsAccurate = false
			validation.Issues = append(validation.Issues, "Swimming environment physics may be inaccurate")
		}
	}
	
	// Check clothing appropriateness
	if !s.isClothingAppropriate(prompt) {
		validation.IsAccurate = false
		validation.Issues = append(validation.Issues, "Clothing may not be appropriate for activity")
	}
	
	return validation
}
