from dataclasses import dataclass, field
from typing import Dict, Any

@dataclass
class Context:
    """Shared execution context."""
    data: Dict[str, Any] = field(default_factory=dict)

@dataclass
class TaskResult:
    """Standardized result for a task."""
    success: bool
    message: str
    data: Any = None
