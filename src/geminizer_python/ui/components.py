from rich.console import Console
from rich.panel import Panel
import typer

console = Console()

def render_input_form() -> tuple[str, str]:
    """Renders a color-coded input form and collects prompt and style."""
    console.print(Panel(
        "[bold cyan]Please enter your prompt details:[/bold cyan]",
        title="[bold magenta]Enhancement Form[/bold magenta]",
        border_style="cyan"
    ))
    prompt = typer.prompt("Enter your [bold white]prompt[/bold white]")
    style = typer.prompt("Enter the [bold white]style[/bold white]", default="photorealistic")
    return prompt, style

def render_result_panel(enhanced_prompt: str) -> None:
    """Renders the final result in a stylized panel."""
    console.print(Panel(
        f"[green]{enhanced_prompt}[/green]",
        title="[bold green]Enhanced Prompt[/bold green]",
        border_style="green"
    ))
