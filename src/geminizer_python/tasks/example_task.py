
from src.geminizer_python.tasks.base import Task
from src.geminizer_python.core.models import Context, TaskResult

class ExampleTask(Task):
    """An example task implementation."""
    def execute(self, context: Context) -> TaskResult:
        print(f"Executing example task with context: {context}")
        return TaskResult(success=True, message="Example task completed.")
