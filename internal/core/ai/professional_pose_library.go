package ai

type ProfessionalPoseLibrary struct {
	standingPoses map[string]ProfessionalPose
	sittingPoses  map[string]ProfessionalPose
	fashionPoses  map[string]ProfessionalPose
	dynamicPoses  map[string]ProfessionalPose
}

type ProfessionalPose struct {
	Name           string
	Category       string
	Difficulty     string
	Description    string
	KeyPoints      []string
	GravityCenter  string
	Balance        string
	Angles         []string
	Lighting       string
	Background     string
	CommonUsage    []string
	PhysicsNotes   string
}

func NewProfessionalPoseLibrary() *ProfessionalPoseLibrary {
	lib := &ProfessionalPoseLibrary{
		standingPoses: make(map[string]ProfessionalPose),
		sittingPoses:  make(map[string]ProfessionalPose),
		fashionPoses:  make(map[string]ProfessionalPose),
		dynamicPoses:  make(map[string]ProfessionalPose),
	}
	
	lib.initializeStandingPoses()
	lib.initializeSittingPoses() 
	lib.initializeFashionPoses()
	lib.initializeDynamicPoses()
	
	return lib
}

func (p *ProfessionalPoseLibrary) initializeStandingPoses() {
	// 10+ Different standing poses
	p.standingPoses["contrapposto"] = ProfessionalPose{
		Name:        "Contrapposto",
		Category:    "standing",
		Difficulty:  "beginner",
		Description: "Classical weight shift pose with one hip higher",
		KeyPoints:   []string{"weight on one leg", "opposite hip raised", "shoulders counter-rotated", "relaxed stance"},
		GravityCenter: "over supporting leg",
		Balance:     "stable asymmetrical",
		Angles:      []string{"three_quarter", "front", "side"},
		Lighting:    "soft directional to emphasize curves",
		Background:  "minimalist studio, gradient backdrop",
		CommonUsage: []string{"fashion", "portrait", "art"},
		PhysicsNotes: "Natural weight distribution, relaxed muscle tension",
	}
	
	p.standingPoses["power_stance"] = ProfessionalPose{
		Name:        "Power Stance",
		Category:    "standing", 
		Difficulty:  "beginner",
		Description: "Confident wide stance with hands on hips",
		KeyPoints:   []string{"feet shoulder-width apart", "hands on hips", "chest forward", "chin slightly raised"},
		GravityCenter: "centered between feet",
		Balance:     "very stable",
		Angles:      []string{"low_angle", "front", "three_quarter_low"},
		Lighting:    "dramatic from below for empowerment",
		Background:  "urban environment, studio with hard light",
		CommonUsage: []string{"fitness", "leadership", "fashion"},
		PhysicsNotes: "Wide base of support, strong core engagement",
	}
	
	p.standingPoses["leaning_elegance"] = ProfessionalPose{
		Name:        "Leaning Elegance",
		Category:    "standing",
		Difficulty:  "intermediate", 
		Description: "Leaning against surface with elongated lines",
		KeyPoints:   []string{"one shoulder against wall", "legs crossed at ankle", "head tilted", "relaxed arms"},
		GravityCenter: "diagonal through body to wall",
		Balance:     "supported stable",
		Angles:      []string{"three_quarter", "side", "front_angled"},
		Lighting:    "soft side lighting for depth",
		Background:  "architectural elements, textured walls",
		CommonUsage: []string{"fashion", "editorial", "lifestyle"},
		PhysicsNotes: "Partial weight support, graceful lines",
	}
	
	p.standingPoses["dynamic_twist"] = ProfessionalPose{
		Name:        "Dynamic Twist",
		Category:    "standing",
		Difficulty:  "intermediate",
		Description: "Upper body twist with flowing movement",
		KeyPoints:   []string{"hips facing forward", "shoulders rotated", "arms in motion", "weight transition"},
		GravityCenter: "shifting during movement",
		Balance:     "dynamic stability",
		Angles:      []string{"three_quarter", "action_angle", "freeze_frame"},
		Lighting:    "motion-blur compatible, directional",
		Background:  "motion-friendly, simple backdrop",
		CommonUsage: []string{"dance", "sports", "action fashion"},
		PhysicsNotes: "Rotational momentum, core stabilization",
	}
	
	p.standingPoses["floating_serenity"] = ProfessionalPose{
		Name:        "Floating Serenity", 
		Category:    "standing",
		Difficulty:  "advanced",
		Description: "Weightless appearance with subtle balance",
		KeyPoints:   []string{"tiptoe preparation", "arms raised gently", "soft gaze upward", "minimal tension"},
		GravityCenter: "high and centered",
		Balance:     "precise and delicate",
		Angles:      []string{"low_angle", "eye_level", "high_angle"},
		Lighting:    "ethereal backlighting, soft glow",
		Background:  "dreamy gradient, soft focus",
		CommonUsage: []string{"artistic", "conceptual", "fine art"},
		PhysicsNotes: "Maximum balance challenge, core strength required",
	}
}

