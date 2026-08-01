import pytest
import os
from src.geminizer_python.core.models import Context
from src.geminizer_python.tasks.file_analysis import FileAnalysisTask

from pathlib import Path
from typing import Any

def test_file_analysis_success(tmp_path: Path) -> None:
    # Create a dummy file
    d = tmp_path / "sub"
    d.mkdir()
    p = d / "hello.txt"
    p.write_text("line1\nline2\nline3")
    
    context = Context(data={"file_path": str(p)})
    task = FileAnalysisTask()
    
    result = task.execute(context)
    
    assert result.success is True
    assert result.data["line_count"] == 3

def test_file_analysis_file_not_found() -> None:
    context = Context(data={"file_path": "non_existent.txt"})
    task = FileAnalysisTask()
    
    result = task.execute(context)
    
    assert result.success is False
    assert "File not found" in result.message
