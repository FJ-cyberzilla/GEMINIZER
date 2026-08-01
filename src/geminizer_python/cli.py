import os
import typer
from rich.panel import Panel
from src.geminizer_python.core.orchestrator import Orchestrator
from src.geminizer_python.tasks.example_task import ExampleTask
from src.geminizer_python.tasks.file_analysis import FileAnalysisTask
from src.geminizer_python.tasks.prompt_task import PromptTask
from src.geminizer_python.core.models import Context
from src.geminizer_python.ui.components import (
    render_input_form, 
    render_result_panel, 
    render_header, 
    render_chat_input,
    console  # Using the console instance with the custom theme
)
from src.geminizer_python.ui.animations import waiting_animation

app = typer.Typer()
orchestrator = Orchestrator()
orchestrator.add_task(ExampleTask())

def ensure_api_key() -> bool:
    """Checks for GOOGLE_API_KEY and guides the user if missing."""
    if not os.environ.get("GOOGLE_API_KEY"):
        console.print(Panel(
            "[error]API Key Missing![/error]\n\n"
            "To use Geminizer, you need a Gemini API key.\n"
            "1. Get one at [link=https://aistudio.google.com/]Google AI Studio[/link]\n"
            "2. Set it in your terminal:\n"
            "   [warning]export GOOGLE_API_KEY='your_key_here'[/warning]\n\n"
            "You can also set it temporarily in this session.",
            title="[title]Setup Required[/title]",
            border_style="error"
        ))
        return False
    return True

@app.command()
def run() -> None:
    """Runs all registered tasks."""
    render_header()
    context = Context()
    results = orchestrator.run_all(context)
    for result in results:
        console.print(f"Result: {result.message}")

@app.command()
def chat() -> None:
    """Interactive chat interface."""
    if not ensure_api_key():
        return
    render_header()
    sentence = render_chat_input()
    
    context = Context(data={"raw_prompt": sentence, "style": "default"})
    task = PromptTask()
    
    with waiting_animation("Processing your prompt..."):
        result = task.execute(context)
        
    if result.success:
        render_result_panel(result.data['enhanced_prompt'])
    else:
        console.print(f"❌ [bold red]Error:[/bold red] {result.message}")

@app.command()
def analyze(file_path: str) -> None:
    """Analyzes a specific file."""
    console.print(f"🔍 [bold green]Analyzing file: {file_path}[/bold green]")
    context = Context(data={"file_path": file_path})
    task = FileAnalysisTask()
    result = task.execute(context)
    if result.success:
        console.print(f"✅ [bold green]Success![/bold green] Line count: {result.data['line_count']}")
    else:
        console.print(f"❌ [bold red]Error:[/bold red] {result.message}")

@app.command()
def enhance() -> None:

    """Enhances a prompt for AI image generation using an interactive UI."""
    prompt, style = render_input_form()
    
    context = Context(data={"raw_prompt": prompt, "style": style})
    task = PromptTask()
    
    with waiting_animation("Enhancing your prompt..."):
        result = task.execute(context)
        
    if result.success:
        render_result_panel(result.data['enhanced_prompt'])
    else:
        console.print(f"❌ [bold red]Error:[/bold red] {result.message}")

@app.command()
def info() -> None:
    """Shows application info."""
    console.print("[blue]Geminizer CLI v0.1.0[/blue]")

if __name__ == "__main__":
    app()
