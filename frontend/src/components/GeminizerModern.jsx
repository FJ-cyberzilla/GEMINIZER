// frontend/src/components/GeminizerModern.jsx - UPDATED
import React, { useState } from 'react';
import { Sparkles, Palette, Sun, Smile, Layers, Zap, Image, Wand2, History, Download } from 'lucide-react';
import { useImageGeneration } from '../hooks/useImageGeneration';
import { useHistory } from '../hooks/useHistory';

export default function GeminizerModern() {
  const [activeMode, setActiveMode] = useState(0);
  const [prompt, setPrompt] = useState('');
  const [showHistory, setShowHistory] = useState(false);
  
  const { generateImage, generating, result, error, clearResult } = useImageGeneration();
  const { history, fetchHistory, deleteGeneration } = useHistory();

  // Your existing modes array remains the same...
  const modes = [ /* ... your existing modes ... */ ];

  const [selectedSettings, setSelectedSettings] = useState(
    modes.reduce((acc, _, idx) => ({ ...acc, [idx]: 0 }), {})
  );

  const handleGenerate = async () => {
    const options = {
      archetype: modes[0].settings[selectedSettings[0]].label.toLowerCase(),
      material: modes[1].settings[selectedSettings[1]].label.toLowerCase(),
      lighting: modes[2].settings[selectedSettings[2]].label.toLowerCase(),
      style: modes[3].settings[selectedSettings[3]].label.toLowerCase(),
      angle: modes[4].settings[selectedSettings[4]].label.toLowerCase(),
      quality: modes[5].settings[selectedSettings[5]].label.toLowerCase(),
    };

    try {
      await generateImage(prompt, options);
      // Refresh history to include new generation
      await fetchHistory();
    } catch (err) {
      // Error handled by hook
    }
  };

  const downloadImage = () => {
    if (result?.image_data) {
      const link = document.createElement('a');
      link.href = `data:image/jpeg;base64,${result.image_data}`;
      link.download = `geminizer-${result.request_id}.jpg`;
      link.click();
    }
  };

  return (
    <div style={{ /* your existing styles */ }}>
      {/* Your existing UI remains the same */}
      
      {/* Enhanced result display */}
      {result && (
        <div style={{
          marginTop: '20px',
          background: 'linear-gradient(135deg, rgba(255,20,147,0.1) 0%, rgba(0,255,255,0.1) 100%)',
          border: '1px solid rgba(255,20,147,0.3)',
          borderRadius: '16px',
          padding: '20px'
        }}>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            marginBottom: '12px',
          }}>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              gap: '10px',
              color: '#ff1493',
              fontSize: '12px',
              fontWeight: '600',
              textTransform: 'uppercase',
              letterSpacing: '0.1em'
            }}>
              <Wand2 size={16} />
              <span>Professional Generation Complete</span>
            </div>
            <button
              onClick={downloadImage}
              style={{
                background: 'rgba(255,255,255,0.1)',
                border: '1px solid rgba(255,255,255,0.2)',
                borderRadius: '8px',
                padding: '8px 16px',
                color: '#fff',
                fontSize: '12px',
                cursor: 'pointer',
                display: 'flex',
                alignItems: 'center',
                gap: '6px'
              }}
            >
              <Download size={14} />
              Download
            </button>
          </div>
          
          <div style={{
            display: 'grid',
            gridTemplateColumns: '1fr 1fr',
            gap: '20px'
          }}>
            <div>
              <img 
                src={`data:image/jpeg;base64,${result.image_data}`}
                alt="Generated professional image"
                style={{
                  width: '100%',
                  borderRadius: '12px',
                  border: '1px solid rgba(255,255,255,0.1)'
                }}
              />
            </div>
            <div>
              <div style={{
                color: 'rgba(255,255,255,0.9)',
                fontSize: '14px',
                lineHeight: '1.6',
                background: 'rgba(0,0,0,0.3)',
                padding: '15px',
                borderRadius: '8px'
              }}>
                <strong>Enriched Prompt:</strong><br/>
                {result.enriched_prompt}
              </div>
              <div style={{
                marginTop: '10px',
                fontSize: '12px',
                color: 'rgba(255,255,255,0.6)'
              }}>
                Request ID: {result.request_id}
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Error display */}
      {error && (
        <div style={{
          marginTop: '20px',
          background: 'rgba(255,0,0,0.1)',
          border: '1px solid rgba(255,0,0,0.3)',
          borderRadius: '16px',
          padding: '20px',
          color: '#ff6b6b'
        }}>
          {error}
        </div>
      )}
    </div>
  );
}
