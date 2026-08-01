import os
from typing import Optional
from google import genai
from src.geminizer_python.tasks.base import Task
from src.geminizer_python.core.models import Context, TaskResult

class PromptTask(Task):
    """Enhances a prompt for AI image generation."""
    
    def __init__(self, model_name: str = "gemini-3.5-flash", client: Optional[genai.Client] = None):
        self.client = client or genai.Client() # Assumes GOOGLE_API_KEY is set in environment
        self.model_name = model_name
        self.system_prompt = (
            "You are an expert AI image prompt engineer. Your goal is to transform a user's "
            "basic, vague, or short prompt into a highly detailed, professional, industrial-level "
            "prompt optimized for high-quality image generation models (like Midjourney, DALL-E, "
            "Stable Diffusion). Maintain the user's core intent but add rich details regarding "
            "lighting, style, composition, camera settings, texture, and technical fidelity."
        )
    
    def execute(self, context: Context) -> TaskResult:
        raw_prompt = context.data.get("raw_prompt")
        style = context.data.get("style", "photorealistic")
        if not raw_prompt:
            return TaskResult(success=False, message="No prompt provided.")
            
        try:
            # Dynamically inject style into instructions
            style_instruction = f"\n\nStyle: {style}. Adapt the prompt to strictly adhere to the aesthetics of this style, incorporating specific visual, lighting, and composition elements characteristic of this aesthetic."
            
            response = self.client.models.generate_content(
                model=self.model_name,
                contents=f"{self.system_prompt}{style_instruction}\n\nUser Prompt: {raw_prompt}",
            )
            
            return TaskResult(
                success=True,
                message="Enhanced prompt successfully.",
                data={"enhanced_prompt": response.text}
            )
        except Exception as e:
            return TaskResult(success=False, message=f"Failed to enhance prompt: {str(e)}")
