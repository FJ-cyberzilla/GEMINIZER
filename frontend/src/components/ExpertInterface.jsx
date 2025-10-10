import React, { useState, useRef } from 'react';
import { Command, Zap, Settings, Physics, Camera, Material } from 'lucide-react';

const ExpertInterface = ({ onExpertCommand, onPhysicsUpdate }) => {
    const [command, setCommand] = useState('');
    const [suggestions, setSuggestions] = useState([]);
    const [isExpertMode, setIsExpertMode] = useState(false);
    const commandHistory = useRef([]);

    const expertCommands = {
        photography: [
            "chiaroscuro lighting with high contrast",
            "shallow depth of field f/1.4",
            "golden hour warm natural lighting", 
            "rim light backlit silhouette",
            "vignette edge darkening",
            "softbox diffused studio light"
        ],
        physics: [
            "soaked fabric with water saturation",
            "underwater refraction and caustics",
            "flowing cloth with wind dynamics",
            "clinging wet material physics",
            "hair with water droplet simulation",
            "billowing fabric air resistance"
        ],
        materials: [
            "silk with subsurface scattering",
            "wet cotton transparency effect",
            "leather with specular highlights",
            "denim with heavy texture",
            "chiffon semi-transparent drape",
            "satin lustrous reflective surface"
        ]
    };

    const handleExpertCommand = async (cmd) => {
        const fullCommand = isExpertMode ? cmd : `EXPERT: ${cmd}`;
        
        try {
            const response = await fetch('/api/v1/expert/command', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ command: fullCommand })
            });

            const result = await response.json();
            
            // Add to history
            commandHistory.current.push({
                command: fullCommand,
                response: result.message,
                timestamp: new Date()
            });

            onExpertCommand?.(result);
            setSuggestions(result.suggestions || []);

        } catch (error) {
            console.error('Expert command failed:', error);
        }
    };

    const quickPhysicsCommand = (type, value) => {
        const commands = {
            wetness: `Apply ${value} wetness physics: fabric saturation, transparency, drape alteration`,
            underwater: "Engage underwater rendering: fluid dynamics, light refraction, particle system",
            flowing: "Enable cloth physics: wind influence, natural folds, movement dynamics",
            material: `Configure ${value} material: subsurface scattering, texture detail, reflectivity`
        };
        
        handleExpertCommand(commands[type]);
    };

    return (
        <div style={{
            background: 'rgba(255,255,255,0.03)',
            border: '1px solid rgba(255,255,255,0.1)',
            borderRadius: '16px',
            padding: '24px',
            marginBottom: '20px'
        }}>
            <div style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'space-between',
                marginBottom: '20px'
            }}>
                <h2 style={{
                    color: '#fff',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '10px'
                }}>
                    <Command size={24} />
                    Expert Command Interface
                </h2>
                
                <label style={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px',
                    color: 'rgba(255,255,255,0.8)',
                    fontSize: '14px',
                    cursor: 'pointer'
                }}>
                    <input
                        type="checkbox"
                        checked={isExpertMode}
                        onChange={(e) => setIsExpertMode(e.target.checked)}
                        style={{ margin: 0 }}
                    />
                    Expert Mode
                </label>
            </div>

            {/* Command Input */}
            <div style={{ marginBottom: '20px' }}>
                <div style={{
                    display: 'flex',
                    gap: '10px',
                    marginBottom: '10px'
                }}>
                    <input
                        type="text"
                        value={command}
                        onChange={(e) => setCommand(e.target.value)}
                        placeholder="Enter expert photography/physics command..."
                        style={{
                            flex: 1,
                            background: 'rgba(0,0,0,0.4)',
                            border: '1px solid rgba(255,255,255,0.1)',
                            borderRadius: '8px',
                            padding: '12px',
                            color: '#fff',
                            fontSize: '14px'
                        }}
                        onKeyPress={(e) => {
                            if (e.key === 'Enter') {
                                handleExpertCommand(command);
                                setCommand('');
                            }
                        }}
                    />
                    <button
                        onClick={() => {
                            handleExpertCommand(command);
                            setCommand('');
                        }}
                        style={{
                            background: 'linear-gradient(135deg, #8b5cf6 0%, #3b82f6 100%)',
                            border: 'none',
                            borderRadius: '8px',
                            padding: '12px 20px',
                            color: '#fff',
                            fontSize: '14px',
                            fontWeight: '600',
                            cursor: 'pointer',
                            display: 'flex',
                            alignItems: 'center',
                            gap: '8px'
                        }}
                    >
                        <Zap size={16} />
                        Execute
                    </button>
                </div>
                
                <div style={{
                    color: 'rgba(255,255,255,0.6)',
                    fontSize: '12px',
                    fontStyle: 'italic'
                }}>
                    {isExpertMode 
                        ? "Expert Mode: Direct physics and photography commands accepted"
                        : "Beginner Mode: Natural language commands"
                    }
                </div>
            </div>

            {/* Quick Physics Controls */}
            <div style={{
                background: 'rgba(0,0,0,0.3)',
                border: '1px solid rgba(255,255,255,0.1)',
                borderRadius: '12px',
                padding: '16px',
                marginBottom: '20px'
            }}>
                <h3 style={{
                    color: '#fff',
                    fontSize: '14px',
                    marginBottom: '12px',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px'
                }}>
                    <Physics size={16} />
                    Quick Physics Controls
                </h3>
                
                <div style={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(auto-fit, minmax(150px, 1fr))',
                    gap: '10px'
                }}>
                    <button
                        onClick={() => quickPhysicsCommand('wetness', 'light')}
                        style={{
                            background: 'rgba(59, 130, 246, 0.1)',
                            border: '1px solid rgba(59, 130, 246, 0.3)',
                            borderRadius: '6px',
                            padding: '8px 12px',
                            color: '#3b82f6',
                            fontSize: '12px',
                            cursor: 'pointer'
                        }}
                    >
                        💧 Light Wetness
                    </button>
                    
                    <button
                        onClick={() => quickPhysicsCommand('wetness', 'heavy')}
                        style={{
                            background: 'rgba(59, 130, 246, 0.2)',
                            border: '1px solid rgba(59, 130, 246, 0.4)',
                            borderRadius: '6px',
                            padding: '8px 12px',
                            color: '#3b82f6',
                            fontSize: '12px',
                            cursor: 'pointer'
                        }}
                    >
                        💦 Heavy Soaked
                    </button>
                    
                    <button
                        onClick={() => quickPhysicsCommand('underwater', '')}
                        style={{
                            background: 'rgba(16, 185, 129, 0.1)',
                            border: '1px solid rgba(16, 185, 129, 0.3)',
                            borderRadius: '6px',
                            padding: '8px 12px',
                            color: '#10b981',
                            fontSize: '12px',
                            cursor: 'pointer'
                        }}
                    >
                        🌊 Underwater
                    </button>
                    
                    <button
                        onClick={() => quickPhysicsCommand('flowing', '')}
                        style={{
                            background: 'rgba(168, 85, 247, 0.1)',
                            border: '1px solid rgba(168, 85, 247, 0.3)',
                            borderRadius: '6px',
                            padding: '8px 12px',
                            color: '#a855f7',
                            fontSize: '12px',
                            cursor: 'pointer'
                        }}
                    >
                        💨 Flowing Fabric
                    </button>
                </div>
            </div>

            {/* Command Suggestions */}
            <div>
                <h3 style={{
                    color: '#fff',
                    fontSize: '14px',
                    marginBottom: '12px',
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px'
                }}>
                    <Settings size={16} />
                    Expert Command Library
                </h3>
                
                <div style={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))',
                    gap: '12px'
                }}>
                    {/* Photography Commands */}
                    <div>
                        <div style={{
                            color: '#3b82f6',
                            fontSize: '12px',
                            marginBottom: '8px',
                            display: 'flex',
                            alignItems: 'center',
                            gap: '6px'
                        }}>
                            <Camera size={14} />
                            Photography
                        </div>
                        {expertCommands.photography.map((cmd, index) => (
                            <div
                                key={index}
                                onClick={() => handleExpertCommand(cmd)}
                                style={{
                                    background: 'rgba(255,255,255,0.05)',
                                    border: '1px solid rgba(255,255,255,0.1)',
                                    borderRadius: '6px',
                                    padding: '8px 12px',
                                    marginBottom: '6px',
                                    fontSize: '11px',
                                    color: 'rgba(255,255,255,0.8)',
                                    cursor: 'pointer',
                                    transition: 'all 0.2s ease'
                                }}
                                onMouseEnter={(e) => {
                                    e.target.style.background = 'rgba(59, 130, 246, 0.2)';
                                    e.target.style.borderColor = 'rgba(59, 130, 246, 0.4)';
                                }}
                                onMouseLeave={(e) => {
                                    e.target.style.background = 'rgba(255,255,255,0.05)';
                                    e.target.style.borderColor = 'rgba(255,255,255,0.1)';
                                }}
                            >
                                {cmd}
                            </div>
                        ))}
                    </div>
                    
                    {/* Physics Commands */}
                    <div>
                        <div style={{
                            color: '#10b981',
                            fontSize: '12px',
                            marginBottom: '8px',
                            display: 'flex',
                            alignItems: 'center',
                            gap: '6px'
                        }}>
                            <Physics size={14} />
                            Physics
                        </div>
                        {expertCommands.physics.map((cmd, index) => (
                            <div
                                key={index}
                                onClick={() => handleExpertCommand(cmd)}
                                style={{
                                    background: 'rgba(255,255,255,0.05)',
                                    border: '1px solid rgba(255,255,255,0.1)',
                                    borderRadius: '6px',
                                    padding: '8px 12px',
                                    marginBottom: '6px',
                                    fontSize: '11px',
                                    color: 'rgba(255,255,255,0.8)',
                                    cursor: 'pointer',
                                    transition: 'all 0.2s ease'
                                }}
                                onMouseEnter={(e) => {
                                    e.target.style.background = 'rgba(16, 185, 129, 0.2)';
                                    e.target.style.borderColor = 'rgba(16, 185, 129, 0.4)';
                                }}
                                onMouseLeave={(e) => {
                                    e.target.style.background = 'rgba(255,255,255,0.05)';
                                    e.target.style.borderColor = 'rgba(255,255,255,0.1)';
                                }}
                            >
                                {cmd}
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* Command History */}
            {commandHistory.current.length > 0 && (
                <div style={{ marginTop: '20px' }}>
                    <h3 style={{
                        color: '#fff',
                        fontSize: '14px',
                        marginBottom: '12px'
                    }}>
                        Command History
                    </h3>
                    <div style={{
                        maxHeight: '200px',
                        overflowY: 'auto',
                        background: 'rgba(0,0,0,0.3)',
                        borderRadius: '8px',
                        padding: '12px'
                    }}>
                        {commandHistory.current.slice(-5).map((item, index) => (
                            <div key={index} style={{
                                marginBottom: '8px',
                                padding: '8px',
                                background: 'rgba(255,255,255,0.05)',
                                borderRadius: '6px'
                            }}>
                                <div style={{
                                    color: '#3b82f6',
                                    fontSize: '11px',
                                    fontWeight: '600'
                                }}>
                                    {item.command}
                                </div>
                                <div style={{
                                    color: 'rgba(255,255,255,0.7)',
                                    fontSize: '10px',
                                    marginTop: '4px'
                                }}>
                                    {item.response}
                                </div>
                                <div style={{
                                    color: 'rgba(255,255,255,0.4)',
                                    fontSize: '9px',
                                    marginTop: '2px'
                                }}>
                                    {item.timestamp.toLocaleTimeString()}
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default ExpertInterface;
