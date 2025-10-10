import React, { useState, useEffect } from 'react';
import { Shield, AlertTriangle, RefreshCw, ThumbsUp } from 'lucide-react';

const SafetyMonitor = ({ currentPrompt, onSafetyIssue, onRecoverySuggestion }) => {
    const [safetyStatus, setSafetyStatus] = useState('checking');
    const [issues, setIssues] = useState([]);
    const [recoverySuggestions, setRecoverySuggestions] = useState([]);

    useEffect(() => {
        checkSafety(currentPrompt);
    }, [currentPrompt]);

    const checkSafety = async (prompt) => {
        if (!prompt.trim()) {
            setSafetyStatus('idle');
            return;
        }

        try {
            const response = await fetch('/api/v1/safety/check', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ prompt: prompt })
            });

            const result = await response.json();
            
            setSafetyStatus(result.is_safe ? 'safe' : 'unsafe');
            setIssues(result.issues || []);
            setRecoverySuggestions(result.recovery_suggestions || []);

            if (!result.is_safe) {
                onSafetyIssue?.(result.issues);
            }

            if (result.recovery_suggestions) {
                onRecoverySuggestion?.(result.recovery_suggestions);
            }

        } catch (error) {
            setSafetyStatus('error');
            console.error('Safety check failed:', error);
        }
    };

    const getStatusColor = () => {
        switch (safetyStatus) {
            case 'safe': return '#10b981';
            case 'unsafe': return '#ef4444';
            case 'checking': return '#f59e0b';
            default: return '#6b7280';
        }
    };

    const getStatusIcon = () => {
        switch (safetyStatus) {
            case 'safe': return <ThumbsUp size={16} color="#10b981" />;
            case 'unsafe': return <AlertTriangle size={16} color="#ef4444" />;
            case 'checking': return <RefreshCw size={16} color="#f59e0b" />;
            default: return <Shield size={16} color="#6b7280" />;
        }
    };

    return (
        <div style={{
            background: 'rgba(255,255,255,0.03)',
            border: `1px solid ${getStatusColor()}30`,
            borderRadius: '12px',
            padding: '16px',
            marginBottom: '20px'
        }}>
            <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '10px',
                marginBottom: '12px'
            }}>
                {getStatusIcon()}
                <span style={{ color: '#fff', fontSize: '14px', fontWeight: '600' }}>
                    AI Safety Monitor
                </span>
                <div style={{
                    marginLeft: 'auto',
                    padding: '4px 8px',
                    background: getStatusColor() + '20',
                    color: getStatusColor(),
                    borderRadius: '6px',
                    fontSize: '11px',
                    fontWeight: '600'
                }}>
                    {safetyStatus.toUpperCase()}
                </div>
            </div>

            {/* Safety Issues */}
            {issues.length > 0 && (
                <div style={{ marginBottom: '12px' }}>
                    <div style={{ color: '#ef4444', fontSize: '12px', marginBottom: '8px' }}>
                        Safety Issues Detected:
                    </div>
                    <ul style={{
                        color: 'rgba(255,255,255,0.8)',
                        fontSize: '11px',
                        paddingLeft: '16px',
                        lineHeight: '1.4'
                    }}>
                        {issues.map((issue, index) => (
                            <li key={index}>{issue}</li>
                        ))}
                    </ul>
                </div>
            )}

            {/* Recovery Suggestions */}
            {recoverySuggestions.length > 0 && (
                <div>
                    <div style={{ color: '#10b981', fontSize: '12px', marginBottom: '8px' }}>
                        Suggested Alternatives:
                    </div>
                    <div style={{
                        background: 'rgba(16, 185, 129, 0.1)',
                        border: '1px solid rgba(16, 185, 129, 0.3)',
                        borderRadius: '8px',
                        padding: '12px'
                    }}>
                        {recoverySuggestions.map((suggestion, index) => (
                            <div key={index} style={{
                                color: '#10b981',
                                fontSize: '11px',
                                lineHeight: '1.4',
                                marginBottom: index < recoverySuggestions.length - 1 ? '8px' : '0',
                                paddingBottom: index < recoverySuggestions.length - 1 ? '8px' : '0',
                                borderBottom: index < recoverySuggestions.length - 1 ? '1px solid rgba(16, 185, 129, 0.2)' : 'none'
                            }}>
                                {suggestion}
                            </div>
                        ))}
                    </div>
                </div>
            )}

            {/* Safe Prompt Message */}
            {safetyStatus === 'safe' && issues.length === 0 && (
                <div style={{
                    color: '#10b981',
                    fontSize: '11px',
                    textAlign: 'center'
                }}>
                    ✓ Prompt meets all safety guidelines
                </div>
            )}
        </div>
    );
};

export default SafetyMonitor;
