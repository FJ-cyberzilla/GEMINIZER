from rich.console import Console
from rich.panel import Panel
from rich.align import Align
from rich.prompt import Prompt
from src.geminizer_python.ui.theme import geminizer_theme

console = Console(theme=geminizer_theme)

def render_header() -> None:
    """Renders a stylized Geminizer header."""
    header = Panel(
        Align.center("[bold cyan]Geminizer CLI[/bold cyan]\n[dim]AI-Powered Automation Framework[/dim]", vertical="middle"),
        border_style="border",
        expand=False,
        padding=(1, 2)
    )
    console.print(header)

def render_input_form() -> tuple[str, str]:
    """Renders a color-coded input form and collects prompt and style."""
    console.print(Panel(
        "[info]Please enter your prompt details:[/info]",
        title="[title]Enhancement Form[/title]",
        border_style="border"
    ))
    prompt = Prompt.ask("Enter your [prompt]prompt[/prompt]")
    style = Prompt.ask("Enter the [prompt]style[/prompt]", default="photorealistic")
    return prompt, style

def render_chat_input() -> str:
    """Renders a panel for chat input and collects user prompt."""
    console.print(Panel(
        "[info]Please copy or put your prompt.[/info]\n[dim]Press Ctrl+C to exit.[/dim]",
        title="[title]Chat[/title]",
        border_style="border"
    ))
    return Prompt.ask("Enter your prompt")

def render_result_panel(enhanced_prompt: str) -> None:
    """Renders the final result in a stylized panel."""
    console.print(Panel(
        f"[success]{enhanced_prompt}[/success]",
        title="[title]Enhanced Prompt[/title]",
        border_style="success"
    ))
