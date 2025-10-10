import React, { useState } from 'react';
import { Play, Sparkles, Zap, CheckCircle, AlertTriangle } from 'lucide-react';

const LiveDemo = () => {
    const [demoRunning, setDemoRunning] = useState(false);
    const [demoResults, setDemoResults] = useState(null);
    const [agentActivity, setAgentActivity] = useState([]);

    const runLiveDemo = async () => {
        setDemoRunning(true);
        setDemoResults(null);
        setAgentActivity([]);

        // Simulate AI agent activity
        const agents = [
            { id: 'nlu', name: 'NLU Engine', role: 'Understanding your request...', delay: 1000 },
            { id: 'quality', name: 'Quality Inspector', role: 'Assessing prompt quality...', delay: 2000 },
            { id: 'curator', name: 'Prompt Curator', role: 'Enhancing with professional terms...', delay: 3000 },
            { id: 'error', name: 'Error Detective', role: 'Checking for issues...', delay: 4000 },
            { id: 'render', name: 'Render Engine', role: 'Creating professional image...', delay: 5000 },
        ];

        for (const agent of agents) {
            await new Promise(resolve => setTimeout(resolve, agent.delay));
            setAgentActivity(prev => [...prev, {
                ...agent,
                status: 'completed',
                timestamp: new Date(),
                icon: <CheckCircle size={16} color="#10b981" />
            }]);
        }

        // Get demo results from backend
        try {
            const response = await fetch('/api/v1/demo/sample');
            const results = await response.json();
            setDemoResults(results);
        } catch (error) {
            console.error('Demo failed:', error);
        }

        setDemoRunning(false);
    };

    return (
        <div style={{
            background: 'rgba(255,255,255,0.03)',
            border: '1px solid rgba(255,255,255,0.1)',
            borderRadius: '16px',
            padding: '24px',
            marginBottom: '20px'
        }}>
            <h2 style={{
                color: '#fff',
                marginBottom: '20px',
                display: 'flex',
                alignItems: 'center',
                gap: '10px'
            }}>
                <Zap size={24} />
                Live AI Demonstration
            </h2>

            {/* Sample Input Display */}
            <div style={{
                background: 'rgba(0,0,0,0.3)',
                border: '1px solid rgba(255,255,255,0.1)',
                borderRadius: '12px',
                padding: '16px',
                marginBottom: '20px'
            }}>
                <h3 style={{ color: '#fff', fontSize: '14px', marginBottom: '10px' }}>
                    Sample Input Being Processed:
                </h3>
                <div style={{
                    color: 'rgba(255,255,255,0.8)',
                    fontSize: '13px',
                    lineHeight: '1.5',
                    fontFamily: 'monospace',
                    background: 'rgba(0,0,0,0.5)',
                    padding: '12px',
                    borderRadius: '8px',
                    border: '1px solid rgba(255,255,255,0.05)'
                }}>
                    {`Character: Young woman with short wavy pink hair, blue eyes
Outfit: White bandeau and matching mini-skirt  
Pose: Full lotus yoga pose with Gyan Mudra
Physique: Athletic but feminine`}
                </div>
            </div>

            {/* Run Demo Button */}
            <button
                onClick={runLiveDemo}
                disabled={demoRunning}
                style={{
                    background: demoRunning 
                        ? 'rgba(139, 92, 246, 0.5)' 
                        : 'linear-gradient(135deg, #8b5cf6 0%, #3b82f6 100%)',
                    border: 'none',
                    borderRadius: '12px',
                    padding: '12px 24px',
                    color: '#fff',
                    fontSize: '14px',
                    fontWeight: '600',
                    cursor: demoRunning ? 'not-allowed' : 'pointer',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px',
                    marginBottom: '20px'
                }}
            >
                {demoRunning ? (
                    <>
                        <div style={{
                            width: '16px',
                            height: '16px',
                            border: '2px solid rgba(255,255,255,0.3)',
                            borderTop: '2px solid #fff',
                            borderRadius: '50%',
                            animation: 'spin 1s linear infinite'
                        }} />
                        AI Agents Working...
                    </>
                ) : (
                    <>
                        <Play size={16} />
                        Run Live AI Demo
                    </>
                )}
            </button>

            {/* Agent Activity Feed */}
            {agentActivity.length > 0 && (
                <div style={{
                    background: 'rgba(0,0,0,0.3)',
                    border: '1px solid rgba(255,255,255,0.1)',
                    borderRadius: '12px',
                    padding: '16px',
                    marginBottom: '20px'
                }}>
                    <h3 style={{ color: '#fff', fontSize: '14px', marginBottom: '12px' }}>
                        AI Agent Activity:
                    </h3>
                    <div style={{ display: 'flex', flexDirection: 'column', gap: '8px' }}>
                        {agentActivity.map((agent, index) => (
                            <div key={agent.id} style={{
                                display: 'flex',
                                alignItems: 'center',
                                gap: '10px',
                                padding: '8px',
                                background: 'rgba(255,255,255,0.05)',
                                borderRadius: '6px',
                                animation: 'fadeIn 0.5s ease-in'
                            }}>
                                {agent.icon}
                                <div style={{ flex: 1 }}>
                                    <div style={{ color: '#fff', fontSize: '12px', fontWeight: '600' }}>
                                        {agent.name}
                                    </div>
                                    <div style={{ color: 'rgba(255,255,255,0.6)', fontSize: '11px' }}>
                                        {agent.role}
                                    </div>
                                </div>
                                <div style={{
                                    fontSize: '10px',
                                    color: 'rgba(255,255,255,0.4)'
                                }}>
                                    {agent.timestamp.toLocaleTimeString()}
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}

            {/* Demo Results */}
            {demoResults && (
                <div style={{
                    background: 'rgba(16, 185, 129, 0.1)',
                    border: '1px solid rgba(16, 185, 129, 0.3)',
                    borderRadius: '12px',
                    padding: '16px'
                }}>
                    <h3 style={{ 
                        color: '#10b981',
                        fontSize: '14px',
                        marginBottom: '12px',
                        display: 'flex',
                        alignItems: 'center',
                        gap: '8px'
                    }}>
                        <Sparkles size={16} />
                        Demo Complete - AI Enhancement Results
                    </h3>
                    
                    <div style={{
                        display: 'grid',
                        gridTemplateColumns: '1fr 1fr',
                        gap: '16px'
                    }}>
                        {/* Original vs Enhanced */}
                        <div>
                            <h4 style={{ color: '#fff', fontSize: '12px', marginBottom: '8px' }}>
                                Original Input:
                            </h4>
                            <div style={{
                                background: 'rgba(0,0,0,0.5)',
                                padding: '12px',
                                borderRadius: '6px',
                                fontSize: '11px',
                                color: 'rgba(255,255,255,0.7)',
                                lineHeight: '1.4'
                            }}>
                                {demoResults.original_input}
                            </div>
                        </div>
                        
                        <div>
                            <h4 style={{ color: '#fff', fontSize: '12px', marginBottom: '8px' }}>
                                AI-Enhanced Professional Prompt:
                            </h4>
                            <div style={{
                                background: 'rgba(0,0,0,0.5)',
                                padding: '12px',
                                borderRadius: '6px',
                                fontSize: '11px',
                                color: '#10b981',
                                lineHeight: '1.4',
                                border: '1px solid rgba(16, 185, 129, 0.3)'
                            }}>
                                {demoResults.enhanced_prompt}
                            </div>
                        </div>
                    </div>

                    {/* Quality Score */}
                    <div style={{
                        marginTop: '12px',
                        padding: '8px',
                        background: 'rgba(255,255,255,0.05)',
                        borderRadius: '6px',
                        display: 'flex',
                        justifyContent: 'space-between',
                        alignItems: 'center'
                    }}>
                        <span style={{ color: '#fff', fontSize: '12px' }}>Quality Improvement:</span>
                        <span style={{ 
                            color: '#10b981', 
                            fontSize: '12px', 
                            fontWeight: 'bold' 
                        }}>
                            +{Math.round((demoResults.quality_score - 0.5) * 100)}%
                        </span>
                    </div>
                </div>
            )}

            <style>{`
                @keyframes spin {
                    to { transform: rotate(360deg); }
                }
                @keyframes fadeIn {
                    from { opacity: 0; transform: translateY(10px); }
                    to { opacity: 1; transform: translateY(0); }
                }
            `}</style>
        </div>
    );
};

export default LiveDemo;
