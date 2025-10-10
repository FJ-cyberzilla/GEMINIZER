package ai

type ExpertCommandParser struct {
	photographyTerms map[string]string
	physicsTerms     map[string]string
	materialTerms    map[string]string
	styleTerms       map[string]string
}

func NewExpertCommandParser() *ExpertCommandParser {
	return &ExpertCommandParser{
		photographyTerms: map[string]string{
			"chiaroscuro":     "dramatic_high_contrast",
			"bokeh":           "shallow_depth_of_field", 
			"vignette":        "edge_darkening",
			"golden_hour":     "warm_natural_lighting",
			"rim_light":       "back_lighting_emphasis",
			"softbox":         "diffused_studio_light",
		},
		physicsTerms: map[string]string{
			"soaked":          "high_saturation_wetness",
			"damp":            "medium_saturation_wetness", 
			"underwater":      "full_immersion_physics",
			"flowing":         "fluid_dynamics_active",
			"clinging":        "wet_fabric_adhesion",
			"billowing":       "wind_affected_drape",
		},
		materialTerms: map[string]string{
			"silk":            "smooth_high_sheen",
			"cotton":          "matte_textured",
			"denim":           "heavy_structured", 
			"leather":         "supple_reflective",
			"chiffon":         "light_transparent",
			"satin":           "lustrous_smooth",
		},
	}
}

// ParseExpertCommand understands professional photography and physics commands
func (e *ExpertCommandParser) ParseExpertCommand(command string) *ExpertCommand {
	parsed := &ExpertCommand{
		Original: command,
		Elements: make(map[string]interface{}),
	}
	
	// Extract photography directives
	parsed.Photography = e.extractPhotographyDirectives(command)
	
	// Extract physics directives  
	parsed.Physics = e.extractPhysicsDirectives(command)
	
	// Extract material specifications
	parsed.Materials = e.extractMaterialSpecifications(command)
	
	// Extract style and mood
	parsed.Style = e.extractStyleDirectives(command)
	
	// Calculate command complexity
	parsed.Complexity = e.calculateComplexity(parsed)
	
	return parsed
}

func (e *ExpertCommandParser) extractPhysicsDirectives(command string) PhysicsDirectives {
	directives := PhysicsDirectives{}
	
	// Wetness detection
	if strings.Contains(command, "soaked") || strings.Contains(command, "dripping") {
		directives.WetnessLevel = 0.9
		directives.HasFluidPhysics = true
	} else if strings.Contains(command, "wet") || strings.Contains(command, "damp") {
		directives.WetnessLevel = 0.6
		directives.HasFluidPhysics = true
	}
	
	// Underwater detection
	if strings.Contains(command, "underwater") || strings.Contains(command, "submerged") {
		directives.IsUnderwater = true
		directives.HasFluidPhysics = true
		directives.HasBuoyancy = true
	}
	
	// Material physics
	if strings.Contains(command, "flowing") || strings.Contains(command, "billowing") {
		directives.HasClothPhysics = true
		directives.WindInfluence = 0.7
	}
	
	if strings.Contains(command, "clinging") || strings.Contains(command, "tight") {
		directives.HasClothPhysics = true
		directives.AdhesionLevel = 0.8
	}
	
	return directives
}

func (e *ExpertCommandParser) extractMaterialSpecifications(command string) []MaterialSpec {
	var specs []MaterialSpec
	
	// Detect material types
	for term, materialType := range e.materialTerms {
		if strings.Contains(strings.ToLower(command), term) {
			spec := MaterialSpec{
				Type: materialType,
				Properties: e.getMaterialProperties(materialType),
			}
			
			// Enhance with wetness if specified
			if strings.Contains(command, "wet") || strings.Contains(command, "soaked") {
				spec.Properties = e.applyWetnessToMaterial(spec.Properties)
			}
			
			specs = append(specs, spec)
		}
	}
	
	return specs
}
