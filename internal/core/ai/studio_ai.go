package ai

type StudioKnowledge struct {
	lightingSetups  map[string]LightingSetup
	backgroundTypes map[string]BackgroundType
	diffuseSystems  map[string]DiffuseSystem
	gradientTypes   map[string]GradientType
	studioProps     map[string]StudioProp
}

type LightingSetup struct {
	Name        string
	KeyLight    string
	FillLight   string
	RimLight    string
	BackgroundLight string
	Mood        string
	BestFor     []string
}

type DiffuseSystem struct {
	Name        string
	Type        string // softbox, umbrella, scrim, bounce
	Size        string
	Softness    string
	UseCase     []string
}

type GradientType struct {
	Name        string
	Colors      []string
	Direction   string
	Intensity   string
	Complexity  string
}

func NewStudioKnowledge() *StudioKnowledge {
	sk := &StudioKnowledge{
		lightingSetups:  make(map[string]LightingSetup),
		backgroundTypes: make(map[string]BackgroundType),
		diffuseSystems:  make(map[string]DiffuseSystem),
		gradientTypes:   make(map[string]GradientType),
		studioProps:     make(map[string]StudioProp),
	}
	
	sk.initializeLightingSetups()
	sk.initializeBackgroundTypes()
	sk.initializeDiffuseSystems()
	sk.initializeGradientTypes()
	sk.initializeStudioProps()
	
	return sk
}

func (s *StudioKnowledge) initializeLightingSetups() {
	s.lightingSetups["butterfly_lighting"] = LightingSetup{
		Name:        "Butterfly Lighting",
		KeyLight:    "directly above camera, high angle",
		FillLight:   "minimal or reflector below",
		RimLight:    "optional for separation",
		BackgroundLight: "even wash",
		Mood:        "glamorous, flattering",
		BestFor:     []string{"beauty", "glamour", "portrait"},
	}
	
	s.lightingSetups["rembrandt_lighting"] = LightingSetup{
		Name:        "Rembrandt Lighting",
		KeyLight:    "45 degrees to side, above eye level",
		FillLight:   "opposite side, subtle",
		RimLight:    "back for depth",
		BackgroundLight: "moody, darker",
		Mood:        "dramatic, artistic",
		BestFor:     []string{"dramatic portrait", "artistic", "character study"},
	}
	
	s.lightingSetups["clamshell_lighting"] = LightingSetup{
		Name:        "Clamshell Lighting",
		KeyLight:    "above subject",
		FillLight:   "directly below subject",
		RimLight:    "optional side light",
		BackgroundLight: "even and bright",
		Mood:        "commercial, clean",
		BestFor:     []string{"beauty", "commercial", "product"},
	}
}

func (s *StudioKnowledge) initializeDiffuseSystems() {
	s.diffuseSystems["large_softbox"] = DiffuseSystem{
		Name:     "Large Softbox",
		Type:     "softbox",
		Size:     "large (4x6 feet)",
		Softness: "very soft, wrap-around",
		UseCase:  []string{"portrait", "beauty", "full body"},
	}
	
	s.diffuseSystems["octabox"] = DiffuseSystem{
		Name:     "Octabox",
		Type:     "softbox", 
		Size:     "medium (3-4 feet)",
		Softness: "soft with circular catchlights",
		UseCase:  []string{"portrait", "fashion", "headshots"},
	}
	
	s.diffuseSystems["scrim"] = DiffuseSystem{
		Name:     "Scrim",
		Type:     "scrim",
		Size:     "large (6x6 feet+)",
		Softness: "diffused natural light quality",
		UseCase:  []string{"natural look", "large scenes", "window light simulation"},
	}
}

// GetProfessionalStudioSetup suggests complete studio configuration
func (s *StudioKnowledge) GetProfessionalStudioSetup(shotType string, mood string) *StudioSetup {
	setup := &StudioSetup{
		ShotType: shotType,
		Mood:     mood,
	}
	
	// Select lighting
	setup.Lighting = s.selectLighting(shotType, mood)
	
	// Select background
	setup.Background = s.selectBackground(shotType, mood)
	
	// Select diffusion
	setup.Diffusion = s.selectDiffusion(shotType, setup.Lighting)
	
	// Select props if needed
	setup.Props = s.selectProps(shotType, mood)
	
	return setup
}