func (p *ProfessionalPoseLibrary) initializeSittingPoses() {
	p.sittingPoses["elegant_perch"] = ProfessionalPose{
		Name:        "Elegant Perch",
		Category:    "sitting",
		Difficulty:  "beginner",
		Description: "Sitting on edge with perfect posture",
		KeyPoints:   []string{"sitting on edge", "spine straight", "legs crossed at ankle", "hands resting lightly"},
		GravityCenter: "over sitting bones",
		Balance:     "stable seated",
		Angles:      []string{"three_quarter", "front", "side_elevated"},
		Lighting:    "soft top lighting for elegance",
		Background:  "minimalist stool, clean backdrop",
		CommonUsage: []string{"professional", "fashion", "portrait"},
		PhysicsNotes: "Proper spinal alignment, even weight distribution",
	}
	
	p.sittingPoses["casual_lounge"] = ProfessionalPose{
		Name:        "Casual Lounge",
		Category:    "sitting",
		Difficulty:  "beginner",
		Description: "Relaxed sitting with natural slouch",
		KeyPoints:   []string{"leaning back", "one arm resting", "legs slightly apart", "comfortable slouch"},
		GravityCenter: "back in chair",
		Balance:     "fully supported",
		Angles:      []string{"three_quarter", "front_relaxed", "side"},
		Lighting:    "natural window light, soft shadows",
		Background:  "comfortable chair, casual setting",
		CommonUsage: []string{"lifestyle", "casual fashion", "environmental portrait"},
		PhysicsNotes: "Relaxed muscle tone, natural gravity effects",
	}
	
	p.sittingPoses["thinker_pose"] = ProfessionalPose{
		Name:        "Thinker Pose",
		Category:    "sitting",
		Difficulty:  "intermediate",
		Description: "Classic contemplative pose with hand on chin",
		KeyPoints:   []string{"elbow on knee", "hand supporting chin", "leaning forward", "intense gaze"},
		GravityCenter: "forward over knees",
		Balance:     "active seated",
		Angles:      []string{"three_quarter_low", "front_engaged", "side_contemplative"},
		Lighting:    "dramatic side lighting for mood",
		Background:  "study environment, focused setting",
		CommonUsage: []string{"conceptual", "editorial", "dramatic portrait"},
		PhysicsNotes: "Forward weight shift, engaged posture",
	}
	
	p.sittingPoses["yoga_seated"] = ProfessionalPose{
		Name:        "Yoga Seated",
		Category:    "sitting",
		Difficulty:  "intermediate",
		Description: "Floor sitting with meditation posture",
		KeyPoints:   []string{"cross-legged", "spine elongated", "hands on knees", "peaceful expression"},
		GravityCenter: "centered over sitting bones",
		Balance:     "grounded stability",
		Angles:      []string{"front_eye_level", "high_angle", "three_quarter_ground"},
		Lighting:    "soft diffused, spiritual quality",
		Background:  "simple floor, meditation space",
		CommonUsage: []string{"wellness", "yoga", "spiritual"},
		PhysicsNotes: "Hip flexibility required, grounded energy",
	}
}

