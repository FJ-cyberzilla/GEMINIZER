package ai

type ProfessionalPoseLibrary struct {
	poses map[string]ProfessionalPose
	styles map[string]PhotographyStyle
}

type ProfessionalPose struct {
	Name         string
	Category     string // standing, sitting, lying, dynamic
	Description  string
	ModelInstructions string
	PhotographerInstructions string
	KeyElements  []string
	Difficulty   string // beginner, intermediate, advanced
}

type PhotographyStyle struct {
	Name        string
	Lighting    string
	CameraSetup string
	Mood        string
	CommonUses  []string
}

func NewProfessionalPoseLibrary() *ProfessionalPoseLibrary {
	lib := &ProfessionalPoseLibrary{
		poses: make(map[string]ProfessionalPose),
		styles: make(map[string]PhotographyStyle),
	}
	
	lib.initializePoses()
	lib.initializeStyles()
	
	return lib
}

func (p *ProfessionalPoseLibrary) initializePoses() {
	// Sample 1: Power Walk Peek
	p.poses["power_walk_peek"] = ProfessionalPose{
		Name:        "Power Walk Peek",
		Category:    "standing",
		Description: "Dynamic walking pose emphasizing athleticism and shoulder definition",
		ModelInstructions: "Initiate deliberate stride away from camera, pause and twist sharply back, look over furthest shoulder, keep core engaged, arms slightly swinging",
		PhotographerInstructions: "Frame at low mid-level angle, use hard directional side lighting to accentuate muscle definition, freeze action moment",
		KeyElements: []string{
			"dynamic movement", "shoulder definition", "twisting motion", 
			"engaged core", "confident energy", "frozen action",
		},
		Difficulty: "intermediate",
	}
	
	// Sample 2: V-Curve Lounge
	p.poses["v_curve_lounge"] = ProfessionalPose{
		Name:        "V-Curve Lounge", 
		Category:    "lying",
		Description: "Reclining pose creating elongated V-shape for flawless skin aesthetic",
		ModelInstructions: "Recline on ground or prop, create full-body arch by pushing chest forward and extending limbs, focus on relaxation and clean lines with minimal tension",
		PhotographerInstructions: "Position camera at high angle (45 degrees down), use broad flat lighting with massive softbox to eliminate harsh shadows, ensure sleek skin texture and symmetrical body lines",
		KeyElements: []string{
			"body arch", "clean lines", "relaxed tension", "flawless skin", 
			"symmetrical composition", "shadow-free lighting",
		},
		Difficulty: "beginner",
	}
	
	// Additional professional poses
	p.poses["flirty_twist"] = ProfessionalPose{
		Name:        "Flirty Twist",
		Category:    "standing", 
		Description: "Playful pose emphasizing S-curve and engaging connection",
		ModelInstructions: "Create S-curve with torso rotation, one hand on hip or playing with hair, slight head tilt with engaging eye contact, playful smile",
		PhotographerInstructions: "Use eye-level angle with slight low angle for empowerment, soft butterfly lighting for flattering facial features, medium close-up framing",
		KeyElements: []string{"s-curve", "playful engagement", "eye contact", "torso rotation"},
		Difficulty: "beginner",
	}
}

func (p *ProfessionalPoseLibrary) initializeStyles() {
	p.styles["athletic_dynamic"] = PhotographyStyle{
		Name:        "Athletic Dynamic",
		Lighting:    "Hard directional side lighting with dramatic shadows",
		CameraSetup: "Low angle with wide lens for empowerment, fast shutter to freeze motion",
		Mood:        "Intense, powerful, energetic",
		CommonUses:  ["fitness photography", "sportswear", "action shots"],
	}
	
	p.styles["flawless_beauty"] = PhotographyStyle{
		Name:        "Flawless Beauty", 
		Lighting:    "Broad flat lighting with massive softbox to eliminate shadows",
		CameraSetup: "High angle for flattering perspective, medium telephoto lens",
		Mood:        "Soft, elegant, pristine",
		CommonUses:  ["beauty shots", "skin care", "lingerie", "portrait"],
	}
	
	p.styles["cinematic_dramatic"] = PhotographyStyle{
		Name:        "Cinematic Dramatic",
		Lighting:    "Chiaroscuro with high contrast and deep shadows", 
		CameraSetup: "Dutch angles, wide aperture for shallow depth of field",
		Mood:        "Moody, intense, artistic",
		CommonUses:  ["fashion editorial", "conceptual art", "dramatic portraits"],
	}
}

// AnalyzeProfessionalPrompt breaks down professional photography prompts
func (p *ProfessionalPoseLibrary) AnalyzeProfessionalPrompt(prompt string) *ProfessionalAnalysis {
	analysis := &ProfessionalAnalysis{
		OriginalPrompt: prompt,
		DetectedPose:   p.detectPose(prompt),
		DetectedStyle:  p.detectStyle(prompt),
		Elements:       p.extractPhotographyElements(prompt),
	}
	
	analysis.Confidence = p.calculateAnalysisConfidence(analysis)
	analysis.Suggestions = p.generateProfessionalSuggestions(analysis)
	
	return analysis
}

func (p *ProfessionalPoseLibrary) detectPose(prompt string) *ProfessionalPose {
	// Advanced pattern matching for professional poses
	for _, pose := range p.poses {
		if p.containsPoseElements(prompt, pose) {
			return &pose
		}
	}
	return nil
}

func (p *ProfessionalPoseLibrary) containsPoseElements(prompt string, pose ProfessionalPose) bool {
	keywords := append(pose.KeyElements, 
		strings.ToLower(pose.Name),
		strings.ToLower(pose.Category),
	)
	
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(prompt), strings.ToLower(keyword)) {
			return true
		}
	}
	return false
}
