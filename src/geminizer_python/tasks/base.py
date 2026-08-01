from abc import ABC, abstractmethod
from src.geminizer_python.core.models import Context, TaskResult

class Task(ABC):
    """Abstract Base Class for all tasks."""

    @abstractmethod
    def execute(self, context: Context) -> TaskResult:
        """Executes the task with the given context."""
        pass
