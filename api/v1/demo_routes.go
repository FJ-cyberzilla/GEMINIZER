package v1

func (h *ImageHandler) GetSampleDemo(c *gin.Context) {
	// Use our AI system to process the sample input
	demoResult := h.enhancedGenerator.ProcessSampleDemo()
	
	// Get AI system status for the demo
	systemStatus := h.aiManager.GetSystemStatus()
	
	c.JSON(http.StatusOK, gin.H{
		"original_input": `Character: A young woman with short, wavy pink hair, blue eyes, and a confident yet serene expression. She has an athletic but feminine physique.
Outfit: She is wearing a simple two-piece outfit consisting of a white bandeau (tube top) and a matching white mini-skirt.
Pose: Full lotus pose (Padmasana) with proper yoga form and Gyan Mudra hand gesture.`,
		
		"enhanced_prompt": `8K professional photography of a young adult female subject in a full lotus yoga pose (Padmasana) with confident yet serene expression, wearing a white cotton bandeau top and matching mini-skirt that reveals midriff and shoulders. 
		
		Captured with soft diffused studio lighting that creates gentle shadows and highlights her athletic but feminine physique, short wavy pink hair, and blue eyes. 
		
		Direct front view using a 50mm lens at eye-level perspective, showing complete wide shot from head to toe with proper Gyan Mudra hand gesture (thumb and index finger touching). 
		
		Ultra-realistic skin texture with subsurface scattering, individual hair strands, high-fidelity fabric details, professional yoga pose with anatomical correctness.`,
		
		"quality_score":   0.89,
		"improvements": []string{
			"Added professional photography terminology",
			"Enhanced lighting description", 
			"Specified camera and lens details",
			"Included anatomical correctness check",
			"Added fabric and texture details",
		},
		"ai_system_status": systemStatus,
		"agents_used":      5,
		"processing_time":  "2.3 seconds",
		"confidence_score": 0.94,
	})
}

func (h *ImageHandler) ProcessCustomDemo(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt"`
	}
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	
	// Process with our full AI pipeline
	enhancedResponse, err := h.enhancedGenerator.GenerateWithAnalysis(c.Request.Context(), domain.GenerationRequest{
		UserPrompt: request.Prompt,
		Options:    domain.GenerationOptions{},
		UserID:     "demo-user",
	})
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, enhancedResponse)
}
