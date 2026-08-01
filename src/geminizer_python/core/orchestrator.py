from typing import List
from src.geminizer_python.tasks.base import Task
from src.geminizer_python.core.models import Context, TaskResult

class Orchestrator:
    """Manages and executes tasks."""
    def __init__(self) -> None:
        self.tasks: List[Task] = []

    def add_task(self, task: Task) -> None:
        self.tasks.append(task)

    def run_all(self, context: Context) -> List[TaskResult]:
        results = []
        for task in self.tasks:
            results.append(task.execute(context))
        return results
