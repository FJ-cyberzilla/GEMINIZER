import typer
from rich.console import Console
from rich.panel import Panel
from src.geminizer_python.core.orchestrator import Orchestrator
from src.geminizer_python.tasks.example_task import ExampleTask
from src.geminizer_python.tasks.file_analysis import FileAnalysisTask
from src.geminizer_python.tasks.prompt_task import PromptTask
from src.geminizer_python.core.models import Context

app = typer.Typer()
console = Console()
orchestrator = Orchestrator()
orchestrator.add_task(ExampleTask())

@app.command()
def run() -> None:
    """Runs all registered tasks."""
    console.print(Panel("🚀 [bold blue]Geminizer Running Tasks[/bold blue]", title="Geminizer", expand=False))
    context = Context()
    results = orchestrator.run_all(context)
    for result in results:
        console.print(f"Result: {result.message}")

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

from src.geminizer_python.ui.components import render_input_form, render_result_panel
from src.geminizer_python.ui.animations import waiting_animation

# ... (other imports)

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
