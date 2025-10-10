package ai

// GenerateComicPanelWithPhysics creates panels with advanced material handling
func (c *ComicBookEngine) GenerateComicPanelWithPhysics(panelDesc string, context ComicContext) (*ComicPanel, error) {
	// Apply advanced physics to panel description
	enhancedDesc := c.materialPhysics.HandleAdvancedMaterials(panelDesc)
	
	// Safety check
	if err := c.safetyFilter.ValidatePanelDescription(enhancedDesc); err != nil {
		return nil, err
	}
	
	// Ensure character consistency
	for _, character := range context.Characters {
		consistentAppearance, err := c.characterManager.EnsureConsistency(
			character.Name, 
			enhancedDesc,
		)
		if err != nil {
			return nil, err
		}
		
		// Update description with consistent appearance
		enhancedDesc = c.injectCharacterAppearance(enhancedDesc, character.Name, *consistentAppearance)
	}
	
	panel := &ComicPanel{
		Description: enhancedDesc,
		Dialogue:    context.Dialogue,
		CameraAngle: c.detectCameraAngle(enhancedDesc),
		Lighting:    c.detectLighting(enhancedDesc),
		Emotion:     c.detectEmotion(enhancedDesc),
		PhysicsFlags: c.extractPhysicsFlags(enhancedDesc),
	}
	
	return panel, nil
}

// Handle specific comic panel with wetness
func (c *ComicBookEngine) handleWetScene(panelDesc string) string {
	// Example: Panel 17 - Warehouse scene could have wet ground from recent rain
	if strings.Contains(panelDesc, "warehouse") && strings.Contains(panelDesc, "concrete") {
		return panelDesc + ", wet concrete floor with water reflections from recent rain, atmospheric moisture in air"
	}
	
	// Example: Underwater or wet hair scenes
	if strings.Contains(panelDesc, "hair") && (strings.Contains(panelDesc, "wet") || strings.Contains(panelDesc, "water")) {
		return c.materialPhysics.applyWetnessEffects(panelDesc, 0.7)
	}
	
	return panelDesc
}
