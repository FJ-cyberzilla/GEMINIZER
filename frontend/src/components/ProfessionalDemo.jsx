import React, { useState } from 'react';
import { Camera, Zap, Lightbulb, Settings, Award } from 'lucide-react';

const ProfessionalDemo = () => {
    const [selectedPose, setSelectedPose] = useState('power_walk_peek');
    const [analysis, setAnalysis] = useState(null);
    const [enhancedPrompt, setEnhancedPrompt] = useState('');

    const professionalPoses = {
        power_walk_peek: {
            name: "Power Walk Peek",
            description: "Dynamic walking pose emphasizing athleticism and shoulder definition",
            sample: `"Power Walk Peek shot to emphasize athleticism and defined shoulders. Model walking away then twisting back, low angle with hard directional side lighting."`
        },
        v_curve_lounge: {
            name: "V-Curve Lounge", 
            description: "Reclining pose creating elongated V-shape for flawless skin aesthetic",
            sample: `"V-Curve Lounge for idealized skin. Model reclining with body arch, high angle with broad flat lighting to eliminate shadows."`
        },
        flirty_twist: {
            name: "Flirty Twist",
            description: "Playful pose emphasizing S-curve and engaging connection", 
            sample: `"Flirty Twist emphasizing S-curve. Model with torso rotation and playful expression, eye-level angle with soft lighting."`
        }
    };

    const analyzeProfessionalPrompt = async () => {
        const samplePrompt = professionalPoses[selectedPose].sample;
        
        try {
            const response = await fetch('/api/v1/analyze/professional', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ prompt: samplePrompt })
            });
            
            const result = await response.json();
            setAnalysis(result.analysis);
            setEnhancedPrompt(result.enhanced_prompt);
        } catch (error) {
            console.error('Analysis failed:', error);
        }
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
                <Award size={24} />
                Professional Photography Intelligence
            </h2>

            {/* Pose Selection */}
            <div style={{ marginBottom: '20px' }}>
                <label style={{
                    display: 'block',
                    color: 'rgba(255,255,255,0.8)',
                    marginBottom: '10px',
                    fontSize: '14px'
                }}>
                    Select Professional Pose:
                </label>
                <select 
                    value={selectedPose}
                    onChange={(e) => setSelectedPose(e.target.value)}
                    style={{
                        width: '100%',
                        background: 'rgba(0,0,0,0.4)',
                        border: '1px solid rgba(255,255,255,0.1)',
                        borderRadius: '8px',
                        padding: '10px',
                        color: '#fff',
                        fontSize: '14px'
                    }}
                >
                    {Object.entries(professionalPoses).map(([key, pose]) => (
                        <option key={key} value={key}>
                            {pose.name}
                        </option>
                    ))}
                </select>
                <div style={{
                    color: 'rgba(255,255,255,0.6)',
                    fontSize: '12px',
                    marginTop: '8px'
                }}>
                    {professionalPoses[selectedPose].description}
                </div>
            </div>

            {/* Sample Prompt Display */}
            <div style={{
                background: 'rgba(0,0,0,0.3)',
                border: '1px solid rgba(255,255,255,0.1)',
                borderRadius: '12px',
                padding: '16px',
                marginBottom: '20px'
            }}>
                <div style={{
                    color: '#fff',
                    fontSize: '14px',
                    marginBottom: '10px',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px'
                }}>
                    <Camera size={16} />
                    Sample Professional Prompt:
                </div>
                <div style={{
                    color: 'rgba(255,255,255,0.8)',
                    fontSize: '13px',
                    lineHeight: '1.5',
                    fontFamily: 'monospace',
                    background: 'rgba(0,0,0,0.5)',
                    padding: '12px',
                    borderRadius: '8px'
                }}>
                    {professionalPoses[selectedPose].sample}
                </div>
            </div>

            {/* Analyze Button */}
            <button
                onClick={analyzeProfessionalPrompt}
                style={{
                    background: 'linear-gradient(135deg, #8b5cf6 0%, #3b82f6 100%)',
                    border: 'none',
                    borderRadius: '12px',
                    padding: '12px 24px',
                    color: '#fff',
                    fontSize: '14px',
                    fontWeight: '600',
                    cursor: 'pointer',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px',
                    marginBottom: '20px'
                }}
            >
                <Zap size={16} />
                Analyze & Enhance with AI
            </button>

            {/* Analysis Results */}
            {analysis && (
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
                        <Lightbulb size={16} />
                        AI Analysis Complete
                    </h3>

                    <div style={{
                        display: 'grid',
                        gridTemplateColumns: '1fr 1fr',
                        gap: '16px',
                        marginBottom: '16px'
                    }}>
                        {/* Quality Score */}
                        <div style={{
                            background: 'rgba(0,0,0,0.3)',
                            padding: '12px',
                            borderRadius: '8px',
                            textAlign: 'center'
                        }}>
                            <div style={{ color: 'rgba(255,255,255,0.6)', fontSize: '12px' }}>
                                Quality Score
                            </div>
                            <div style={{
                                color: '#10b981',
                                fontSize: '24px',
                                fontWeight: 'bold'
                            }}>
                                {Math.round(analysis.quality_score * 100)}%
                            </div>
                        </div>

                        {/* Detected Pose */}
                        <div style={{
                            background: 'rgba(0,0,0,0.3)',
                            padding: '12px',
                            borderRadius: '8px',
                            textAlign: 'center'
                        }}>
                            <div style={{ color: 'rgba(255,255,255,0.6)', fontSize: '12px' }}>
                                Detected Pose
                            </div>
                            <div style={{
                                color: '#3b82f6', 
                                fontSize: '14px',
                                fontWeight: 'bold'
                            }}>
                                {analysis.detected_pose || 'Custom Pose'}
                            </div>
                        </div>
                    </div>

                    {/* Enhanced Prompt */}
                    <div style={{ marginBottom: '16px' }}>
                        <div style={{
                            color: '#fff',
                            fontSize: '12px',
                            marginBottom: '8px',
                            display: 'flex',
                            alignItems: 'center',
                            gap: '8px'
                        }}>
                            <Settings size={14} />
                            AI-Enhanced Professional Prompt:
                        </div>
                        <div style={{
                            background: 'rgba(0,0,0,0.5)',
                            padding: '12px',
                            borderRadius: '8px',
                            fontSize: '12px',
                            color: '#10b981',
                            lineHeight: '1.4',
                            border: '1px solid rgba(16, 185, 129, 0.3)'
                        }}>
                            {enhancedPrompt}
                        </div>
                    </div>

                    {/* Improvements */}
                    {analysis.improvements && analysis.improvements.length > 0 && (
                        <div>
                            <div style={{
                                color: '#fff', 
                                fontSize: '12px',
                                marginBottom: '8px'
                            }}>
                                AI-Applied Improvements:
                            </div>
                            <ul style={{
                                color: 'rgba(255,255,255,0.8)',
                                fontSize: '11px',
                                paddingLeft: '16px',
                                lineHeight: '1.4'
                            }}>
                                {analysis.improvements.map((improvement, index) => (
                                    <li key={index}>{improvement}</li>
                                ))}
                            </ul>
                        </div>
                    )}
                </div>
            )}
        </div>
    );
};

export default ProfessionalDemo;
