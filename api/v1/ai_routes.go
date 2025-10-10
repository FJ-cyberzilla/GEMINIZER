package v1

// AI Management endpoints
func (h *ImageHandler) GetAIStatus(c *gin.Context) {
	status := h.aiManager.GetSystemStatus()
	
	c.JSON(http.StatusOK, gin.H{
		"status":              "operational",
		"overall_health":      status.OverallHealth,
		"average_performance": status.AveragePerformance,
		"system_load":         status.SystemLoad,
		"uptime":              status.Uptime.String(),
		"total_agents":        status.TotalAgents,
		"agents_used":         len(status.Agents),
		"agents":              status.Agents,
		"timestamp":           time.Now().UTC(),
	})
}

func (h *ImageHandler) GetAgentDetails(c *gin.Context) {
	agentID := c.Param("id")
	
	// Get detailed agent information
	agentInfo := h.aiManager.GetAgentDetails(agentID)
	
	c.JSON(http.StatusOK, gin.H{
		"agent":     agentInfo,
		"timestamp": time.Now().UTC(),
	})
}

func (h *ImageHandler) SystemDiagnostics(c *gin.Context) {
	diagnostics := h.healthMonitor.CheckSystemHealth()
	
	c.JSON(http.StatusOK, gin.H{
		"diagnostics": diagnostics,
		"timestamp":   time.Now().UTC(),
	})
}
