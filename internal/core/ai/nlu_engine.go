package ai

import (
	"regexp"
	"strings"
)

type NLUEngine struct {
	intentRecognizer *IntentRecognizer
	entityExtractor  *EntityExtractor
	sentimentAnalyzer *SentimentAnalyzer
}

func NewNLUEngine() *NLUEngine {
	return &NLUEngine{
		intentRecognizer: NewIntentRecognizer(),
		entityExtractor:  NewEntityExtractor(),
		sentimentAnalyzer: NewSentimentAnalyzer(),
	}
}

// UnderstandPrompt extracts intent, entities, and sentiment from user input
func (n *NLUEngine) UnderstandPrompt(prompt string) *PromptUnderstanding {
	understanding := &PromptUnderstanding{
		RawPrompt: prompt,
	}
	
	understanding.Intent = n.intentRecognizer.RecognizeIntent(prompt)
	understanding.Entities = n.entityExtractor.ExtractEntities(prompt)
	understanding.Sentiment = n.sentimentAnalyzer.AnalyzeSentiment(prompt)
	understanding.Complexity = n.analyzeComplexity(prompt)
	
	return understanding
}

type IntentRecognizer struct {
	intentPatterns map[string]*regexp.Regexp
}

func NewIntentRecognizer() *IntentRecognizer {
	return &IntentRecognizer{
		intentPatterns: map[string]*regexp.Regexp{
			"portrait":    regexp.MustCompile(`(?i)(portrait|face|person|model)`),
			"fashion":     regexp.MustCompile(`(?i)(fashion|clothing|outfit|dress|wear)`),
			"product":     regexp.MustCompile(`(?i)(product|object|item|thing)`),
			"landscape":   regexp.MustCompile(`(?i)(landscape|scene|view|nature|outdoor)`),
			"conceptual":  regexp.MustCompile(`(?i)(concept|idea|abstract|surreal|fantasy)`),
		},
	}
}

func (i *IntentRecognizer) RecognizeIntent(prompt string) string {
	intentScores := make(map[string]int)
	
	for intent, pattern := range i.intentPatterns {
		if pattern.MatchString(prompt) {
			matches := pattern.FindAllString(prompt, -1)
			intentScores[intent] += len(matches)
		}
	}
	
	// Return intent with highest score
	maxScore := 0
	dominantIntent := "general"
	for intent, score := range intentScores {
		if score > maxScore {
			maxScore = score
			dominantIntent = intent
		}
	}
	
	return dominantIntent
}

type EntityExtractor struct {
	entityPatterns map[string]*regexp.Regexp
}

func NewEntityExtractor() *EntityExtractor {
	return &EntityExtractor{
		entityPatterns: map[string]*regexp.Regexp{
			"material":   regexp.MustCompile(`(?i)(cotton|silk|denim|leather|wool|fabric|material)`),
			"lighting":   regexp.MustCompile(`(?i)(light|sun|bright|dark|shadow|illumination)`),
			"emotion":    regexp.MustCompile(`(?i)(happy|sad|angry|confident|playful|serious|emotional)`),
			"color":      regexp.MustCompile(`(?i)(red|blue|green|yellow|black|white|color|colour)`),
			"composition": regexp.MustCompile(`(?i)(closeup|wide|angle|view|perspective|composition)`),
		},
	}
}

func (e *EntityExtractor) ExtractEntities(prompt string) []Entity {
	var entities []Entity
	
	for entityType, pattern := range e.entityPatterns {
		matches := pattern.FindAllString(prompt, -1)
		for _, match := range matches {
			entities = append(entities, Entity{
				Type:  entityType,
				Value: match,
				Start: strings.Index(prompt, match),
				End:   strings.Index(prompt, match) + len(match),
			})
		}
	}
	
	return entities
}
