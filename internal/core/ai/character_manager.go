package ai

type CharacterManager struct {
	consistencyDB *ConsistencyDatabase
	appearanceEngine *AppearanceEngine
	traitManager *TraitManager
}

func NewCharacterManager() *CharacterManager {
	return &CharacterManager{
		consistencyDB:    NewConsistencyDatabase(),
		appearanceEngine: NewAppearanceEngine(),
		traitManager:     NewTraitManager(),
	}
}

func (c *CharacterManager) DevelopCharacters(characterDescriptions []CharacterDescription) []ComicCharacter {
	var characters []ComicCharacter
	
	for _, desc := range characterDescriptions {
		character := c.createCharacter(desc)
		characters = append(characters, character)
		
		// Store in consistency database
		c.consistencyDB.StoreCharacter(character)
	}
	
	return characters
}

func (c *CharacterManager) createCharacter(desc CharacterDescription) ComicCharacter {
	character := ComicCharacter{
		Name:        desc.Name,
		Description: desc.Description,
		Personality: c.traitManager.DevelopPersonality(desc.PersonalityTraits),
		ConsistencyID: generateConsistencyID(desc.Name),
	}
	
	// Generate consistent appearance
	character.Appearance = c.appearanceEngine.GenerateAppearance(
		desc.AppearanceDescription,
		character.ConsistencyID,
	)
	
	return character
}

// EnsureConsistency verifies character remains consistent across panels
func (c *CharacterManager) EnsureConsistency(characterName string, panelDescription string) (*CharacterAppearance, error) {
	character, exists := c.consistencyDB.GetCharacter(characterName)
	if !exists {
		return nil, fmt.Errorf("character not found: %s", characterName)
	}
	
	// Extract appearance cues from panel description
	detectedAppearance := c.appearanceEngine.ExtractAppearanceFromDescription(panelDescription)
	
	// Verify consistency with stored character
	if err := c.consistencyDB.VerifyAppearanceConsistency(character.ConsistencyID, detectedAppearance); err != nil {
		// Auto-correct inconsistencies
		correctedAppearance := c.appearanceEngine.CorrectInconsistencies(
			character.Appearance, 
			detectedAppearance,
		)
		return &correctedAppearance, nil
	}
	
	return &character.Appearance, nil
}
