import React, { useState, useEffect } from 'react';
import { Heart, Cpu, Zap, Activity, Users, Shield, Brain } from 'lucide-react';

const LivingAIDashboard = () => {
    const [systemStatus, setSystemStatus] = useState(null);
    const [agentActivity, setAgentActivity] = useState([]);

    useEffect(() => {
        fetchSystemStatus();
        const interval = setInterval(fetchSystemStatus, 10000); // Update every 10 seconds
        return () => clearInterval(interval);
    }, []);

    const fetchSystemStatus = async () => {
        try {
            const response = await fetch('/api/v1/ai/status');
            const status = await response.json();
            setSystemStatus(status);
            updateAgentActivity(status.agents);
        } catch (error) {
            console.error('Failed to fetch system status:', error);
        }
    };

    const updateAgentActivity = (agents) => {
        const activity = Object.entries(agents).map(([id, agent]) => ({
            id,
            name: agent.name,
            role: agent.role,
            health: agent.health,
            performance: agent.performance,
            lastActive: new Date(agent.lastActive),
            isActive: agent.isActive
        }));
        setAgentActivity(activity);
    };

    const getHealthColor = (score) => {
        if (score >= 0.8) return '#10b981'; // green
        if (score >= 0.6) return '#f59e0b'; // yellow
        return '#ef4444'; // red
    };

    const getPerformanceColor = (score) => {
        if (score >= 0.8) return '#8b5cf6'; // purple
        if (score >= 0.6) return '#3b82f6'; // blue
        return '#6b7280'; // gray
    };

    if (!systemStatus) {
        return <div className="loading">Connecting to AI System...</div>;
    }

    return (
        <div style={{
            background: 'rgba(255,255,255,0.03)',
            border: '1px solid rgba(255,255,255,0.1)',
            borderRadius: '16px',
            padding: '24px',
            marginBottom: '20px'
        }}>
            {/* System Overview */}
            <div style={{
                display: 'grid',
                gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))',
                gap: '16px',
                marginBottom: '24px'
            }}>
                <div style={{
                    background: 'rgba(16, 185, 129, 0.1)',
                    border: '1px solid rgba(16, 185, 129, 0.3)',
                    borderRadius: '12px',
                    padding: '16px',
                    textAlign: 'center'
                }}>
                    <Heart color="#10b981" size={24} />
                    <div style={{ fontSize: '12px', color: 'rgba(255,255,255,0.6)', marginTop: '8px' }}>
                        System Health
                    </div>
                    <div style={{ fontSize: '20px', fontWeight: 'bold', color: '#10b981' }}>
                        {(systemStatus.overallHealth * 100).toFixed(0)}%
                    </div>
                </div>

                <div style={{
                    background: 'rgba(139, 92, 246, 0.1)',
                    border: '1px solid rgba(139, 92, 246, 0.3)',
                    borderRadius: '12px',
                    padding: '16px',
                    textAlign: 'center'
                }}>
                    <Zap color="#8b5cf6" size={24} />
                    <div style={{ fontSize: '12px', color: 'rgba(255,255,255,0.6)', marginTop: '8px' }}>
                        Performance
                    </div>
                    <div style={{ fontSize: '20px', fontWeight: 'bold', color: '#8b5cf6' }}>
                        {(systemStatus.averagePerformance * 100).toFixed(0)}%
                    </div>
                </div>

                <div style={{
                    background: 'rgba(59, 130, 246, 0.1)',
                    border: '1px solid rgba(59, 130, 246, 0.3)',
                    borderRadius: '12px',
                    padding: '16px',
                    textAlign: 'center'
                }}>
                    <Users color="#3b82f6" size={24} />
                    <div style={{ fontSize: '12px', color: 'rgba(255,255,255,0.6)', marginTop: '8px' }}>
                        Active Agents
                    </div>
                    <div style={{ fontSize: '20px', fontWeight: 'bold', color: '#3b82f6' }}>
                        {systemStatus.agentsUsed}/{systemStatus.totalAgents}
                    </div>
                </div>

                <div style={{
                    background: 'rgba(245, 158, 11, 0.1)',
                    border: '1px solid rgba(245, 158, 11, 0.3)',
                    borderRadius: '12px',
                    padding: '16px',
                    textAlign: 'center'
                }}>
                    <Activity color="#f59e0b" size={24} />
                    <div style={{ fontSize: '12px', color: 'rgba(255,255,255,0.6)', marginTop: '8px' }}>
                        System Load
                    </div>
                    <div style={{ fontSize: '20px', fontWeight: 'bold', color: '#f59e0b' }}>
                        {(systemStatus.systemLoad * 100).toFixed(0)}%
                    </div>
                </div>
            </div>

            {/* Agent Activity */}
            <div>
                <h3 style={{ 
                    color: '#fff', 
                    marginBottom: '16px',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px'
                }}>
                    <Brain size={20} />
                    AI Agent Team
                </h3>
                
                <div style={{
                    display: 'grid',
                    gap: '12px'
                }}>
                    {agentActivity.map(agent => (
                        <div key={agent.id} style={{
                            background: 'rgba(255,255,255,0.05)',
                            border: `1px solid ${getHealthColor(agent.health)}30`,
                            borderRadius: '12px',
                            padding: '16px',
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'center'
                        }}>
                            <div>
                                <div style={{ 
                                    color: '#fff', 
                                    fontWeight: '600',
                                    display: 'flex',
                                    alignItems: 'center',
                                    gap: '8px'
                                }}>
                                    {agent.name}
                                    <div style={{
                                        width: '8px',
                                        height: '8px',
                                        borderRadius: '50%',
                                        background: agent.isActive ? '#10b981' : '#6b7280',
                                        animation: agent.isActive ? 'pulse 2s infinite' : 'none'
                                    }} />
                                </div>
                                <div style={{ 
                                    fontSize: '12px', 
                                    color: 'rgba(255,255,255,0.6)',
                                    marginTop: '4px'
                                }}>
                                    {agent.role}
                                </div>
                            </div>
                            
                            <div style={{
                                display: 'flex',
                                gap: '16px',
                                alignItems: 'center'
                            }}>
                                <div style={{ textAlign: 'center' }}>
                                    <div style={{ 
                                        fontSize: '11px', 
                                        color: 'rgba(255,255,255,0.6)'
                                    }}>
                                        Health
                                    </div>
                                    <div style={{ 
                                        fontSize: '14px', 
                                        fontWeight: 'bold',
                                        color: getHealthColor(agent.health)
                                    }}>
                                        {(agent.health * 100).toFixed(0)}%
                                    </div>
                                </div>
                                
                                <div style={{ textAlign: 'center' }}>
                                    <div style={{ 
                                        fontSize: '11px', 
                                        color: 'rgba(255,255,255,0.6)'
                                    }}>
                                        Performance
                                    </div>
                                    <div style={{ 
                                        fontSize: '14px', 
                                        fontWeight: 'bold',
                                        color: getPerformanceColor(agent.performance)
                                    }}>
                                        {(agent.performance * 100).toFixed(0)}%
                                    </div>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>

            {/* System Messages */}
            <div style={{
                marginTop: '20px',
                padding: '12px',
                background: 'rgba(59, 130, 246, 0.1)',
                border: '1px solid rgba(59, 130, 246, 0.3)',
                borderRadius: '8px',
                fontSize: '14px',
                color: '#3b82f6'
            }}>
                <Shield size={16} style={{ display: 'inline', marginRight: '8px' }} />
                AI System is operational and monitoring all agents. All systems nominal.
            </div>
        </div>
    );
};

export default LivingAIDashboard;
