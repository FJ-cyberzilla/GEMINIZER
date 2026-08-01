import os
from src.geminizer_python.tasks.base import Task
from src.geminizer_python.core.models import Context, TaskResult

class FileAnalysisTask(Task):
    """Analyzes a file by counting lines."""
    
    def execute(self, context: Context) -> TaskResult:
        file_path = context.data.get("file_path")
        if not file_path or not os.path.exists(file_path):
            return TaskResult(success=False, message=f"File not found: {file_path}")
        
        with open(file_path, "r") as f:
            lines = f.readlines()
            
        return TaskResult(
            success=True, 
            message=f"Analyzed {file_path}", 
            data={"line_count": len(lines)}
        )
