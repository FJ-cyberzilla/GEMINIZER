import pytest
from src.geminizer_python.core.models import Context
from src.geminizer_python.core.orchestrator import Orchestrator
from src.geminizer_python.tasks.example_task import ExampleTask

def test_orchestrator_runs_task() -> None:
    orchestrator = Orchestrator()
    orchestrator.add_task(ExampleTask())
    results = orchestrator.run_all(Context())
    assert len(results) == 1
    assert results[0].success is True
    assert results[0].message == "Example task completed."
