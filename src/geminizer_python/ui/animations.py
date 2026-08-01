from rich.console import Console
from contextlib import contextmanager
from typing import Generator

console = Console()

@contextmanager
def waiting_animation(message: str = "Processing...") -> Generator[None, None, None]:
    """Shows a waiting status animation."""
    with console.status(f"[bold yellow]{message}[/bold yellow]", spinner="dots"):
        yield