func (p *ProfessionalPoseLibrary) initializeFashionPoses() {
	// 10+ Fashion-specific poses
	p.fashionPoses["catwalk_pause"] = ProfessionalPose{
		Name:        "Catwalk Pause",
		Category:    "fashion",
		Difficulty:  "intermediate",
		Description: "Fashion model runway pause with attitude",
		KeyPoints:   []string{"one foot forward", "hip thrust", "hands on hips", "direct eye contact", "slight pout"},
		GravityCenter: "over back leg with forward momentum",
		Balance:     "dynamic pause",
		Angles:      []string{"three_quarter_runway", "front_dramatic", "low_angle_power"},
		Lighting:    "dramatic runway lighting, strong shadows",
		Background:  "minimalist runway, geometric backdrop",
		CommonUsage: []string{"high fashion", "runway", "editorial"},
		PhysicsNotes: "Forward energy, confident posture",
	}
	
	p.fashionPoses["over_shoulder"] = ProfessionalPose{
		Name:        "Over Shoulder",
		Category:    "fashion", 
		Difficulty:  "beginner",
		Description: "Looking back over shoulder with mysterious gaze",
		KeyPoints:   []string{"body facing away", "head turned back", "shoulders angled", "mysterious expression"},
		GravityCenter: "slightly back-weighted",
		Balance:     "stable with twist",
		Angles:      []string{"three_quarter_back", "side_profile", "dramatic_angle"},
		Lighting:    "backlight for silhouette, face in shadow",
		Background:  "moody backdrop, textured surfaces",
		CommonUsage: []string{"mysterious fashion", "lingerie", "artistic"},
		PhysicsNotes: "Spinal twist, neck rotation",
	}
	
	p.fashionPoses["fabric_flow"] = ProfessionalPose{
		Name:        "Fabric Flow",
		Category:    "fashion",
		Difficulty:  "advanced",
		Description: "Pose emphasizing fabric movement and drape",
		KeyPoints:   []string{"arms raised", "fabric billowing", "body twist", "flowing lines", "wind effect"},
		GravityCenter: "dynamic and shifting",
		Balance:     "movement-based",
		Angles:      []string{"action_freeze", "three_quarter_dynamic", "side_movement"},
		Lighting:    "freeze motion lighting, highlight fabric",
		Background:  "wind machine setup, clean studio",
		CommonUsage: []string{"fabric showcase", "dramatic fashion", "editorial"},
		PhysicsNotes: "Wind physics, fabric dynamics, body in motion",
	}
	
	p.fashionPoses["arch_back"] = ProfessionalPose{
		Name:        "Arch Back",
		Category:    "fashion",
		Difficulty:  "advanced",
		Description: "Dramatic back arch emphasizing silhouette",
		KeyPoints:   []string{"spine arched", "head back", "arms extended", "chest forward", "graceful curve"},
		GravityCenter: "shifted backward",
		Balance:     "controlled instability",
		Angles:      []string{"low_angle_dramatic", "side_silhouette", "three_quarter_arched"},
		Lighting:    "rim lighting for silhouette, dramatic shadows",
		Background:  "minimalist for silhouette focus",
		CommonUsage: []string{"high fashion", "artistic", "dramatic editorial"},
		PhysicsNotes: "Spinal flexibility, core strength, balance challenge",
	}
	
	p.fashionPoses["casual_lean"] = ProfessionalPose{
		Name:        "Casual Lean",
		Category:    "fashion",
		Difficulty:  "beginner",
		Description: "Relaxed lean against wall with casual confidence",
		KeyPoints:   []string{"shoulder against wall", "one leg bent", "relaxed arms", "natural smile", "casual gaze"},
		GravityCenter: "diagonal support",
		Balance:     "supported relaxed",
		Angles:      []string{"three_quarter_casual", "front_engaged", "side_relaxed"},
		Lighting:    "natural casual light, soft shadows",
		Background:  "urban wall, textured background",
		CommonUsage: []string{"streetwear", "casual fashion", "lifestyle"},
		PhysicsNotes: "Partial weight support, relaxed muscle tone",
	}
	
	// Additional fashion poses...
	p.fashionPoses["hand_on_hip"] = ProfessionalPose{
		Name:        "Hand on Hip",
		Category:    "fashion",
		Difficulty:  "beginner",
		Description: "Classic fashion pose with hand on hip for shape definition",
		KeyPoints:   []string{"hand on hip", "weight shift", "shoulders back", "confident stance"},
		GravityCenter: "slightly asymmetrical",
		Balance:     "stable fashion pose",
		Angles:      []string{"front", "three_quarter", "side_definition"},
		Lighting:    "shape-defining lighting",
		Background:  "clean studio backdrop",
		CommonUsage: []string{"catalog", "commercial", "fashion"},
		PhysicsNotes: "Creates hourglass silhouette, balanced asymmetry",
	}
	
	p.fashionPoses["walking_pose"] = ProfessionalPose{
		Name:        "Walking Pose",
		Category:    "fashion",
		Difficulty:  "intermediate", 
		Description: "Frozen walking motion showing garment movement",
		KeyPoints:   []string{"mid-stride", "arms swinging", "hair movement", "garment flow", "natural gait"},
		GravityCenter: "forward momentum",
		Balance:     "dynamic walking",
		Angles:      []string{"three_quarter_action", "side_motion", "front_approach"},
		Lighting:    "motion-friendly lighting",
		Background:  "simple for focus on motion",
		CommonUsage: []string{"garment showcase", "lifestyle", "action fashion"},
		PhysicsNotes: "Walking physics, momentum capture",
	}
}

// GetPoseRecommendation suggests appropriate poses based on context
func (p *ProfessionalPoseLibrary) GetPoseRecommendation(context PoseContext) []ProfessionalPose {
	var recommendations []ProfessionalPose
	
	// Filter poses based on context
	allPoses := p.getAllPoses()
	
	for _, pose := range allPoses {
		if p.matchesContext(pose, context) {
			recommendations = append(recommendations, pose)
		}
	}
	
	// Sort by relevance
	sort.Slice(recommendations, func(i, j int) bool {
		return p.calculateRelevance(recommendations[i], context) > 
			   p.calculateRelevance(recommendations[j], context)
	})
	
	return recommendations
}

func (p *ProfessionalPoseLibrary) matchesContext(pose ProfessionalPose, context PoseContext) bool {
	// Check category match
	if context.Category != "" && pose.Category != context.Category {
		return false
	}
	
	// Check difficulty match
	if context.Difficulty != "" && pose.Difficulty != context.Difficulty {
		return false
	}
	
	// Check usage context
	for _, usage := range context.Usage {
		for _, poseUsage := range pose.CommonUsage {
			if strings.Contains(strings.ToLower(poseUsage), strings.ToLower(usage)) {
				return true
			}
		}
	}
	
	return len(context.Usage) == 0
}
